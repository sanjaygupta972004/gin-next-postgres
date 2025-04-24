import {
    API_USER_GET_PROFILE,
    API_USER_UPDATE_PROFILE,
    API_USER_UPLOAD_AVATAR,
    API_USER_UPLOAD_BANNER
} from "@/constants/endpoint";
import { User } from "@/types/user.type";
import api from "..";

export const api_user_get_profile = () => {
    return api.get(API_USER_GET_PROFILE);
}

export const api_user_update_profile = (user: User) => {
    return api.patch(API_USER_UPDATE_PROFILE, user)
}

export const api_user_upload_banner = (data: FormData) => {
    return api.patch(API_USER_UPLOAD_BANNER, data);
}

export const api_user_upload_avatar = (data: FormData) => {
    return api.patch(API_USER_UPLOAD_AVATAR, data);
}