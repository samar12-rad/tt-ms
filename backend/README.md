# TMS-backend
Timetable management system Backend

## Build & Run
- cd into `backend/`
- run `go mod download` first to install go packages
- run directly with `go run .` command or build binary and execute it `go build . && ./tms-sever`

## Development
- For hot-reloading install `air`: [github.com/air-verse/air](https://github.com/air-verse/air)
- Copy `.env.example` to `env`: `cp .env.example .env`
- Inside `.env` add your own postgres DATABASE_URL, Recommeneded to get from [Supabase](https://supabase.com).

## API Endpoints Documentation

### Base URL
`/api/v1`

---

### Public Routes (No JWT Required)
- `GET /ping` - Service health check
- `POST /login` - User authentication

---

### Protected Routes (JWT Required)
All endpoints below require valid JWT authentication

#### Course Management
- `GET /course` - Get all courses
- `POST /course` - Create new course
- `GET /course/:id` - Get single course
- `PUT /course/:id` - Update course
- `DELETE /course/:id` - Delete course

#### Subject Management
- `GET /subject` - Get all subjects
- `POST /subject` - Create new subject
- `GET /subject/:id` - Get single subject
- `PUT /subject/:id` - Update subject
- `DELETE /subject/:id` - Delete subject

#### Faculty Management
- `GET /faculty` - Get all faculties
- `POST /faculty` - Create new faculty
- `GET /faculty/:id` - Get single faculty
- `PUT /faculty/:id` - Update faculty
- `DELETE /faculty/:id` - Delete faculty

#### Room Management
- `GET /room` - Get all rooms
- `POST /room` - Create new room
- `GET /room/:id` - Get single room
- `PUT /room/:id` - Update room
- `DELETE /room/:id` - Delete room

#### User Management (Experimental)
- `GET /user` - Get all users
- `POST /user` - Create new user
- `GET /user/:id` - Get single user
- `PUT /user/:id` - Update user
- `DELETE /user/:id` - Delete user

#### Lecture Management (Experimental)
- `GET /lecture` - Get all timetable entries
- `POST /lecture` - Create new timetable entry
- `GET /lecture/:id` - Get single timetable entry
- `PUT /lecture/:id` - Update timetable entry
- `DELETE /lecture/:id` - Delete timetable entry

---

## Access Notes
- All endpoints except `/ping` and `/login` require JWT authentication
- Include JWT token in cookie for protected routes
- Replace `:id` in URLs with actual resource IDs
- Experimental endpoints may have limited functionality
- CORS enabled for all routes
