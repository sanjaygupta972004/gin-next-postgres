export type User = {
    email: string;
    name: string;
    role?: string;
}

// eslint-disable-next-line
export function isUser(obj: any): obj is User {
    return !!obj && !!obj.email && !!obj.name && !!obj.role
}

