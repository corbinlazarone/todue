"use server";

import { getSession } from "@/utils/session";
import { AppJWTPayload, defaultResponse } from "../types";
import { revalidatePath } from "next/cache";
import { jwtDecode } from "jwt-decode";

export async function login(authCode: string | undefined) {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  if (!authCode) {
    console.error("AuthCode is undefined");
    throw new Error();
  }

  const rep = await fetch(`${apiUrl}/api/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      authCode,
    }),
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    console.error("Backend error:", rep.status, errorText);
    throw new Error();
  }

  const data: defaultResponse = await rep.json();
  await saveUserSession(data.message);

  revalidatePath("/");
  revalidatePath("/dashboard");
}

export async function logout() {
  const session = await getSession();
  session.destroy();

  // Revalidate pages to show logged-out state
  revalidatePath("/");
  revalidatePath("/dashboard");

  return { success: true };
}

// decode our jwt so we can get user data from it.
async function saveUserSession(token: string) {
  try {
    const decodedToken = jwtDecode<AppJWTPayload>(token);

    const session = await getSession();
    session.token = token;
    session.isLoggedIn = true;
    session.userData = {
      email: decodedToken.email,
      firstName: decodedToken.first_name,
      lastName: decodedToken.last_name,
      picture: decodedToken.picture,
    };

    await session.save();
  } catch (err) {
    console.error("Failed to decode token: ", err);
    throw new Error();
  }
}
