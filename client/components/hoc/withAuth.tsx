"use client";
import { ROUTER } from "@/constants/common";
import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

const withAuth = (WrappedComponent: React.FC): React.FC => {
  const AuthenticatedComponent: React.FC = (props) => {
    const { isLoading, user } = useAuth();
    const router = useRouter();

    useEffect(() => {
      if (isLoading) return;
      if (!user)
        router.push(ROUTER.Login);
      else
        router.push(ROUTER.Home);
    }, [router, isLoading, user]);

    return <WrappedComponent {...props} />;
  };

  return AuthenticatedComponent;
};

export default withAuth;