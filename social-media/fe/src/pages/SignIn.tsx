import { useForm } from "react-hook-form";
import { Button, Box, Typography, CircularProgress } from "@mui/material";
import axios from "../apis/HttpClient";
import { useNavigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";
import {
  SignInRequestViewModel,
  SignInResponseViewModel,
} from "../models/user";
import ValidatedTextField from "../components/ValidatedTextField";
import { useState } from "react";
import FormErrorMessage from "../components/FormErrorMessage";
import HttpClient from "../apis/HttpClient";

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
      const response = await HttpClient.Auth.SignIn(signinRequestViewModel);

      signIn(response);
      navigate("/feed");
    } catch (err) {
      console.error("Login failed:", err);
      setSignInError("Signin failed! Email or password incorrect.");
    } finally {
      setSubmitting(false);
    }
  };

  return (
    <Box sx={{ maxWidth: 400, mx: "auto", mt: 10 }}>
      <Typography variant="h5" mb={3}>
        Sign In
      </Typography>
      <form onSubmit={handleSubmit(onSubmit)}>
        <ValidatedTextField
          label="Email"
          register={register("email", {
            required: "Email is required.",
            pattern: {
              value: /^[^@]+@[^@]+\.[^@]+$/,
              message: "Invalid email",
            },
          })}
          error={errors.email}
        />

        <ValidatedTextField
          label="Password"
          type="password"
          register={register("password", {
            required: "Password is required",
            minLength: {
              value: 6,
              message: "Password must contains at least 6 characters.",
            },
          })}
          error={errors.password}
        />

        <Button
          type="submit"
          variant="contained"
          fullWidth
          sx={{ mt: 2 }}
          disabled={isSubmitting}
          startIcon={isSubmitting ? <CircularProgress size={20} /> : null}
        >
          {isSubmitting ? "Sign in..." : "SIGN IN"}
        </Button>

        {signInError && <FormErrorMessage message={signInError} />}
      </form>
    </Box>
  );
}
