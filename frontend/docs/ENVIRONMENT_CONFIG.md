# Environment Configuration

## Frontend Environment Variables

The frontend uses environment variables to configure the backend connection. This allows for easy switching between development, staging, and production environments.

### Required Environment Variables

Create a `.env` file in the frontend directory with the following variables:

```bash
# Node environment
NODE_ENV='development'

# Backend server URL (without /api/v1)
VITE_BACKEND_URL='http://localhost:8080'

# API base URL (with /api/v1)
VITE_API_BASE_URL='http://localhost:8080/api/v1'
```

### Environment-Specific Configurations

#### Development
```bash
VITE_BACKEND_URL='http://localhost:8080'
VITE_API_BASE_URL='http://localhost:8080/api/v1'
```

#### Production
```bash
VITE_BACKEND_URL='https://your-production-domain.com'
VITE_API_BASE_URL='https://your-production-domain.com/api/v1'
```

#### Staging
```bash
VITE_BACKEND_URL='https://staging.your-domain.com'
VITE_API_BASE_URL='https://staging.your-domain.com/api/v1'
```

### Usage in Code

The project uses a centralized configuration system:

```javascript
// Import the API configuration
import { API_CONFIG, buildApiUrl } from '../config/api.js';

// Use the configuration
const response = await fetch(buildApiUrl('/login'), {
  method: 'POST',
  // ... other options
});
```

### Vite Configuration

The Vite development server is also configured to use environment variables for the proxy configuration, allowing the proxy to automatically point to the correct backend URL.

### Files Updated

1. **`.env`** - Environment variables
2. **`src/config/api.js`** - Centralized API configuration
3. **`src/services/api.js`** - Axios instance configuration
4. **`src/context/AuthContext.jsx`** - Authentication API calls
5. **`vite.config.ts`** - Vite proxy configuration

### Benefits

- **Environment Flexibility**: Easy switching between different environments
- **Centralized Configuration**: All API URLs managed in one place
- **Type Safety**: Helper functions prevent URL construction errors
- **Maintainability**: Single source of truth for backend URLs
