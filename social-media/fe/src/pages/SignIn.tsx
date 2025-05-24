import { useForm } from "react-hook-form";
import { Button, Box, Typography, CircularProgress,  Container, CssBaseline, Avatar, TextField } from "@mui/material";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import {
  SignInRequestViewModel,
} from "../models/user";
import LockOutlinedIcon from '@mui/icons-material/LockOutlined';
import { useState } from "react";
import HttpClient from "../apis/HttpClient";
import { createTheme, ThemeProvider } from '@mui/material/styles';


const theme = createTheme();

export default function SignIn() {
  const [isSubmitting, setSubmitting] = useState(false);
  const [signInError, setSignInError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm<SignInRequestViewModel>();

  const { signIn } = useAuth();
  const navigate = useNavigate();

  const onSubmit = async (signinRequestViewModel: SignInRequestViewModel) => {
    setSubmitting(true);
    try {
      const response = await HttpClient.User.SignIn(signinRequestViewModel);

      signIn(response);
     
    } catch (err) {
      setSignInError(String(err));
      console.error(err)
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <ThemeProvider theme={theme}>
      <Container component="main" maxWidth="xs">
        <CssBaseline />
        <Box
          sx={{
            marginTop: 8,
            display: 'flex',
            flexDirection: 'column',
            alignItems: 'center',
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: 'secondary.main' }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign in
          </Typography>
          <Box component="form" onSubmit={handleSubmit(onSubmit)} sx={{ mt: 1 }}>
            <TextField
              margin="normal"
              fullWidth
              label="Email Address"
              autoComplete="email"
              autoFocus
              error={!!errors.email}
              helperText={errors.email?.message}
              {...register("email", {
                required: "Email is required",
                pattern: {
                  value: /^[^@]+@[^@]+\.[^@]+$/,
                  message: "Invalid email",
                },
              })}
            />
            <TextField
              margin="normal"
              fullWidth
              label="Password"
              type="password"
              autoComplete="current-password"
              error={!!errors.password}
              helperText={errors.password?.message}
              {...register("password", {
                required: "Password is required",
                minLength: { value: 6, message: "Password must contains at least 6 characters" },
              })}
            />
            {signInError && (
              <Typography variant="body2" color="error" mt={1}>
                {signInError}
              </Typography>
            )}
            <Button
              type="submit"
              fullWidth
              variant="contained"
              sx={{ mt: 3, mb: 2 }}
              disabled={isSubmitting}
              startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
            >
              {isSubmitting ? 'Signin...' : 'Sign In'}
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
