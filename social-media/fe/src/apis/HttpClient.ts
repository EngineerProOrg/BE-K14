import axios, { AxiosError } from "axios";
import {
  SignInRequestViewModel,
  SignInResponseViewModel,
} from "../models/user";
import { ErrorResponseViewModel } from "../models/error";

const axiosInstance = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  withCredentials: true,
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
  },
});

const sleep = (delay: number) =>
  new Promise((resolve) => setTimeout(resolve, delay));

axiosInstance.interceptors.response.use(
  async (response) => {
    await sleep(300);
    return response.data;
  },
  (error: AxiosError) => {
    const { response } = error;

    const customError =
      response?.data &&
      typeof response.data === "object" &&
      "error" in response.data
        ? (response.data as ErrorResponseViewModel).error
        : "Undefined error from server";

    return Promise.reject(customError);
  }
);

const User = {
  SignIn: (body: SignInRequestViewModel): Promise<SignInResponseViewModel> =>
    axiosInstance.post("/users/signin", body),
};

const HttpClient = {
  User,
};

export default HttpClient;
