import axios from 'axios';
import router from '@/router/router';

const api = axios.create({
  baseURL: 'http://localhost:8080',
  withCredentials: true,
});

api.interceptors.response.use(
  response => response,
  error => {
    if (error.response && error.response.status === 401) {
      console.warn('Usuário não autorizado. Redirecionando para login...');
      router.push('/login')
    }

    return Promise.reject(error);
  }
);

export default api;