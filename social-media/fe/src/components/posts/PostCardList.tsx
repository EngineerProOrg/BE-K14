import { Box, Paper, Typography } from "@mui/material";
import PostCard from "./PostCard";

function Main() {
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h5">News Feed</Typography>
      <Paper sx={{ p: 2, mt: 2 }}>
        <PostCard />
      </Paper>
      <Paper sx={{ p: 2, mt: 2 }}>
        <Typography variant="body1">Post 2 - Welcome back!</Typography>
      </Paper>
    </Box>
  );
}

export default Main;
