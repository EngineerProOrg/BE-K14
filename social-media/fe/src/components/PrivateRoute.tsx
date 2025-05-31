import { Navigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export default function PrivateRoute({ children }: { children: JSX.Element }) {
  const { token } = useAuth();

  return token ? children : <Navigate to="/signin" replace />;
}
