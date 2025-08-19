# CORS Configuration Guide

## Overview

The backend now supports flexible CORS (Cross-Origin Resource Sharing) configuration to allow your frontend deployed on different platforms to access the backend API.

## Current Configuration

### Allowed Origins
By default, the following origins are allowed to access the backend:

1. **Local Development**:
   - `http://localhost:3000` (React default)
   - `http://localhost:5173` (Vite default)

2. **Production Deployment**:
   - `https://tt-ms.vercel.app` (Your Vercel deployment)

## Environment Configuration

### Backend (.env)
```bash
# CORS Configuration
# Frontend origins that can access this backend
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,https://tt-ms.vercel.app
```

### Frontend (.env)
```bash
NODE_ENV='development'
VITE_BACKEND_URL='https://tt-ms.onrender.com'
VITE_API_BASE_URL='https://tt-ms.onrender.com/api/v1'
```

## How It Works

### 1. **Environment Variable Override**
- If `CORS_ALLOWED_ORIGINS` is set, it uses those origins
- If not set, falls back to default origins
- Origins are comma-separated in the environment variable

### 2. **Dynamic Origin Parsing**
```go
// Backend automatically parses comma-separated origins
envOrigins := "http://localhost:3000,https://myapp.vercel.app"
// Results in: ["http://localhost:3000", "https://myapp.vercel.app"]
```

### 3. **CORS Headers**
The backend sends appropriate CORS headers:
- `Access-Control-Allow-Origin`
- `Access-Control-Allow-Methods`
- `Access-Control-Allow-Headers`
- `Access-Control-Allow-Credentials`

## Adding New Origins

### For Development
Add new local development URLs:
```bash
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,http://localhost:8080
```

### For Production
Add new deployment URLs:
```bash
CORS_ALLOWED_ORIGINS=http://localhost:3000,http://localhost:5173,https://tt-ms.vercel.app,https://staging.myapp.com
```

## Deployment Scenarios

### Scenario 1: Frontend on Vercel, Backend on Render
```bash
# Backend .env (on Render)
CORS_ALLOWED_ORIGINS=https://tt-ms.vercel.app,http://localhost:5173

# Frontend .env (for Vercel)
VITE_API_BASE_URL=https://tt-ms.onrender.com/api/v1
```

### Scenario 2: Both on Same Platform
```bash
# Backend .env
CORS_ALLOWED_ORIGINS=https://myapp.platform.com,http://localhost:5173

# Frontend .env
VITE_API_BASE_URL=https://api.myapp.platform.com/api/v1
```

### Scenario 3: Multiple Environments
```bash
# Backend .env
CORS_ALLOWED_ORIGINS=https://prod.myapp.com,https://staging.myapp.com,https://dev.myapp.com,http://localhost:5173
```

## Security Considerations

### ✅ **Good Practices**
- Only add trusted domains to allowed origins
- Use HTTPS in production
- Avoid wildcard origins (`*`) in production
- Regularly review and update allowed origins

### ❌ **Avoid**
```bash
# DON'T: Allow all origins (security risk)
CORS_ALLOWED_ORIGINS=*

# DON'T: Mix HTTP and HTTPS carelessly
CORS_ALLOWED_ORIGINS=http://myapp.com,https://myapp.com
```

## Troubleshooting

### Error: "CORS policy: No 'Access-Control-Allow-Origin' header"

**Solution**: Add your frontend URL to `CORS_ALLOWED_ORIGINS`
```bash
CORS_ALLOWED_ORIGINS=existing-origins,https://your-new-domain.com
```

### Error: "Credentials not allowed"

**Solution**: Ensure `AllowCredentials: true` is set (already configured)

### Error: "Method not allowed"

**Solution**: Check if your HTTP method is in the allowed methods list:
- Currently allowed: `GET`, `POST`, `PUT`, `DELETE`, `OPTIONS`

## Testing CORS

### 1. Browser Dev Tools
Check the Network tab for CORS-related errors

### 2. curl Command
```bash
curl -H "Origin: https://tt-ms.vercel.app" \
     -H "Access-Control-Request-Method: GET" \
     -H "Access-Control-Request-Headers: X-Requested-With" \
     -X OPTIONS \
     https://tt-ms.onrender.com/api/v1/health
```

### 3. Browser Console
```javascript
fetch('https://tt-ms.onrender.com/api/v1/health', {
  method: 'GET',
  credentials: 'include'
})
.then(response => console.log('Success:', response))
.catch(error => console.log('CORS Error:', error));
```

## Current Setup Summary

✅ **https://tt-ms.vercel.app** is now allowed to access your backend  
✅ **Local development** (localhost:5173) is allowed  
✅ **Environment-based configuration** for easy deployment  
✅ **Secure defaults** with credentials support
