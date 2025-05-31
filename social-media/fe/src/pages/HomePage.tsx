import { Navigate } from "react-router-dom";
import { useAuth } from "../contexts/AuthContext";

export default function HomeRedirect() {
  const { token } = useAuth();

  return <Navigate to={token ? "/posts" : "/signin"} replace />;
}
