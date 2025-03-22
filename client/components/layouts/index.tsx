"use client";
import { ToastContainer } from "react-toastify";
import Header from "./header";

export default function Layout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {

  return (
    <>
      <div className="w-screen min-h-screen flex flex-col">
        <Header />
        <div className="flex-1 flex flex-col max-w-[1440px] w-full m-auto p-8">
          {children}
        </div>
        <footer className="border-t border-dashed border-t-zinc-800 py-4">
          <p className="text-sm text-center font-semibold">Copyright</p>
        </footer>
      </div>
      <ToastContainer theme="dark" />
    </>
  )
}