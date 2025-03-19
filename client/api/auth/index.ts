import { API_AUTH_LOGIN } from "@/constants/endpoint"
import api from ".."

export const api_login = (email: string, password: string) => {
    return api.post(API_AUTH_LOGIN, {email, password})
}