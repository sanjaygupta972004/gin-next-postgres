'use client'
import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { CookiesStorage } from '@/lib/storage/cookie';
import { AuthContextType, AuthCredentials, User } from '@/types/auth.type';
import { api_login } from '@/api/auth';
import { toast } from 'react-toastify';

const AuthContext = createContext<AuthContextType>({
  isLoading: false,
  user: null,
  login: () => { },
  logout: () => { },
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);

  useEffect(() => {
    const storedUser = CookiesStorage.getUser();
    if (storedUser) {
      setUser(JSON.parse(storedUser))
    }
    setIsLoading(false);
  }, [])

  const login = async (credentials: AuthCredentials) => {
    try {
      const res = (await api_login(credentials.email, credentials.password)).data;
      const token = res.token || null;
      CookiesStorage.setAccessToken(token);
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || "Invalid credentials");
      console.error("Failed to login", err);
    }
  }

  const logout = () => {
    setUser(null);
    CookiesStorage.clearCookieData("user");
    CookiesStorage.clearAccessToken();
  };

  return (
    <AuthContext.Provider value={{
      isLoading,
      user,
      login,
      logout
    }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext); 