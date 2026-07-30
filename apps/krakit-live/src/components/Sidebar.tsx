"use client";

import { useContext } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/Button";
import { ProUpgrade } from "./ProUpgrade";
import { krakitIcon } from "@/utils/constant";
import { SidebarContext } from "@/components/SidebarProvider";

export function Sidebar() {
  const { isOpen, setIsOpen } = useContext(SidebarContext);
  const router = useRouter();

  const handleLogout = async () => {
    try {
      router.push("/logout");
    } catch (error) {
      console.error("Logout failed:", error);
    }
  };

  return (
    <>
      <button
        onClick={() => setIsOpen(!isOpen)}
        className="fixed left-4 top-4 z-50 p-2 rounded-lg bg-gray-800/50 backdrop-blur-sm text-gray-300 hover:bg-gray-700/50 transition-colors"
      >
        {isOpen ? (
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
              d="M11 19l-7-7 7-7m8 14l-7-7 7-7"
            />
          </svg>
        ) : (
          <img
            src={krakitIcon}
            alt="Krakit Logo"
            className="w-6 h-6"
            width="24"
            height="24"
          />
        )}
      </button>

      <div
        className={`sidebar-container bg-gray-800/50 backdrop-blur-sm h-screen fixed left-0 top-0 pt-16 z-40 transition-transform duration-300 ease-in-out flex flex-col ${
          isOpen ? "translate-x-0" : "-translate-x-full"
        }`}
      >
        {/* Logo and Brand Name */}
        <div className="px-4 py-3 border-b border-gray-700/50 flex-shrink-0">
          <Link href="/" className="flex items-center space-x-2">
            <img
              src={krakitIcon}
              alt="Krakit Logo"
              className="w-6 h-6"
              width="24"
              height="24"
            />
            <span className="text-xl font-bold text-white">Krakit</span>
          </Link>
        </div>

        <nav className="p-4 flex-1 overflow-y-auto">
          <ul className="space-y-2">
            <li>
              <Link
                href="/dashboard"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M3 12l2-2m0 0l7-7 7 7M5 10v10a1 1 0 001 1h3m10-11l2 2m-2-2v10a1 1 0 01-1 1h-3m-6 0a1 1 0 001-1v-4a1 1 0 011-1h2a1 1 0 011 1v4a1 1 0 001 1m-6 0h6"
                  />
                </svg>
                <span className="text-gray-300">Dashboard</span>
              </Link>
            </li>
            <li>
              <Link
                href="/exams"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 5H7a2 2 0 00-2 2v12a2 2 0 002 2h10a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2"
                  />
                </svg>
                <span className="text-gray-300">Exams</span>
              </Link>
            </li>
            <li>
              <Link
                href="/materials"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M12 6.253v13m0-13C10.832 5.477 9.246 5 7.5 5S4.168 5.477 3 6.253v13C4.168 18.477 5.754 18 7.5 18s3.332.477 4.5 1.253m0-13C13.168 5.477 14.754 5 16.5 5c1.747 0 3.332.477 4.5 1.253v13C19.832 18.477 18.247 18 16.5 18c-1.746 0-3.332.477-4.5 1.253"
                  />
                </svg>
                <span className="text-gray-300">Study Materials</span>
              </Link>
            </li>
            <li>
              <Link
                href="/results"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M9 19v-6a2 2 0 00-2-2H5a2 2 0 00-2 2v6a2 2 0 002 2h2a2 2 0 002-2zm0 0V9a2 2 0 012-2h2a2 2 0 012 2v10m-6 0a2 2 0 002 2h2a2 2 0 002-2m0 0V5a2 2 0 012-2h2a2 2 0 012 2v14a2 2 0 01-2 2h-2a2 2 0 01-2-2z"
                  />
                </svg>
                <span className="text-gray-300">Results</span>
              </Link>
            </li>
            <li>
              <Link
                href="/progress"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M13 7h8m0 0v8m0-8l-8 8-4-4-6 6"
                  />
                </svg>
                <span className="text-gray-300">Progress</span>
              </Link>
            </li>
            <li>
              <Link
                href="/profile"
                className="flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50 rounded-lg transition-colors"
              >
                <svg
                  className="w-5 h-5"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    strokeLinecap="round"
                    strokeLinejoin="round"
                    strokeWidth={2}
                    d="M16 7a4 4 0 11-8 0 4 4 0 018 0zM12 14a7 7 0 00-7 7h14a7 7 0 00-7-7z"
                  />
                </svg>
                <span className="text-gray-300">Profile</span>
              </Link>
            </li>
          </ul>
        </nav>

        {/* Pro Upgrade Section */}
        <div className="flex-shrink-0">
          <ProUpgrade />
        </div>

        {/* Logout Button */}
        <div className="p-4 border-t border-gray-700/50 flex-shrink-0">
          <Button
            onClick={handleLogout}
            variant="outline"
            className="w-full flex items-center space-x-3 px-4 py-3 text-gray-300 hover:bg-gray-700/50"
          >
            <svg
              className="w-5 h-5"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
              xmlns="http://www.w3.org/2000/svg"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1"
              />
            </svg>
            <span>Logout</span>
          </Button>
        </div>
      </div>
    </>
  );
}
