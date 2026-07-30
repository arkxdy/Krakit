"use client";

import { useContext, useState } from "react";
import Link from "next/link";
import { ProUpgrade } from "../ProUpgrade";
import { Notifications } from "../Notifications";
import { SidebarContext } from "@/components/SidebarProvider";
import type { User } from "@/types/user.types";

export function PrivateHeader({ user }: { user: User }) {
  const { isOpen } = useContext(SidebarContext);
  const [isNotificationsOpen, setIsNotificationsOpen] = useState(false);

  return (
    <>
      <header
        className={`fixed top-0 right-0 z-50 bg-gray-900/80 backdrop-blur-sm border-b border-gray-800 transition-all duration-300 ease-in-out ${isOpen ? "left-64" : "left-0 pl-8"}`}
        style={{ minHeight: "64px" }}
      >
        <div className="w-full px-4 sm:px-6 lg:px-8">
          <div className="flex justify-between items-center h-16">
            {/* Welcome Message */}
            <div className={`flex items-center`}>
              <div className="hidden md:block">
                <h2 className="text-gray-300">
                  Welcome back,{" "}
                  <span className="text-white font-semibold">
                    {user?.name || "User"}
                  </span>
                </h2>
              </div>
            </div>

            {/* Right Side Icons */}
            <div className="flex items-center space-x-6">
              {/* Notifications */}
              <button
                onClick={() => setIsNotificationsOpen(true)}
                className="rounded-full relative text-gray-300 hover:text-white transition-colors"
              >
                <svg
                  className="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M15 17h5l-1.405-1.405A2.032 2.032 0 0118 14.158V11a6.002 6.002 0 00-4-5.659V5a2 2 0 10-4 0v.341C7.67 6.165 6 8.388 6 11v3.159c0 .538-.214 1.055-.595 1.436L4 17h5m6 0v1a3 3 0 11-6 0v-1m6 0H9"
                  />
                </svg>
                {/* Notification Badge */}
                <span className="absolute -top-1 -right-1 bg-red-500 text-white text-xs rounded-full h-4 w-4 flex items-center justify-center">
                  3
                </span>
              </button>

              {/* User Profile */}
              <Link
                href="/profile"
                className="flex items-center space-x-2 text-gray-300 hover:text-white transition-colors"
              >
                <div className="relative">
                  <div className="w-8 h-8 rounded-full bg-gray-700 flex items-center justify-center">
                    <span className="text-white font-medium">
                      {user?.name?.charAt(0) || "U"}
                    </span>
                  </div>
                  {/* Online Status */}
                  <span className="absolute bottom-0 right-0 w-3 h-3 bg-green-500 rounded-full border-2 border-gray-900"></span>
                </div>
                <span className="hidden md:block text-sm font-medium">
                  {user?.name || "User"}
                </span>
              </Link>
            </div>
          </div>
        </div>
      </header>

      <Notifications
        isOpen={isNotificationsOpen}
        onClose={() => setIsNotificationsOpen(false)}
      />
    </>
  );
}
