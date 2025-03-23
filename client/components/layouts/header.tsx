"use client";
import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image"
import { ROUTER } from "@/constants/common";
import { FaSignInAlt, FaSignOutAlt, FaUserCircle, FaUserTie } from "react-icons/fa";
import { useAuth } from "@/context/AuthContext";

export default function Header() {
  const router = useRouter();
  const { user, logout } = useAuth();

  const [isShowUserDropdown, setIsShowUserDropdown] = useState<boolean>(false);
  const dropdownRef = useRef<HTMLDivElement | null>(null);


  useEffect(() => {
    const handleOutSideClick = (event: MouseEvent) => {
      if (!dropdownRef.current) return;
      const target = event.target as Node;
      if (isShowUserDropdown && dropdownRef.current && !dropdownRef.current.contains(target)) {
        setIsShowUserDropdown(false)
      }
    }

    document.addEventListener("mousedown", handleOutSideClick);
    return () => {
      document.removeEventListener("mousedown", handleOutSideClick);
    };
  }, [isShowUserDropdown]);

  return (
    <header className="border-b border-dashed border-b-zinc-800 py-2 px-8 flex justify-between items-center">
      <Link href="/" >
        <Image src="/logo.png" alt="logo" width={71} height={48} className="py-1.5 h-12 w-auto object-contain" />
      </Link>
      {user ?
        <div className="relative">
          <p
            className="flex items-center gap-2 font-semibold cursor-pointer"
            onClick={() => setIsShowUserDropdown(!isShowUserDropdown)}
          >
            <FaUserTie size={18} />{user.name}
          </p>
          {isShowUserDropdown &&
            <div
              ref={dropdownRef}
              className="flex flex-col gap-4 absolute translate-y-2 top-full right-0 px-6 py-3 border border-solid border-zinc-800 rounded-lg bg-zinc-900/10 backdrop-blur-2xl"
            >
              <a
                href={ROUTER.Profile}
                className="flex items-center gap-2 font-semibold cursor-pointer"
              >
                <FaUserCircle size={18} />Profile
              </a>
              <div
                className="flex items-center gap-2 font-semibold cursor-pointer"
                onClick={() => { logout(); setIsShowUserDropdown(false) }}
              >
                <FaSignOutAlt size={18} />Logout
              </div>
            </div>
          }
        </div>
        : <button
          className="py-2 px-4 rounded-lg bg-white text-black flex items-center gap-2 cursor-pointer font-medium"
          onClick={() => router.push(ROUTER.Login)}
        >
          <FaSignInAlt />Login
        </button>
      }
    </header>
  )
}