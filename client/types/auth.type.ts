import { User } from "./user.type";

export interface AuthCredentials {
  email: string;
  password: string;
}

export interface AuthContextType {
  isLoading: boolean;
  user: User | null,
  login: (credentials: AuthCredentials) => void;
  logout: () => void;
}
