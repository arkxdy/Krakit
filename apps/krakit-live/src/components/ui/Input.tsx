import { type InputHTMLAttributes, forwardRef } from "react";

interface InputProps extends InputHTMLAttributes<HTMLInputElement> {
  label?: string;
  error?: string;
  fullWidth?: boolean;
}

export const Input = forwardRef<HTMLInputElement, InputProps>(
  ({ label, error, fullWidth = true, className = "", ...props }, ref) => {
    const baseInputStyles =
      "appearance-none px-3 py-2 border rounded-md shadow-sm placeholder-purple-400 focus:outline-none focus:ring-2 focus:ring-purple-500 focus:border-purple-500 dark:bg-purple-700 dark:text-white transition-all duration-200";
    const errorStyles = error
      ? "border-red-500 focus:ring-red-500 focus:border-red-500"
      : "border-purple-300 dark:border-purple-600";
    const widthClass = fullWidth ? "w-full" : "";

    return (
      <div className="space-y-1">
        {label && (
          <label className="block text-sm font-medium text-purple-700 dark:text-purple-200">
            {label}
          </label>
        )}
        <input
          ref={ref}
          className={`${baseInputStyles} ${errorStyles} ${widthClass} ${className}`}
          {...props}
        />
        {error && (
          <p className="text-sm text-red-600 dark:text-red-400">{error}</p>
        )}
      </div>
    );
  },
);
