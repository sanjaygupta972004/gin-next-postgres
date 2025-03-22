"use client"
import { useState } from "react";
import { FaEnvelope, FaKey } from "react-icons/fa";
import { AuthCredentials } from "@/types/auth.type";
import { useAuth } from "@/context/AuthContext";
import withAuth from "@/components/hoc/withAuth";


const LoginPage: React.FC = () => {
  const [credentials, setCredentials] = useState<AuthCredentials>({ email: "", password: "" })

  const { login } = useAuth();

  return (

    <div className="h-full w-[400px] flex flex-col gap-6 justify-center items-center m-auto border border-zinc-800 rounded-lg p-8">
      <div className="w-full flex flex-col gap-2">
        <p className="flex items-center gap-2 font-semibold"><FaEnvelope />Email</p>
        <input
          className="w-full border border-zinc-800 rounded-lg py-2 px-4"
          value={credentials.email}
          onChange={(e) => setCredentials({ ...credentials, email: e.target.value })}
        />
      </div>
      <div className="w-full flex flex-col gap-2">
        <p className="flex items-center gap-2 font-semibold"><FaKey />Password</p>
        <input
          className="w-full border border-zinc-800 rounded-lg py-2 px-4"
          value={credentials.password}
          onChange={(e) => setCredentials({ ...credentials, password: e.target.value })}
          type="password"
        />
      </div>
      <div className="flex gap-4">
        <button
          className="bg-white text-zinc-950 font-semibold py-2 px-4 rounded-lg cursor-pointer"
          onClick={() => login(credentials)}
        >
          Login
        </button>
        <button className="font-semibold py-2 px-4 rounded-lg cursor-pointer">
          Register
        </button>
      </div>
    </div>
  )
}

export default withAuth(LoginPage);