"use client";

import type { ButtonHTMLAttributes, ReactNode } from "react";

interface ButtonProps extends ButtonHTMLAttributes<HTMLButtonElement> {
  children: ReactNode;
  variant?: "primary" | "secondary" | "outline" | "white";
  isLoading?: boolean;
  fullWidth?: boolean;
}

export function Button({
  children,
  variant = "primary",
  isLoading = false,
  fullWidth = false,
  className = "",
  disabled,
  ...props
}: ButtonProps) {
  const baseStyles =
    "flex justify-center py-2 px-4 rounded-md shadow-sm text-sm font-medium transition-all duration-200";

  const variants = {
    primary:
      "text-white bg-purple-600 hover:bg-purple-700 focus:ring-2 focus:ring-offset-2 focus:ring-purple-500",
    secondary:
      "text-purple-700 bg-purple-100 hover:bg-purple-200 focus:ring-2 focus:ring-offset-2 focus:ring-purple-500",
    outline:
      "text-purple-700 bg-white border border-purple-300 hover:bg-purple-50 focus:ring-2 focus:ring-offset-2 focus:ring-purple-500",
    white:
      "text-primary-900 bg-white hover:bg-gray-100 focus:ring-2 focus:ring-offset-2 focus:ring-white",
  };

  const widthClass = fullWidth ? "w-full" : "";
  const loadingClass = isLoading
    ? "opacity-75 cursor-not-allowed"
    : "hover:scale-[1.02]";

  return (
    <button
      className={`${baseStyles} ${variants[variant]} ${widthClass} ${loadingClass} ${className}`}
      disabled={disabled || isLoading}
      {...props}
    >
      {isLoading ? (
        <span className="flex items-center">
          <svg
            className="animate-spin -ml-1 mr-3 h-5 w-5 text-current"
            xmlns="http://www.w3.org/2000/svg"
            fill="none"
            viewBox="0 0 24 24"
          >
            <circle
              className="opacity-25"
              cx="12"
              cy="12"
              r="10"
              stroke="currentColor"
              strokeWidth="4"
            ></circle>
            <path
              className="opacity-75"
              fill="currentColor"
              d="M4 12a8 8 0 018-8V0C5.373 0 0 5.373 0 12h4zm2 5.291A7.962 7.962 0 014 12H0c0 3.042 1.135 5.824 3 7.938l3-2.647z"
            ></path>
          </svg>
          Loading...
        </span>
      ) : (
        children
      )}
    </button>
  );
}
