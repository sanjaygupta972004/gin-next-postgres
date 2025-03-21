"use client";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image"
import { ROUTER } from "@/constants/common";
import { FaSignInAlt } from "react-icons/fa";

export default function Header() {
  const router = useRouter()

  return (
    <header className="border-b border-dashed border-b-zinc-800 py-2 px-8 flex justify-between items-center">
      <Link href="/" >
        <Image src="/logo.png" alt="logo" width={71} height={48} className="py-1.5 h-12 w-auto object-contain" />
      </Link>
      <button
        className="py-2 px-4 rounded-lg bg-white text-black flex items-center gap-2 cursor-pointer font-medium"
        onClick={() => router.push(ROUTER.Login)}
      >
        <FaSignInAlt />Login
      </button>
    </header>
  )
}