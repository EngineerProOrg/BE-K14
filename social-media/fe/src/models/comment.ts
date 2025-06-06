import { UserBaseViewModel } from "./user";

export interface CommentResponseViewModel {
  id: number;
  content: string;
  createdAt: Date;
  updatedAt: Date | null;
  author: UserBaseViewModel;
}
