'use client'
import { useRouter } from 'next/navigation';
import { createContext, useContext, useState, useEffect, ReactNode } from 'react';
import { toast } from 'react-toastify';
import { isUser, User } from '@/types/user.type';
import { AuthContextType, AuthCredentials, AuthRegistrationFormRequest } from '@/types/auth.type';
import { ROUTER } from '@/constants/common';
import { CookiesStorage } from '@/lib/storage/cookie';
import { api_user_profile, api_user_login, api_refresh_token, api_user_register } from '@/api/auth';

const AuthContext = createContext<AuthContextType>({
  isLoading: false,
  user: null,
  login: async () => { },
  register: async () => { return null },
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
          const me = (await api_user_profile()).data.data;
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
      const res = (await api_user_login(credentials.email, credentials.password)).data.data;

      const me = res.data as User;
      if (!me.isEmailVerified) {
        router.push(ROUTER.Verification(me.userID!))
        return;
      }
      const accessToken = res.accessToken;
      CookiesStorage.setAccessToken(accessToken);

      if (!isUser(me))
        setUser(null);
      else
        setUser(me as User);
      setUser(me);
      CookiesStorage.setUser(me as User);
      router.push(ROUTER.Home)
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || "Invalid credentials");
      console.error("Failed to login", err);
    }
  }

  const register = async (request: AuthRegistrationFormRequest) => {
    try {
      const user = (await api_user_register(request)).data.data as User;
      toast.success("Successfully registered!");
      return user;
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || "Invalid credentials");
      console.error("Failed to register", err);
      return null;
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