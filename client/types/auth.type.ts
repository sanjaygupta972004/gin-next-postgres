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

// eslint-disable-next-line
export function isUser(obj: any): obj is User {
  return !!obj && !!obj.email && !!obj.name && !!obj.role
}