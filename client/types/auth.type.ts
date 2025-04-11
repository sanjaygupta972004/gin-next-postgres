import { User } from "./user.type";

export interface AuthCredentials {
  email: string;
  password: string;
}

export interface AuthRegisterRequest {
  fullName: string;
  username: string;
  gender: string;
  email: string;
  password: string;
  role: string;
}

export interface AuthContextType {
  isLoading: boolean;
  user: User | null,
  login: (credentials: AuthCredentials) => void;
  register: (request: AuthRegisterRequest) => void;
  logout: () => void;
}
