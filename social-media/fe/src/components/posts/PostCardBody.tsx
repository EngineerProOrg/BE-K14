import { CardContent, Chip, Typography } from "@mui/material";
import { PostBaseViewModel } from "../../models/post";

interface PostCardBodyProps {
  post: PostBaseViewModel;
}

function PostCardBody({ post }: PostCardBodyProps) {
  return (
    <>
      <CardContent>
        <Typography variant="h6">
          {post.title}
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
          {post.content}
        </Typography>
        <Chip
          label="ĐÓN TƯƠNG LAI"
          color="primary"
          size="small"
          sx={{ mt: 1 }}
        />
      </CardContent>
    </>
  );
}

export default PostCardBody;
