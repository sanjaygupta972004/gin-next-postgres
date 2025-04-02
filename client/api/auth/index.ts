import { API_AUTH_GETME, API_AUTH_LOGIN, API_REFRESH_TOKEN, API_UPDATE_PROFILE } from "@/constants/endpoint";
import api from "..";
import { User } from "@/types/user.type";

export const api_login = (email: string, password: string) => {
    return api.post(API_AUTH_LOGIN, { email, password });
}

export const api_getme = () => {
    return api.get(API_AUTH_GETME);
}

export const api_refresh_token = () => {
    return api.get(API_REFRESH_TOKEN);
}

export const api_update_profile = (user: User) => {
    return api.post(API_UPDATE_PROFILE, { ...user })
}