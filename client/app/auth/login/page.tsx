"use client"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FaEnvelope, FaKey } from "react-icons/fa";
import { AuthCredentials } from "@/types/auth.type";
import { useAuth } from "@/context/AuthContext";
import withAuth from "@/components/hoc/withAuth";
import InputText from "@/components/common/InputBox";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { ROUTER } from "@/constants/common";


const LoginPage: React.FC = () => {
  const [credentials, setCredentials] = useState<AuthCredentials>({ email: "", password: "" })

  const { login } = useAuth();
  const router = useRouter();

  return (
    <Section className="w-[400px] m-auto">
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
          customClass="flex-1"
          onClick={() => login(credentials)}
        >
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