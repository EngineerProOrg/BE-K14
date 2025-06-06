import { Box, TextField } from "@mui/material";

function PostComment() {
  return (
    <Box sx={{ px: 2, pb: 2 }}>
      <TextField
        fullWidth
        placeholder="Write your comment"
        variant="outlined"
        size="small"
      />
    </Box>
  );
}

export default PostComment;
