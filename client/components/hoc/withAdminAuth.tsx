"use client";
import { ROUTER } from "@/constants/common";
import { useAuth } from "@/context/AuthContext";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

const withAdminAuth = (WrappedComponent: React.FC): React.FC => {
  const AuthenticatedComponent: React.FC = (props) => {
    const { isLoading, user } = useAuth();
    const router = useRouter();

    useEffect(() => {
      if (isLoading) return;
      if (user?.role !== 'admin')
        router.push(ROUTER.Forbidden);
    }, [router, isLoading, user]);

    return <WrappedComponent {...props} />;
  };

  return AuthenticatedComponent;
};

export default withAdminAuth;