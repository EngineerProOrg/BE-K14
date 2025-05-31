// File: src/pages/SignUp.tsx
import {
  Button,
  TextField,
  Box,
  Typography,
  Container,
  CssBaseline,
  Avatar,
  CircularProgress,
} from "@mui/material";
import { useForm, Controller } from "react-hook-form";
import LockOutlinedIcon from "@mui/icons-material/LockOutlined";
import { useState } from "react";
import { createTheme, ThemeProvider } from "@mui/material/styles";
import { useNavigate } from "react-router-dom";
import { SignUpRequestViewModel } from "../models/user";

import AdapterDateFns from "@mui/lab/AdapterDateFns";
import LocalizationProvider from "@mui/lab/LocalizationProvider";
import DatePicker from "@mui/lab/DatePicker";

import HttpClient from "../apis/HttpClient";

const theme = createTheme();

export default function SignUp() {
  const [isSubmitting, setSubmitting] = useState(false);
  const [signUpError, setSignUpError] = useState<string | null>(null);

  const {
    register,
    handleSubmit,
    control,
    formState: { errors },
  } = useForm<SignUpRequestViewModel>({
    defaultValues: {
      firstName: "",
      lastName: "",
      email: "",
      password: "",
      birthday: undefined,
    },
  });

  const navigate = useNavigate();

  const onSubmit = async (data: SignUpRequestViewModel) => {
    setSubmitting(true);
    try {
      await HttpClient.User.SignUp(data);
      navigate("/signin");
    } catch (err) {
      console.error(err);
      setSignUpError(String(err));
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
            display: "flex",
            flexDirection: "column",
            alignItems: "center",
          }}
        >
          <Avatar sx={{ m: 1, bgcolor: "secondary.main" }}>
            <LockOutlinedIcon />
          </Avatar>
          <Typography component="h1" variant="h5">
            Sign up
          </Typography>
          <Box
            component="form"
            onSubmit={handleSubmit(onSubmit)}
            sx={{ mt: 1 }}
          >
            <TextField
              fullWidth
              margin="normal"
              label="First name"
              {...register("firstName", {
                required: "First name is required",
                minLength: {
                  value: 2,
                  message: "First name must contain at least 2 characters",
                },
              })}
              error={!!errors.firstName}
              helperText={errors.firstName?.message}
            />
            <TextField
              fullWidth
              margin="normal"
              label="Last name"
              {...register("lastName", {
                required: "Last name is required",
                minLength: {
                  value: 2,
                  message: "Last name must contain at least 2 characters",
                },
              })}
              error={!!errors.lastName}
              helperText={errors.lastName?.message}
            />

            <LocalizationProvider dateAdapter={AdapterDateFns}>
              <Controller
                name="birthday"
                control={control}
                rules={{ required: "Birthday is required" }}
                render={({ field }) => (
                  <DatePicker
                    label="Birthday"
                    value={field.value ? new Date(field.value) : null}
                    onChange={(date: Date | null) => field.onChange(date)}
                    format="dd/MMM/yy"
                    renderInput={{
                      textField: {
                        margin: "normal",
                        fullWidth: true,
                        error: !!errors.birthday,
                        helperText: errors.birthday?.message,
                      },
                    }}
                  />
                )}
              />
            </LocalizationProvider>

            <TextField
              fullWidth
              margin="normal"
              label="Email"
              {...register("email", {
                required: "Email is required",
                pattern: {
                  value: /^[^@]+@[^@]+\.[^@]+$/,
                  message: "Invalid email",
                },
              })}
              error={!!errors.email}
              helperText={errors.email?.message}
            />
            <TextField
              fullWidth
              margin="normal"
              label="Password"
              type="password"
              {...register("password", {
                required: "Password is required",
                minLength: {
                  value: 6,
                  message: "Password must contain at least 6 characters",
                },
              })}
              error={!!errors.password}
              helperText={errors.password?.message}
            />
            {signUpError && (
              <Typography variant="body2" color="error" mt={1}>
                {signUpError}
              </Typography>
            )}
            <Button
              type="submit"
              fullWidth
              variant="contained"
              disabled={isSubmitting}
              sx={{ mt: 3, mb: 2 }}
              startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
            >
              {isSubmitting ? "Signing up..." : "Sign Up"}
            </Button>
          </Box>
        </Box>
      </Container>
    </ThemeProvider>
  );
}
