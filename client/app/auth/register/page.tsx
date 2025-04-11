"use client"
import { useRouter } from "next/navigation";
import { useState } from "react";
import { FaEnvelope, FaLock, FaUnlock, FaUser, FaUserTag } from "react-icons/fa";
import { AuthRegisterRequest } from "@/types/auth.type";
import InputText from "@/components/common/InputBox";
import Section from "@/components/common/Section";
import { Button } from "@/components/common/Button";
import { ROUTER } from "@/constants/common";
import { useAuth } from "@/context/AuthContext";


const LoginPage: React.FC = () => {

  const { register } = useAuth();
  const router = useRouter();

  const [request, setRequest] = useState<AuthRegisterRequest>({
    fullName: "",
    username: "",
    email: "",
    password: "",
    gender: "male",
    role: "user",
  })

  return (
    <Section className="w-[400px] m-auto">
      <InputText
        label="Full Name"
        labelIcon={<FaUser />}
        onChange={(e) => setRequest({ ...request, fullName: e.target.value })}
      />
      <InputText
        label="Username"
        labelIcon={<FaUserTag />}
        onChange={(e) => setRequest({ ...request, username: e.target.value })}
      />
      <InputText
        label="Email"
        labelIcon={<FaEnvelope />}
        value={request.email}
        onChange={(e) => setRequest({ ...request, email: e.target.value })}
        type="email"
      />
      <InputText
        label="Password"
        labelIcon={<FaLock />}
        onChange={(e) => setRequest({ ...request, password: e.target.value })}
        type="password"
      />
      <InputText
        label="Confirm Password"
        labelIcon={<FaUnlock />}
        onChange={(e) => setRequest({ ...request, password: e.target.value })}
        type="password"
      />
      <div className="flex justify-evenly gap-4 mt-4">
        <Button
          customClass="w-30"
          onClick={() => register(request)}
        >
          Reigster
        </Button>
        <Button
          customClass="text-nowrap"
          onClick={() => router.push(ROUTER.Login)}
          isPrimary={false}
        >
          Already have account?
        </Button>
      </div>
    </Section >
  )
}

export default LoginPage;