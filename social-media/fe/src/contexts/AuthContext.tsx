import { createContext, useContext, useState } from "react";
import {
  AuthContextViewModel,
  SignInResponseViewModel,
  UserInfo,
} from "../models/user";

const AuthContext = createContext<AuthContextViewModel | null>(null);

export const AuthProvider: React.FC<React.PropsWithChildren<{}>> = ({
  children,
}) => {
  const [user, setUser] = useState<UserInfo | null>(null);
  const [token, setToken] = useState<string | null>(null);

  const signIn = (signinResponseViewModel: SignInResponseViewModel) => {
    setToken(signinResponseViewModel.accessToken);
    setUser(signinResponseViewModel.userInfo);
    localStorage.setItem("accessToken", signinResponseViewModel.accessToken);
  };

  const signOut = () => {
    setUser(null);
    setToken(null);
    localStorage.removeItem("accessToken");
  };

  return (
    <AuthContext.Provider value={{ user, token, signIn, signOut }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
};
