import axios, { AxiosError } from "axios";
import {
  SignInRequestViewModel,
  SignInResponseViewModel,
  SignUpRequestViewModel,
  UserBaseViewModel,
} from "../models/user";
import { ErrorResponseViewModel } from "../models/error";
import { PostsWithAuthorResponse } from "../models/post";
import { CommentResponseViewModel } from "../models/comment";
import { UserReactionResponseViewModel } from "../models/reaction";

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
    const status = error.response?.status;
    const data = error.response?.data;

    if (status === 401) {
      localStorage.removeItem("accessToken");
      window.location.href = "/signin";
    }

    const message =
      typeof data === "object" && data && "error" in data
        ? (data as ErrorResponseViewModel).error
        : "Something went wrong";

    return Promise.reject(new Error(message));
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

const Comment = {
  GetCommentsByPostId: (postId: number): Promise<CommentResponseViewModel> => {
    return axiosInstance.get(`/posts/${postId}/comments`);
  },
};

const Reaction = {
  GetReactionsByTarget: (
    targetId: number,
    targetType: "post" | "comment"
  ): Promise<UserReactionResponseViewModel> => {
    return axiosInstance.get(`/reactions/${targetId}`, {
      params: { target_type: targetType },
    });
  },
};

const HttpClient = {
  User,
  Post,
  Comment,
  Reaction
};

export default HttpClient;
