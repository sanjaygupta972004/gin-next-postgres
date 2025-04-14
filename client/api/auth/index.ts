import {
    API_USER_LOGIN,
    API_USER_REGISTER,
    API_USER_OTP_RESEND,
    API_USER_EMAIL_VERIFICATION,
    API_USER_PROFILE,
    API_UPDATE_PROFILE,
    API_REFRESH_TOKEN,
} from "@/constants/endpoint";
import api from "..";
import { User } from "@/types/user.type";
import { AuthRegistrationFormRequest } from "@/types/auth.type";

export const api_user_login = (email: string, password: string) => {
    return api.post(API_USER_LOGIN, { email, password });
}

export const api_user_register = (request: AuthRegistrationFormRequest) => {
    return api.post(API_USER_REGISTER, { ...request });
}

export const api_user_resend_otp_code = (userID: string) => {
    return api.post(`${API_USER_OTP_RESEND}/${userID}`);
}

export const api_user_email_verify = (userID: string, otpCode: number) => {
    return api.post(`${API_USER_EMAIL_VERIFICATION}/${userID}`, { authOtp: otpCode });
}

export const api_user_profile = () => {
    return api.get(API_USER_PROFILE);
}

export const api_update_profile = (user: User) => {
    return api.post(API_UPDATE_PROFILE, { ...user })
}

export const api_refresh_token = () => {
    return api.get(API_REFRESH_TOKEN);
}
