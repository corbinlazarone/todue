"use client";

import { GoogleLogin } from "@react-oauth/google";
import { login } from "../api/login";

export function GoogleLoginButton() {
  return (
    <div>
      <GoogleLogin
        onSuccess={(credentialResponse) => {
          const jwt = credentialResponse.credential;
          login(jwt);
        }}
        onError={() => {
          console.log("Login Failed");
        }}
      />
    </div>
  );
}
