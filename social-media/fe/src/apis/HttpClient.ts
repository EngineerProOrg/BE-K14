import axios, { AxiosError } from "axios";
import {
  SignInRequestViewModel,
  SignInResponseViewModel,
  SignUpRequestViewModel,
  UserBaseViewModel,
} from "../models/user";
import { ErrorResponseViewModel } from "../models/error";
import { PostsWithAuthorResponse } from "../models/post";

const axiosInstance = axios.create({
  baseURL: "http://localhost:8080/api/v1",
  withCredentials: true,
  headers: {
    "Content-Type": "application/json",
    Accept: "application/json",
  },
});

const sleep = (delay: number) =>
  new Promise((resolve) => setTimeout(resolve, delay));

axiosInstance.interceptors.request.use(
  (config) => {
    const accessToken = localStorage.getItem("accessToken");
    if (accessToken) {
      config.headers["Authorization"] = `Bearer ${accessToken}`;
    }
    return config;
  },
  (error) => {
    return Promise.reject(error);
  }
);


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
  SignIn: (
    signInRequestViewModel: SignInRequestViewModel
  ): Promise<SignInResponseViewModel> =>
    axiosInstance.post("/users/signin", signInRequestViewModel),
  SignUp: (
    signUpRequestViewModel: SignUpRequestViewModel
  ): Promise<UserBaseViewModel> =>
    axiosInstance.post("/users/signup", signUpRequestViewModel),
};

const Post = {
  GetAllPosts: (): Promise<PostsWithAuthorResponse> => {
    return axiosInstance.get("/posts");
  },
};

const HttpClient = {
  User,
  Post
};

export default HttpClient;
