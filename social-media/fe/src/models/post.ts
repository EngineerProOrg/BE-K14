import { UserBaseViewModel } from "./user";

export interface PostBaseViewModel {
  postId: number;
  title: string;
  content: string;
  createdAt: Date | null;
  updatedAt: Date | null;
}

export interface PostWithAuthorViewModel {
  post: PostBaseViewModel;
  author: UserBaseViewModel;
}

export interface PostsWithAuthorResponse {
  posts: PostWithAuthorViewModel[];
}
