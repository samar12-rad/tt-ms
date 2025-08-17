import axios from './api.js';

const backendService = {

  get: async (url, data = {}) => {
    try {
      const response = await axios.get(url, {
        params: data,
      });
      return response.data;
    } catch (err) {
      throw err;
    }
  },

  post: async (url, data = {}) => {
    try {
      const response = await axios.post(url, data);
      return response.data;
    } catch (err) {
      throw err;
    }
  },

  put: async (url, data = {}) => {
    try {
      const response = await axios.put(url, data);
      return response.data;
    } catch (err) {
      throw err;
    }
  },

  delete: async (url, data = {}) => {
    try {
      const response = await axios.delete(url, { data }); // body in config
      return response.data;
    } catch (err) {
      throw err;
    }
  }
};

export default backendService;
