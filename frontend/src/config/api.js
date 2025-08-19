// API configuration utility
export const API_CONFIG = {
  BACKEND_URL: import.meta.env.VITE_BACKEND_URL || 'http://localhost:8080',
  API_BASE_URL: import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080/api/v1',
};

// Helper function to build API URLs
export const buildApiUrl = (endpoint) => {
  const baseUrl = API_CONFIG.API_BASE_URL;
  return endpoint.startsWith('/') ? `${baseUrl}${endpoint}` : `${baseUrl}/${endpoint}`;
};

// Helper function to build backend URLs
export const buildBackendUrl = (endpoint) => {
  const baseUrl = API_CONFIG.BACKEND_URL;
  return endpoint.startsWith('/') ? `${baseUrl}${endpoint}` : `${baseUrl}/${endpoint}`;
};

export default API_CONFIG;
