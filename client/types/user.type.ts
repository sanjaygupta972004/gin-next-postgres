export type User = {
    userID?: string;
    fullName: string;
    gender: string;
    email: string;
    role?: string;
    username: string;
    profileImage?: string;
    bannerImage?: string;
    createdAt?: string;
    updatedAt?: string;
    isEmailVerified?: boolean;
}

// eslint-disable-next-line
export function isUser(obj: any): obj is User {
    return !!obj && !!obj.email && !!obj.fullName && !!obj.role
}

