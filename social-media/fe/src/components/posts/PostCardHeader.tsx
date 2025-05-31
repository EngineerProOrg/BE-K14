import { CardMedia } from "@mui/material";

import CardAuthor from "../userProfile/CardAuthor";
import { PostWithAuthorViewModel } from "../../models/post";

interface PostCardHeaderProps {
  postWithAuthor: PostWithAuthorViewModel;
}

export default function PostCardHeader({
  postWithAuthor,
}: PostCardHeaderProps) {
  return (
    <>
      <CardAuthor
        avatarSrc="/avatar.png"
        title={postWithAuthor.author.name}
        subheader={postWithAuthor.author.username}
      />

      <CardMedia
        component="img"
        height="180"
        image="/images/posts/paella.jpg"
        alt="Future"
      />
    </>
  );
}
