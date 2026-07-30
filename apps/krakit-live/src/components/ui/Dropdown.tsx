import { useState, useRef, useEffect, type ReactNode } from "react";

interface DropdownOption {
  value: string;
  label: string;
  disabled?: boolean;
}

"use client";

interface DropdownProps {
  options: DropdownOption[];
  value?: string;
  onChange: (value: string) => void;
  placeholder?: string;
  disabled?: boolean;
  className?: string;
  trigger?: ReactNode;
  size?: "sm" | "md" | "lg";
}

export function Dropdown({
  options,
  value,
  onChange,
  placeholder = "Select an option",
  disabled = false,
  className = "",
  trigger,
  size = "md",
}: DropdownProps) {
  const [isOpen, setIsOpen] = useState(false);
  const dropdownRef = useRef<HTMLDivElement>(null);

  const selectedOption = options.find((option) => option.value === value);

  const sizeClasses = {
    sm: "px-3 py-2 text-sm",
    md: "px-4 py-2 text-base",
    lg: "px-4 py-3 text-lg",
  };

  useEffect(() => {
    const handleClickOutside = (event: MouseEvent) => {
      if (
        dropdownRef.current &&
        !dropdownRef.current.contains(event.target as Node)
      ) {
        setIsOpen(false);
      }
    };

    document.addEventListener("mousedown", handleClickOutside);
    return () => document.removeEventListener("mousedown", handleClickOutside);
  }, []);

  const handleOptionClick = (optionValue: string) => {
    onChange(optionValue);
    setIsOpen(false);
  };

  return (
    <div className={`relative ${className}`} ref={dropdownRef}>
      {/* Trigger Button */}
      <button
        type="button"
        onClick={() => !disabled && setIsOpen(!isOpen)}
        disabled={disabled}
        className={`
          w-full flex items-center justify-between
          bg-gray-800/50 backdrop-blur-sm border border-gray-700/50 rounded-lg
          text-white placeholder-gray-400
          transition-all duration-200
          ${sizeClasses[size]}
          ${
            disabled
              ? "opacity-50 cursor-not-allowed"
              : "hover:bg-gray-700/50 hover:border-gray-600/50 focus:outline-none focus:ring-2 focus:ring-primary-500/50 focus:border-primary-500/50"
          }
        `}
      >
        {trigger || (
          <>
            <span className={selectedOption ? "text-white" : "text-gray-400"}>
              {selectedOption ? selectedOption.label : placeholder}
            </span>
            <svg
              className={`w-5 h-5 text-gray-400 transition-transform duration-200 ${
                isOpen ? "rotate-180" : ""
              }`}
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth={2}
                d="M19 9l-7 7-7-7"
              />
            </svg>
          </>
        )}
      </button>

      {/* Dropdown Menu */}
      {isOpen && (
        <div className="absolute z-50 w-full mt-1 bg-gray-800/95 backdrop-blur-sm border border-gray-700/50 rounded-lg shadow-xl">
          <div className="py-1 max-h-60 overflow-auto">
            {options.map((option) => (
              <button
                key={option.value}
                type="button"
                onClick={() =>
                  !option.disabled && handleOptionClick(option.value)
                }
                disabled={option.disabled}
                className={`
                  w-full text-left px-4 py-2 text-sm
                  transition-colors duration-150
                  ${
                    option.disabled
                      ? "text-gray-500 cursor-not-allowed"
                      : option.value === value
                        ? "bg-primary-600/20 text-primary-400 hover:bg-primary-600/30"
                        : "text-gray-300 hover:bg-gray-700/50 hover:text-white"
                  }
                `}
              >
                {option.label}
              </button>
            ))}
          </div>
        </div>
      )}
    </div>
  );
}
