"use client";

import { useEffect, useState } from "react";
import { api_get_users } from "@/api/user";
import withAdminAuth from "@/components/hoc/withAdminAuth";
import { User } from "@/types/user.type";

function UsersPage() {
  const [users, setUsers] = useState<User[]>([]);
  useEffect(() => {
    const getUsers = async () => {
      try {
        const _users = (await api_get_users()).data.users as User[];
        setUsers(_users)
      } catch (err) {
        console.error(err);
        setUsers([]);
      }
    }
    getUsers();
  }, [])
  return (
    <div>
      <h1 className="text-center text-xl"><b>Users list</b></h1>
      <div className="mt-8">
        {users.map((user, index) =>
          <div
            className="text-sm flex items-center gap-8"
            key={user.email}
          >
            <p>{index + 1}.</p>
            <p>{user.name}</p>
            <p>{user.email}</p>
            {user.role === 'admin'
              ? <div className="flex items-center bg-sky-600 px-2 py-1 rounded-lg text-white font-semibold">Admin</div>
              : <div className="flex items-center bg-green-600 px-2 py-1 rounded-lg text-white font-semibold">User</div>
            }
          </div>
        )}
      </div>
    </div>
  )
}

export default withAdminAuth(UsersPage);