import { createContext, useContext, useEffect, useState } from "react";
import {
  AuthContextViewModel,
  SignInResponseViewModel,
  UserBaseViewModel,
} from "../models/user";

const AuthContext = createContext<AuthContextViewModel | null>(null);

export const AuthProvider: React.FC<React.PropsWithChildren<{}>> = ({
  children,
}) => {
  const [userInfo, setUserInfo] = useState<UserBaseViewModel | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);

  // Get access token from local storage
  useEffect(() => {
    const tokenInStorage = localStorage.getItem("accessToken");
    if (tokenInStorage) {
      setAccessToken(tokenInStorage);
    }
  }, []);

  const signIn = (signinResponseViewModel: SignInResponseViewModel) => {
    if (signinResponseViewModel) {
      setAccessToken(signinResponseViewModel.accessToken);
      setUserInfo(signinResponseViewModel.userInfo);
      localStorage.setItem("accessToken", signinResponseViewModel.accessToken);
    }
  };

  const signOut = () => {
    setUserInfo(null);
    setAccessToken(null);
    localStorage.removeItem("accessToken");
  };

  return (
    <AuthContext.Provider
      value={{ user: userInfo, token: accessToken, signIn, signOut }}
    >
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth must be used within AuthProvider");
  return context;
};

export const isValidToken = (token: string): boolean => {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    return payload.exp * 1000 > Date.now();
  } catch {
    return false;
  }
};
