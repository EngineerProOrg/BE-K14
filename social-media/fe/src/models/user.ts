export interface AuthContextViewModel {
  user: UserBaseViewModel | null;
  token: string | null;
  signIn: (data: SignInResponseViewModel) => void;
  signOut: () => void;
}

export interface SignInResponseViewModel {
  userInfo: UserBaseViewModel;
  accessToken: string;
}

export interface UserBaseViewModel {
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
  birthday: Date | undefined;
  email: string;
  password: string;
  avatar: string;
}
