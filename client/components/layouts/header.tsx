"use client";
import { useEffect, useRef, useState } from "react";
import { useRouter } from "next/navigation";
import Link from "next/link";
import Image from "next/image"
import { ROUTER } from "@/constants/common";
import { FaBook, FaComments, FaSignInAlt, FaSignOutAlt, FaTrophy, FaUserCircle, FaUserTie } from "react-icons/fa";
import { useAuth } from "@/context/AuthContext";
import { Button } from "../common/Button";

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
    <div className="z-10 fixed left-0 right-0 top-0 bg-zinc-950/60 backdrop-blur-lg border-b border-dashed border-b-zinc-800 py-2 px-8 flex justify-between items-center">
      <Link href="/" >
        <Image src="/logo.png" alt="logo" width={71} height={48} className="py-1.5 h-12 w-auto object-contain" />
      </Link>
      <div className="flex flex-row gap-4">
        <Button
          isPrimary={false}
          customClass="flex items-center justify-center gap-2 py-2.5"
          onClick={() => router.push(ROUTER.Tutorials)}
        >
          <FaBook /> Tutorials
        </Button>
        <Button
          isPrimary={false}
          customClass="flex items-center justify-center gap-2 py-2.5"
          onClick={() => router.push(ROUTER.Discussion)}
        >
          <FaComments /> Q&A
        </Button>
        <Button
          isPrimary={false}
          customClass="flex items-center justify-center gap-2 py-2.5"
          onClick={() => router.push(ROUTER.LeaderBoard)}
        >
          <FaTrophy /> LeaderBoard
        </Button>
        {user ?
          <div className="relative">
            <div
              className="flex items-center gap-2 font-semibold cursor-pointer text-[14px]"
              onClick={() => setIsShowUserDropdown(!isShowUserDropdown)}
            >
              {user.profileImage ?
                <Image
                  src={user.profileImage}
                  alt="avatar"
                  width={32}
                  height={32}
                  className="rounded-full border-solid border-zinc-50"
                />
                : <div className="w-8 h-8 border border-solid border-zinc-500 rounded-full p-1">
                  <FaUserTie size={18} />
                </div>
              }
              {user.fullName}
            </div>
            {isShowUserDropdown &&
              <div
                ref={dropdownRef}
                className="!z-10 flex flex-col gap-2 absolute translate-y-2 top-full right-0 p-2 border border-solid border-zinc-900 rounded-lg bg-zinc-900"
              >
                <a
                  href={ROUTER.Profile}
                  className="w-40 flex items-center gap-2 px-4 py-2 font-semibold cursor-pointer rounded-lg hover:bg-zinc-800 text-nowrap"
                >
                  <FaUserCircle size={18} />Profile
                </a>
                <div
                  className="w-40 flex items-center gap-2 px-4 py-2 font-semibold cursor-pointer rounded-lg hover:bg-zinc-800 text-nowrap"
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
      </div>
    </div>
  )
}