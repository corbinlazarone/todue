"use client";

import { Button } from "@/shared/ui/button";
import { useRouter } from "next/navigation";

export function Landing({ isLoggedIn }: { isLoggedIn: boolean }) {
  const router = useRouter();
  return (
    <div className="flex flex-col items-center justify-center h-screen">
      <h1>Todue Landing Page</h1>
      {isLoggedIn ? (
        <Button onClick={() => router.push("/dashboard")}>Dashboard</Button>
      ) : (
        <Button onClick={() => router.push("/login")}>Sign in</Button>
      )}
    </div>
  );
}
