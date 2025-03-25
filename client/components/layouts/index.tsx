"use client";
import { ToastContainer } from "react-toastify";
import Header from "./header";
import { useAuth } from "@/context/AuthContext";
import { InfinitySpin } from "react-loader-spinner";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const { isLoading } = useAuth();
  return (
    <>
      <div className="w-screen min-h-screen flex flex-col">
        <Header />
        <div className="flex-1 flex flex-col max-w-[1440px] w-full m-auto p-8">
          {isLoading ?
            <div className="flex justify-center items-center m-auto">
              <InfinitySpin color="#FFF" />
            </div>
            : children
          }
        </div>
        <footer className="border-t border-dashed border-t-zinc-800 py-4">
          <p className="text-sm text-center font-semibold">Copyright © 2025</p>
        </footer>
      </div>
      <ToastContainer theme="dark" />
    </>
  )
}