"use client";

import { GoogleLogin } from "@react-oauth/google";

export function GoogleLoginButton() {
  return (
    <div>
      <GoogleLogin
        onSuccess={(credentialResponse) => {
          const jwt = credentialResponse.credential;
          // TODO: send to backend to validate
        }}
        onError={() => {
          console.log("Login Failed");
        }}
      />
    </div>
  );
}
