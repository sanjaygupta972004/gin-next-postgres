"use client"
import { useState } from "react";
import { FaEnvelope, FaKey } from "react-icons/fa";
import { api_login } from "@/api/auth";
import { toast } from "react-toastify";

interface LoginCredentials {
  email: string;
  password: string;
}

export default function LoginPage() {
  const [loginData, setLoginData] = useState<LoginCredentials>({ email: "", password: "" })

  const onLogin = async () => {
    try {
      await api_login(loginData.email, loginData.password);
    } catch (err) {
      // eslint-disable-next-line
      toast.error((err as any)?.message || 'Invalid credentials');
      console.error("Failed to login", err);
    }
  }
  return (

    <div className="h-full w-[400px] flex flex-col gap-6 justify-center items-center m-auto border border-zinc-800 rounded-lg p-8">
      <div className="w-full flex flex-col gap-2">
        <p className="flex items-center gap-2 font-semibold"><FaEnvelope />Email</p>
        <input
          className="w-full border border-zinc-800 rounded-lg py-2 px-4"
          value={loginData.email}
          onChange={(e) => setLoginData({ ...loginData, email: e.target.value })}
        />
      </div>
      <div className="w-full flex flex-col gap-2">
        <p className="flex items-center gap-2 font-semibold"><FaKey />Password</p>
        <input
          className="w-full border border-zinc-800 rounded-lg py-2 px-4"
          value={loginData.password}
          onChange={(e) => setLoginData({ ...loginData, password: e.target.value })}
          type="password"
        />
      </div>
      <div className="flex gap-4">
        <button
          className="bg-white text-zinc-950 font-semibold py-2 px-4 rounded-lg cursor-pointer"
          onClick={onLogin}
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