import { Box, Typography } from "@mui/material";

function Header() {
  return (
    <Box component="header" sx={{ p: 2, bgcolor: 'primary.main', color: 'white' }}>
      <Typography variant="h6">My Social App</Typography>
    </Box>
  );
}

export default Header