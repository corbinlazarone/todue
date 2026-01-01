"use client";

import { createContext, useContext, useState, ReactNode } from "react";

interface DashboardContextType {
  sidebarDisabled: boolean;
  setSidebarDisabled: (disabled: boolean) => void;
}

const DashboardContext = createContext<DashboardContextType | undefined>(
  undefined
);

interface DashboardProviderProps {
  children: ReactNode;
}

export function DashboardProvider({ children }: DashboardProviderProps) {
  const [sidebarDisabled, setSidebarDisabled] = useState(false);

  return (
    <DashboardContext.Provider value={{ sidebarDisabled, setSidebarDisabled }}>
      {children}
    </DashboardContext.Provider>
  );
}

export function useDashboard() {
  const context = useContext(DashboardContext);
  if (context === undefined) {
    throw new Error("useDashboard must be used within a DashboardProvider");
  }
  return context;
}