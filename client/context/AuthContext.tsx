'use client'
import { useRouter } from 'next/navigation';
import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { toast } from 'react-toastify';
import { isUser, User } from '@/types/user.type';
import { AuthContextType, AuthCredentials, AuthRegisterRequest } from '@/types/auth.type';
import { ROUTER } from '@/constants/common';
import { CookiesStorage } from '@/lib/storage/cookie';
import { api_getme, api_login, api_refresh_token, api_user_register } from '@/api/auth';

const AuthContext = createContext<AuthContextType>({
  isLoading: false,
  user: null,
  login: () => { },
  register: () => { },
  logout: () => { },
});

export function AuthProvider({ children }: { children: ReactNode }) {
  const [user, setUser] = useState<User | null>(null);
  const [isLoading, setIsLoading] = useState<boolean>(true);
  const router = useRouter();

  useEffect(() => {
    const getUserInfo = async () => {
      try {
        setIsLoading(true);
        if (CookiesStorage.getAccessToken() !== null) {
          const refreshToken = (await api_refresh_token()).data;
          CookiesStorage.setAccessToken(refreshToken.token);

          const me = (await api_getme()).data;
          CookiesStorage.setUser(me);
          setUser(me);
        }
      } catch (err) {
        console.error(err);
        setUser(null);
        CookiesStorage.clearUser();
        CookiesStorage.clearAccessToken();
      } finally {
        setIsLoading(false)
      }
    }
    getUserInfo();
  }, [])

  const login = async (credentials: AuthCredentials) => {
    try {
      const res = (await api_login(credentials.email, credentials.password)).data;
      const token = res.token || null;
      CookiesStorage.setAccessToken(token);

      const me = (await api_getme()).data;
      if (!isUser(me))
        setUser(null);
      else
        setUser(me as User);
      CookiesStorage.setUser(me as User);
      router.push(ROUTER.Home)
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || "Invalid credentials");
      console.error("Failed to login", err);
    }
  }

  const register = async (request: AuthRegisterRequest) => {
    try {
      const res = (await api_user_register(request)).data;
      toast.success("Successfully registered!");
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || "Invalid credentials");
      console.error("Failed to register", err);
    }
  }

  const logout = () => {
    setUser(null);
    CookiesStorage.clearCookieData("user");
    CookiesStorage.clearAccessToken();
    router.push(ROUTER.Home)
  };

  return (
    <AuthContext.Provider value={{
      isLoading,
      user,
      login,
      logout,
      register
    }}>
      {children}
    </AuthContext.Provider>
  );
}

export const useAuth = () => useContext(AuthContext); 