import { Typography } from "@mui/material";

export default function FormErrorMessage({ message }: { message: string }) {
  return (
    <Typography variant="body2" color="error" mb={2}>
      {message}
    </Typography>
  );
}

 