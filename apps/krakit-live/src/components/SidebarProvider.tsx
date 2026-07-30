"use client";

import { createContext, useMemo, useState, type ReactNode } from "react";

interface SidebarContextValue {
  isOpen: boolean;
  setIsOpen: (isOpen: boolean) => void;
}

export const SidebarContext = createContext<SidebarContextValue>({
  isOpen: false,
  setIsOpen: () => {},
});

export function SidebarProvider({ children }: { children: ReactNode }) {
  const [isOpen, setIsOpen] = useState(true);

  const value = useMemo(
    () => ({ isOpen, setIsOpen }),
    [isOpen],
  );

  return (
    <SidebarContext.Provider value={value}>{children}</SidebarContext.Provider>
  );
}
