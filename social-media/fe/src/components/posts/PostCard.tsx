import {
  Card,
  CardHeader,
  CardMedia,
  CardContent,
  CardActions,
  Avatar,
  Typography,
  IconButton,
  Chip,
  TextField,
  Box,
} from "@mui/material";
import FavoriteIcon from "@mui/icons-material/Favorite";
import ShareIcon from "@mui/icons-material/Share";
import ThumbUpIcon from "@mui/icons-material/ThumbUp";
import CardAuthor from "../userProfile/CardAuthor";

export default function PostCard() {
  return (
    <Card sx={{ maxWidth: 600, margin: "auto", mt: 3 }}>
      <CardAuthor
        avatarSrc="/avatar.png"
        title="EVOL EDU"
        subheader="Marry Tran - Marketing Lead"
      />

      <CardMedia
        component="img"
        height="180"
        image="/images/posts/paella.jpg"
        alt="Future"
      />

      <CardContent>
        <Typography variant="h6">
          Làm Gì Để Nhân Viên Và Công Ty Bạn Không Tụt Hậu Trong Tương Lai
        </Typography>
        <Typography variant="body2" color="text.secondary" sx={{ mt: 1 }}>
          Tìm hiểu lý do vì sao mỗi người phải tự nâng cấp bản thân và học tập
          suốt đời.
        </Typography>
        <Chip
          label="ĐÓN TƯƠNG LAI"
          color="primary"
          size="small"
          sx={{ mt: 1 }}
        />
      </CardContent>

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

      <Box sx={{ px: 2, pb: 2 }}>
        <TextField
          fullWidth
          placeholder="Write your comment"
          variant="outlined"
          size="small"
        />
      </Box>
    </Card>
  );
}
