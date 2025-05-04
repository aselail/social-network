import React from "react";
import {BellIcon, ChatBubbleIcon, GlobeIcon, PersonIcon} from "@radix-ui/react-icons";
import {Theme} from "@radix-ui/themes";
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger
} from "@radix-ui/react-dropdown-menu";
import Link from "next/link";

export default function HomePage() {
    return (
        <Theme appearance="light">
            <div className="min-h-screen bg-white text-gray-800">
                <nav className="flex items-center justify-between px-6 py-4 border-b shadow-sm">
                    <div className="flex gap-4 items-center">
                        <span className="text-xl font-bold">SocialNet</span>
                        <Link href="/groups" className="hover:text-blue-500 flex items-center gap-1">
                            <GlobeIcon/> Groups
                        </Link>
                        <a href="/chat" className="hover:text-blue-500 flex items-center gap-1">
                            <ChatBubbleIcon/> Chat
                        </a>
                        <Link href="/profile" className="hover:text-blue-500 flex items-center gap-1">
                            <PersonIcon/> Profile
                        </Link>
                    </div>

                    <div className="flex items-center gap-4">
                        <DropdownMenu>
                            <DropdownMenuTrigger asChild>
                                <button className="relative p-2 hover:bg-gray-100 rounded-full">
                                    <BellIcon className="w-5 h-5"/>
                                    <span
                                        className="absolute top-0 right-0 inline-flex items-center justify-center px-1 text-xs font-bold leading-none text-white bg-red-500 rounded-full">
                    3
                  </span>
                                </button>
                            </DropdownMenuTrigger>
                            <DropdownMenuContent className="w-64 p-2 bg-white shadow-lg border rounded-md">
                                <DropdownMenuLabel>Notifications</DropdownMenuLabel>
                                <DropdownMenuSeparator/>
                                <DropdownMenuItem>You have a new message</DropdownMenuItem>
                                <DropdownMenuItem>New group invite</DropdownMenuItem>
                                <DropdownMenuItem>Someone liked your post</DropdownMenuItem>
                            </DropdownMenuContent>
                        </DropdownMenu>

                        <Link href="/login" className="text-blue-600 font-medium hover:underline">Login</Link>
                        <Link href="/register" className="text-blue-600 font-medium hover:underline">Register</Link>
                    </div>
                </nav>

                <main className="p-10 text-center">
                    <h1 className="text-4xl font-bold mb-4">Welcome to SocialNet</h1>
                    <p className="text-gray-600 text-lg">Connect with friends, join groups, and chat in real-time.</p>
                </main>
            </div>
        </Theme>
    );
}
