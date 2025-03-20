export interface User {
    email: string;
    name: string;
    role: string;
}

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
