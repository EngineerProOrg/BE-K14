import { Navigate } from "react-router-dom";
import { isValidToken, useAuth } from "../contexts/AuthContext";

export default function PrivateRoute({ children }: { children: JSX.Element }) {
  const { token } = useAuth();

  return token && isValidToken(token) ? (
    children
  ) : (
    <Navigate to="/signin" replace />
  );
}
