'use client'

import { getUserInfo, isUserLoggedIn, Logout } from '@/hooks/Auth'
import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuLabel,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from '@radix-ui/react-dropdown-menu'
import { BellIcon, ChatBubbleIcon, GlobeIcon, PersonIcon } from '@radix-ui/react-icons'
import Link from 'next/link'
import React, { useEffect } from 'react'
import { useAuth } from './AuthContext'

export default function NavBar() {
  const isAuth = isUserLoggedIn()
  const token = (isAuth && getUserInfo()) || null
  const {update, isAuthenticated} = useAuth()

  function navLogout() {
    Logout()
    update()
  }

  useEffect(() => {}, [isAuthenticated])

  return (
    <nav className="flex items-center justify-between px-6 py-4 border-b shadow-sm">
      <div className="flex gap-4 items-center">
        <Link href="/" className="text-xl font-bold">
          SocialNet
        </Link>

        <Link href="/groups" className="hover:text-blue-500 flex items-center gap-1">
          <GlobeIcon /> Groups
        </Link>
        {isAuth && (
          <a href="/chat" className="hover:text-blue-500 flex items-center gap-1">
            <ChatBubbleIcon /> Chat
          </a>
        )}
        <Link href="/profile" className="hover:text-blue-500 flex items-center gap-1">
          <PersonIcon /> Profile
        </Link>
      </div>

      <div className="flex items-center gap-4">
        <DropdownMenu>
          <DropdownMenuTrigger asChild>
            <button className="relative p-2 hover:bg-gray-100 rounded-full">
              <BellIcon className="w-5 h-5" />
              <span className="absolute top-0 right-0 inline-flex items-center justify-center px-1 text-xs font-bold leading-none text-white bg-red-500 rounded-full">
                3
              </span>
            </button>
          </DropdownMenuTrigger>
          <DropdownMenuContent className="w-64 p-2 bg-white shadow-lg border rounded-md">
            <DropdownMenuLabel>Notifications</DropdownMenuLabel>
            <DropdownMenuSeparator />
            <DropdownMenuItem>You have a new message</DropdownMenuItem>
            <DropdownMenuItem>New group invite</DropdownMenuItem>
            <DropdownMenuItem>Someone liked your post</DropdownMenuItem>
          </DropdownMenuContent>
        </DropdownMenu>

        {isAuth ? (
          <>
            Hello, {token!.nickname}
            <button onClick={navLogout} className="text-blue-600 font-medium hover:underline">
              Logout
            </button>
          </>
        ) : (
          <>
            <Link href="/login" className="text-blue-600 font-medium hover:underline">
              Login
            </Link>
            <Link href="/register" className="text-blue-600 font-medium hover:underline">
              Register
            </Link>
          </>
        )}
      </div>
    </nav>
  )
}
