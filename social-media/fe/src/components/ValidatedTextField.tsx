import { TextField, TextFieldProps } from "@mui/material";
import { FieldError, UseFormRegisterReturn } from "react-hook-form";

type Props = {
  label: string;
  register: UseFormRegisterReturn;
  error?: FieldError;
} & Omit<TextFieldProps, "error" | "helperText">;

export default function ValidatedTextField({
  label,
  register,
  error,
  ...rest
}: Props) {
  return (
    <TextField
      label={label}
      {...register}
      error={!!error}
      helperText={error?.message}
      fullWidth
      {...rest}
    />
  );
}
