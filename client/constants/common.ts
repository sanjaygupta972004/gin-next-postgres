
export const CookieKey = {
  accessToken: 'jwt',
  user: 'user',
};

export const ROUTER = {
  Home: '/',
  Login: '/auth/login',
  Register: '/auth/register',
  Verification: (userID: string) => `/auth/verification/${userID}`,
  Profile: '/profile',
  Forbidden: '/auth/forbidden',
  Users: '/users',
}