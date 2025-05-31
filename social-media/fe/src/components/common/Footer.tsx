import { Box, Typography } from "@mui/material";

export default function Footer() {
  return (
    <Box component="footer" sx={{ p: 2, bgcolor: "grey.200", mt: 4 }}>
      <Typography variant="body2" color="text.secondary" align="center">
        © 2025 My Social App. All rights reserved.
      </Typography>
    </Box>
  );
}
