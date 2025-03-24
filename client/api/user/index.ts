import { API_GET_USERS } from "@/constants/endpoint";
import api from "..";

export const api_get_users = () => {
    return api.get(API_GET_USERS);
}