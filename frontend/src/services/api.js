import axios from "axios";
import { API_CONFIG } from "../config/api.js";

const instance = axios.create({
  baseURL: API_CONFIG.BACKEND_URL,
  withCredentials: true,
});

instance.interceptors.response.use(
  (res) => res,
  (err) => {
    if (err.response && err.response.status === 401) {
      window.location.href = "/";
    }
    return Promise.reject(err);
  }
);

export default instance;
