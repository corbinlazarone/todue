"use client";

import { useTheme } from "next-themes";

export function Landing() {
  const { setTheme } = useTheme();

  return (
    <div>
      <button onClick={() => setTheme("dark")}>switch to dark</button>
      <button onClick={() => setTheme("light")}>switch to light</button>
    </div>
  );
}
