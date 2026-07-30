import { useState } from "react";

interface AccordionProps {
  title: string;
  children: React.ReactNode;
  isActive?: boolean;
  defaultOpen?: boolean;
  isMobile?: boolean;
}

"use client";

function Accordion({
  title,
  children,
  isActive = false,
  defaultOpen = false,
  isMobile = false,
}: AccordionProps) {
  const [isOpen, setIsOpen] = useState(defaultOpen);

  return (
    <div className="border rounded-md overflow-hidden">
      <button
        className={`w-full p-2 sm:p-3 text-left font-medium transition-colors flex justify-between items-center text-sm sm:text-base
          ${isActive ? "bg-blue-50 border-l-4 border-blue-500" : "bg-gray-50"}
          hover:bg-gray-100`}
        onClick={() => setIsOpen(!isOpen)}
      >
        <span>{title}</span>
        <span className="text-gray-500">{isOpen ? "−" : "+"}</span>
      </button>
      {isOpen && <div className="p-1 sm:p-2">{children}</div>}
    </div>
  );
}

export default Accordion;
