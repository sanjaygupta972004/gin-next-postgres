import { CookieKey } from "@/constants/common";

export const CookiesStorage = {
  getCookieData(key: string) {
    return localStorage.getItem(key);
  },

  setCookieData(key: string, data: string) {
    localStorage.setItem(key, data);
  },

  clearCookieData(key: string) {
    localStorage.removeItem(key);
  },

  getAccessToken() {
    return localStorage.getItem(CookieKey.accessToken)
  },
  
  setAccessToken(accessToken: string) {
    localStorage.setItem(CookieKey.accessToken, accessToken)
  }
}