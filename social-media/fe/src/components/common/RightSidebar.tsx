import { Box, Typography } from "@mui/material";

function RightSidebar() {
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h6">Trending</Typography>
      <ul>
        <li>#AI</li>
        <li>#React</li>
        <li>#GoLang</li>
      </ul>
    </Box>
  );
}

export default RightSidebar