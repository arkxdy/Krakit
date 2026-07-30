import type { HTMLAttributes, ReactNode } from "react";

interface CardProps extends HTMLAttributes<HTMLDivElement> {
  children: ReactNode;
  hover?: boolean;
}

export function Card({
  children,
  hover = true,
  className = "",
  ...props
}: CardProps) {
  const baseStyles =
    "bg-white dark:bg-purple-800 py-8 px-4 shadow-xl sm:rounded-lg sm:px-10";
  const hoverStyles = hover
    ? "transform transition-all duration-300 hover:shadow-2xl"
    : "";

  return (
    <div className={`${baseStyles} ${hoverStyles} ${className}`} {...props}>
      {children}
    </div>
  );
}
