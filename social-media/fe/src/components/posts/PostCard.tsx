import { Card } from "@mui/material";
import PostCardHeader from "./PostCardHeader";
import PostCardBody from "./PostCardBody";
import PostCardReaction from "./PostCardReaction";
import PostComment from "./PostCardComment";
import { PostWithAuthorViewModel } from "../../models/post";

interface PostCardProps {
  postWithAuthor: PostWithAuthorViewModel;
}

export default function PostCard({ postWithAuthor }: PostCardProps) {
  return (
    <Card sx={{ maxWidth: 600, margin: "auto", mt: 3 }}>
      <PostCardHeader postWithAuthor={postWithAuthor}/>
      <PostCardBody post={postWithAuthor.post}/>
      <PostCardReaction />
      <PostComment />
    </Card>
  );
}
