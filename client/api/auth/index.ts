import {
    API_AUTH_EMAIL_VERIFICATION,
    API_AUTH_LOGIN,
    API_AUTH_OTP_RESEND,
    API_AUTH_REFRESH_TOKEN,
    API_AUTH_REGISTER
} from "@/constants/endpoint";
import { AuthRegistrationFormRequest } from "@/types/auth.type";
import api from "..";

export const api_auth_login = (email: string, password: string) => {
    return api.post(API_AUTH_LOGIN, { email, password });
}

export const api_auth_register = (request: AuthRegistrationFormRequest) => {
    return api.post(API_AUTH_REGISTER, { ...request });
}

export const api_auth_resend_otp_code = (userID: string) => {
    return api.post(`${API_AUTH_OTP_RESEND}/${userID}`);
}

export const api_auth_email_verify = (userID: string, otpCode: number) => {
    return api.post(`${API_AUTH_EMAIL_VERIFICATION}/${userID}`, { authOtp: otpCode });
}

export const api_auth_refresh_token = () => {
    return api.get(API_AUTH_REFRESH_TOKEN);
}
