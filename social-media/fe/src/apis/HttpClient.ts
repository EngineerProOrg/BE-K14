import axios, { AxiosError, AxiosResponse } from 'axios';
import { SignInRequestViewModel, SignInResponseViewModel } from '../models/user';

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8080/api/v1',
  withCredentials: true,
});

const sleep = (delay: number) =>
  new Promise(resolve => setTimeout(resolve, delay));

axiosInstance.interceptors.response.use(
  async (response) => {
    await sleep(300);
    return response.data;
  },
  (error: AxiosError) => {
    const { response } = error;

    if (!response) throw error;

    if (response.status === 404) {
      console.error('404 Not Found');
    }

    throw error;
  }
);

axiosInstance.interceptors.response.use(
  (response) => response.data,
  (error) => Promise.reject(error)
);


const Auth = {
  SignIn: (body: SignInRequestViewModel): Promise<SignInResponseViewModel> =>
    axiosInstance.post('/users/signin', body),
};


const HttpClient = {
  Auth
};

export default HttpClient;
