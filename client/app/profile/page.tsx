"use client";
import withAuth from "@/components/hoc/withAuth";
import { useAuth } from "@/context/AuthContext";

const ProfilePage: React.FC = () => {
  const { user } = useAuth();

  return (
    <div className="flex flex-col gap-8 rounded-lg p-8 border border-solid border-zinc-800">
      <p>
        Name: <b>{user?.name}</b>
      </p>
      <p>
        Email:<b> {user?.email}</b>
      </p>
      <p>
        Role: <b>{user?.role}</b>
      </p>
    </div>
  );
}


export default withAuth(ProfilePage);