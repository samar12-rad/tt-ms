# TMS-backend
Timetable management system Backend

## Features
- **Connection Pooling**: Efficient database connection management using pgxpool
- **JWT Authentication**: Secure user authentication and authorization
- **Role-based Access Control**: Admin, Faculty, and Student roles
- **RESTful API**: Clean and consistent API design
- **Database Migrations**: Automated database schema management
- **Flexible CORS**: Environment-based CORS configuration for deployment
- **Health Monitoring**: Connection pool statistics and health checks

## Build & Run
- cd into `backend/`
- run `go mod download` first to install go packages
- run directly with `go run .` command or build binary and execute it `go build . && ./tms-sever`

## Development
- For hot-reloading install `air`: [github.com/air-verse/air](https://github.com/air-verse/air)
- Copy `.env.example` to `env`: `cp .env.example .env`
- Inside `.env` add your own postgres DATABASE_URL, Recommeneded to get from [Supabase](https://supabase.com).

## Database Connection Pooling

This application uses **dual-layer connection pooling**:
1. **Supabase Transaction Pooling** - External pooling managed by Supabase
2. **Application Connection Pooling** - Internal pooling using pgxpool

See [docs/SUPABASE_POOLING.md](docs/SUPABASE_POOLING.md) for detailed information about the dual-pooling strategy.

### Quick Setup for Supabase
1. Use the Supabase pooler connection string (port 6543)
2. Configure smaller local pool since Supabase handles external pooling
3. Monitor both application and Supabase pool performance

### Environment Variables for Connection Pool
```bash
# Supabase Pooler Connection
PG_HOST=aws-0-us-east-2.pooler.supabase.com
PG_PORT=6543
PG_USER=postgres.your_ref_id
PG_PASSWORD=your_password

# Optimized for Supabase (smaller local pool)
PG_MAX_CONNS=10              # Reduced for Supabase pooling
PG_MIN_CONNS=2               # Minimal persistent connections
PG_MAX_CONN_LIFETIME=30m     # Shorter rotation
PG_MAX_CONN_IDLE_TIME=15m    # Quick cleanup
```

## CORS Configuration

Configure allowed frontend origins for production deployment:
```bash
# Frontend origins that can access this backend
CORS_ALLOWED_ORIGINS=http://localhost:5173,https://tt-ms.vercel.app
```

See [docs/CORS_CONFIGURATION.md](docs/CORS_CONFIGURATION.md) for detailed CORS setup guide.

## API Endpoints Documentation

### Base URL
`/api/v1`

---

### System Routes (Public)
- `GET /health` - Service health check
- `GET /stats` - Database connection pool statistics

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
