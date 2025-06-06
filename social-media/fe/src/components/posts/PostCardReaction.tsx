import { CardActions, IconButton, Typography } from "@mui/material";
import FavoriteIcon from "@mui/icons-material/Favorite";
import ShareIcon from "@mui/icons-material/Share";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";

function PostCardReaction() {
  return (
    <>
      <CardActions disableSpacing>
        <IconButton>
          <ThumbUpIcon />
        </IconButton>
        <IconButton>
          <FavoriteIcon color="error" />
        </IconButton>
        <IconButton>
          <ShareIcon />
        </IconButton>
        <Typography variant="body2" sx={{ ml: 1 }}>
          553
        </Typography>
      </CardActions>
    </>
  );
}

export default PostCardReaction