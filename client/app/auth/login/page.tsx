"use client"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FaApple, FaAppStore, FaEnvelope, FaGithub, FaGoogle, FaKey, FaTwitter } from "react-icons/fa";
import { AuthCredentials } from "@/types/auth.type";
import { useAuth } from "@/context/AuthContext";
import withAuth from "@/components/hoc/withAuth";
import InputText from "@/components/common/InputBox";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { ROUTER } from "@/constants/common";
import { ThreeDots } from "react-loader-spinner";


const LoginPage: React.FC = () => {
  const [isLoggingIn, setIsLoggingIn] = useState<boolean>(false);
  const [credentials, setCredentials] = useState<AuthCredentials>({ email: "", password: "" })

  const { login } = useAuth();
  const router = useRouter();

  const onLogin = () => {
    try {
      setIsLoggingIn(true);
      login(credentials);
    } catch (err) {
      console.error(err);
    } finally {
      setIsLoggingIn(false);
    }
  }
  return (
    <Section className="w-[400px] m-auto">
      <p className="text-base text-center font-semibold">Login With</p>
      <div className="flex flex-row justify-center gap-6">
        <Button customClass="w-14 h-14 flex items-center justify-center" isPrimary={false}>
          <FaGoogle size={24} />
        </Button>
        <Button customClass="w-14 h-14 flex items-center justify-center" isPrimary={false}>
          <FaApple size={24} />
        </Button>
        <Button customClass="w-14 h-14 flex items-center justify-center" isPrimary={false}>
          <FaGithub size={24} />
        </Button>
        <Button customClass="w-14 h-14 flex items-center justify-center" isPrimary={false}>
          <FaTwitter size={24} />
        </Button>
      </div>
      <div className="w-full h-[1px] bg-zinc-800 my-2" />
      <InputText
        label="Email"
        labelIcon={<FaEnvelope />}
        value={credentials.email}
        onChange={(e) => setCredentials({ ...credentials, email: e.target.value })}
      />
      <InputText
        label="Password"
        labelIcon={<FaKey />}
        onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
        type="password"
        onKeyDown={(e) => { if (e.key === 'Enter') login(credentials) }}
      />
      <div className="flex justify-evenly gap-4 mt-4">
        <Button
          customClass="flex-1 flex items-center justify-center gap-2"
          onClick={onLogin}
        >
          {isLoggingIn && <ThreeDots color="#000" width={16} height={16} />}
          Login
        </Button>
        <Button
          customClass="flex-1"
          onClick={() => router.push(ROUTER.Register)}
          isPrimary={false}
        >
          Register
        </Button>
      </div>
    </Section >
  )
}

export default withAuth(LoginPage);