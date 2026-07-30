"use client";

import { useContext } from "react";
import Link from "next/link";
import { SidebarContext } from "@/components/SidebarProvider";

const navigation = {
  main: [
    { name: "About", href: "/about" },
    { name: "Pricing", href: "/pricing" },
    { name: "FAQ", href: "/faq" },
    { name: "Contact", href: "/contact" },
    { name: "Privacy Policy", href: "/privacy" },
    { name: "Terms of Service", href: "/terms" },
  ],
  social: [
    {
      name: "Twitter",
      href: "#",
      icon: "𝕏",
    },
    {
      name: "GitHub",
      href: "#",
      icon: "🐙",
    },
    {
      name: "LinkedIn",
      href: "#",
      icon: "💼",
    },
  ],
};

export function Footer() {
  const { isOpen } = useContext(SidebarContext);

  return (
    <footer
      className={`relative bg-gray-900 border-t border-gray-800 transition-all duration-300 ease-in-out ${isOpen ? "ml-64" : "ml-0"}`}
    >
      <div className="w-full mx-auto py-12 px-4 sm:px-6 lg:px-8">
        <nav className="flex flex-wrap justify-center -mx-5 -my-2">
          {navigation.main.map((item) => (
            <div key={item.name} className="px-5 py-2">
              <Link
                href={item.href}
                className="text-gray-400 hover:text-white transition-colors"
              >
                {item.name}
              </Link>
            </div>
          ))}
        </nav>
        <div className="mt-8 flex justify-center space-x-6">
          {navigation.social.map((item) => (
            <a
              key={item.name}
              href={item.href}
              className="text-gray-400 hover:text-white transition-colors"
            >
              <span className="sr-only">{item.name}</span>
              <span className="text-xl">{item.icon}</span>
            </a>
          ))}
        </div>
        <p className="mt-8 text-center text-gray-400">
          © {new Date().getFullYear()} Krakit. All rights reserved.
        </p>
      </div>
    </footer>
  );
}
