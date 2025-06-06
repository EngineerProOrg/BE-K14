import { Box, Typography } from "@mui/material";


function LeftSidebar() {
  return (
    <Box sx={{ p: 2 }}>
      <Typography variant="h6">Navigation</Typography>
      <ul>
        <li>Home</li>
        <li>Profile</li>
        <li>Settings</li>
      </ul>
    </Box>
  );
}

export default LeftSidebar