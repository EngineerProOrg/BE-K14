export interface AuthContextViewModel {
  user: UserInfo | null;
  token: string | null;
  signIn: (data: SignInResponseViewModel) => void;
  signOut: () => void;
}

export interface SignInResponseViewModel {
  userInfo: UserInfo;
  accessToken: string;
}

export interface UserInfo {
  userId: number;
  firstName: string;
  lastName: string;
  name: string;
  birthday: Date;
  email: string;
  username: string;
  avatar: string;
}

export interface SignInRequestViewModel {
  email: string;
  password: string;
}

export interface SignUpRequestViewModel {
  firstName: string;
  lastName: string;
  birthday: Date;
  email: string;
  password: string;
  avatar: string;
}
