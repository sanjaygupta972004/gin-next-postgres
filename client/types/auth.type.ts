import { User } from "./user.type";

export interface AuthCredentials {
  email: string;
  password: string;
}

export type AuthRegistrationFormRequest = {
  fullName: string;
  username: string;
  email: string;
  gender: string;
  password: string;
  role: string;
  confirmPassword: string;
}

export interface AuthContextType {
  isLoading: boolean;
  user: User | null,
  login: (credentials: AuthCredentials) => Promise<void>;
  register: (request: AuthRegistrationFormRequest) => Promise<User | null>;
  logout: () => void;
}
