package main

import (
	"fmt"
	"log"
	"time"
	"tms-server/config"
	"tms-server/models"
)

func init() {
	config.LoadEnvVariables()
}

func main() {
	config.ConnectDB()

	// Get count of lectures
	var lectureCount int64
	if err := config.DB.Model(&models.Lecture{}).Count(&lectureCount).Error; err != nil {
		log.Fatalf("Failed to count lectures: %v", err)
	}

	fmt.Printf("Total lectures in database: %d\n", lectureCount)

	if lectureCount == 0 {
		fmt.Println("No lectures found. No sessions will be created.")
		return
	}

	// Calculate for 35 days
	days := 35
	
	// Count lectures by day of week to get a more accurate estimate
	dayOfWeekCounts := make(map[string]int64)
	days_of_week := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}
	
	for _, day := range days_of_week {
		var count int64
		config.DB.Model(&models.Lecture{}).Where("LOWER(day_of_week) = LOWER(?)", day).Count(&count)
		dayOfWeekCounts[day] = count
		fmt.Printf("Lectures on %s: %d\n", day, count)
	}

	// Calculate estimated sessions for 35 days
	totalEstimatedSessions := int64(0)
	startDate := time.Now().AddDate(0, 0, -days+1)
	
	for i := 0; i < days; i++ {
		date := startDate.AddDate(0, 0, i)
		dayOfWeek := date.Weekday().String()
		totalEstimatedSessions += dayOfWeekCounts[dayOfWeek]
	}

	// Check how many sessions already exist for this date range
	endDate := time.Now()
	var existingSessions int64
	config.DB.Model(&models.Session{}).
		Where("date >= ? AND date <= ?", startDate.Format("2006-01-02"), endDate.Format("2006-01-02")).
		Count(&existingSessions)

	newSessions := totalEstimatedSessions - existingSessions
	if newSessions < 0 {
		newSessions = 0
	}

	fmt.Printf("\n=== ESTIMATION FOR %d DAYS ===\n", days)
	fmt.Printf("Date range: %s to %s\n", startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	fmt.Printf("Total estimated sessions: %d\n", totalEstimatedSessions)
	fmt.Printf("Existing sessions in range: %d\n", existingSessions)
	fmt.Printf("New sessions to be created: %d\n", newSessions)

	// Estimate execution time
	// Assuming each database operation takes about 10-50ms
	avgTimePerSession := 30 * time.Millisecond
	estimatedTime := time.Duration(newSessions) * avgTimePerSession
	
	fmt.Printf("\n=== TIME ESTIMATION ===\n")
	fmt.Printf("Estimated time per session: %v\n", avgTimePerSession)
	fmt.Printf("Estimated total execution time: %v\n", estimatedTime)
	
	if estimatedTime.Minutes() > 1 {
		fmt.Printf("Estimated execution time (minutes): %.2f\n", estimatedTime.Minutes())
	}
	if estimatedTime.Seconds() > 60 {
		fmt.Printf("Estimated execution time (seconds): %.1f\n", estimatedTime.Seconds())
	}
}
