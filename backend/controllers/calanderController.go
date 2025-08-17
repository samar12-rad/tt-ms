package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
	"tms-server/config"
	"tms-server/models"
)

func GetCalendarSummaryByMonth(c *gin.Context) {
	month := c.Query("month")
	year := c.Query("year")
	semester := c.Query("semester")
	facultyID := c.Query("faculty_id")
	courseID := c.Query("course_id")

	if month == "" || year == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Both 'month' and 'year' query parameters are required.",
		})
		return
	}

	if _, err := strconv.Atoi(month); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'month' parameter. Must be a number."})
		return
	}
	if _, err := strconv.Atoi(year); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid 'year' parameter. Must be a number."})
		return
	}

	query := config.DB.Model(&models.Session{}).
		Joins("JOIN lectures ON lectures.id = sessions.lecture_id").
		Where("EXTRACT(MONTH FROM sessions.date) = ?", month).
		Where("EXTRACT(YEAR FROM sessions.date) = ?", year)

	// Optional Filters (faculty_id, course_id, semester)
	if semester != "" {
		query = query.Where("lectures.semester = ?", semester)
	}
	if facultyID != "" {
		query = query.Where("lectures.faculty_id = ?", facultyID)
	}
	if courseID != "" {
		query = query.
			Joins("JOIN batches ON batches.id = lectures.batch_id").
			Where("batches.course_id = ?", courseID)
	}

	var sessions []models.Session
	err := query.Preload("Lecture").Find(&sessions).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetching sessions"})
		return
	}

	if len(sessions) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no sessions found", "data": []gin.H{}})
		return
	}

	type DayStat struct {
		Held      int
		Cancelled int
		Nil       int
	}
	summary := make(map[string]*DayStat)
	for _, s := range sessions {
		key := s.Date.Format("2006-01-02")
		if summary[key] == nil {
			summary[key] = &DayStat{}
		}
		if s.Status == "held" {
			summary[key].Held++
		} else if s.Status == "cancelled" {
			summary[key].Cancelled++
		} else if s.Status == "" {
			summary[key].Nil++
		}
	}

	result := []gin.H{}
	for dateStr, stat := range summary {
		result = append(result, gin.H{
			"date":            dateStr,
			"total_held":      stat.Held,
			"total_cancelled": stat.Cancelled,
			"no_data":         stat.Nil,
		})
	}

	c.JSON(http.StatusOK, gin.H{"data": result})
}

func GetLectureDetailsByDate(c *gin.Context) {
	dateStr := c.Query("date")
	semester := c.Query("semester")
	facultyID := c.Query("faculty_id")
	courseID := c.Query("course_id")

	if dateStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "date is required in YYYY-MM-DD format"})
		return
	}

	date, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid date format, use YYYY-MM-DD"})
		return
	}

	var sessions []models.Session
	if err := config.DB.Where("date = ?", date).Find(&sessions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch sessions"})
		return
	}

	if len(sessions) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no sessions found", "data": []gin.H{}})
		return
	}

	lectureIDs := make([]uint, 0, len(sessions))
	for _, s := range sessions {
		lectureIDs = append(lectureIDs, s.LectureID)
	}

	lectureQuery := config.DB.
		Preload("Subject").
		Preload("Faculty").
		Preload("Room").
		Preload("Batch.Course").
		Where("lectures.id IN ?", lectureIDs)

	// Optional Filters (faculty_id, course_id, semester)
	if semester != "" {
		lectureQuery = lectureQuery.Where("lectures.semester = ?", semester)
	}
	if facultyID != "" {
		lectureQuery = lectureQuery.Where("lectures.faculty_id = ?", facultyID)
	}
	if courseID != "" {
		lectureQuery = lectureQuery.
			Joins("JOIN batches ON batches.id = lectures.batch_id").
			Where("batches.course_id = ?", courseID)
	}

	var lectures []models.Lecture
	if err := lectureQuery.Find(&lectures).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch lecture details"})
		return
	}

	if len(lectures) == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "no lectures found", "data": []gin.H{}})
		return
	}

	lectureMap := make(map[uint]models.Lecture)
	for _, l := range lectures {
		lectureMap[l.ID] = l
	}

	result := []gin.H{}
	for _, s := range sessions {
		lecture, exists := lectureMap[s.LectureID]
		if !exists {
			continue
		}
		result = append(result, gin.H{
			"lecture_id":    s.LectureID,
			"subject":       lecture.Subject.Name,
			"faculty":       lecture.Faculty.Name,
			"start_time":    lecture.StartTime,
			"end_time":      lecture.EndTime,
			"status":        s.Status,
			"semester":      lecture.Semester,
			"room":          lecture.Room.Name,
			"batch_year":    lecture.Batch.Year,
			"batch_section": lecture.Batch.Section,
			"course_name":   lecture.Batch.Course.Name,
			"session_id":   s.ID,
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"date": dateStr,
		"data": result,
	})
}
