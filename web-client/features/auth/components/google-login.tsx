"use client";

import { GoogleLogin } from "@react-oauth/google";
import { login, logout } from "../api/actions";
import { useRouter } from "next/navigation";

export function GoogleLoginButton() {
  const router = useRouter();

  return (
    <div>
      <GoogleLogin
        onSuccess={async (credentialResponse) => {
          const jwt = credentialResponse.credential;
          try {
            await login(jwt);
            router.push("/dashboard");
          } catch (error) {
            console.error("Login error:", error);
          }
        }}
        onError={() => {
          console.log("Login Failed");
        }}
      />
    </div>
  );
}

export function LogoutButton() {
  return <button onClick={async () => await logout()}>Logout</button>;
}
