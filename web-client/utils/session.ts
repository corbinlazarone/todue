import { getIronSession, SessionOptions } from "iron-session";
import { cookies } from "next/headers";

export interface SessionData {
  token: string;
  isLoggedIn: boolean;
  userData: {
    email: string;
    fistName: string | null;
    lastName: string | null;
    picture: string | null;
  };
}

export const defaultSession: SessionData = {
  token: "",
  isLoggedIn: false,
  userData: {
    fistName: "",
    lastName: "",
    email: "",
    picture: "",
  },
};

if (!process.env.SESSION_PASSWORD) {
  throw new Error("SESSION_PASSWORD not set");
}

export const sessionOptions: SessionOptions = {
  password: process.env.SESSION_PASSWORD,
  cookieName: "todue-session",
  cookieOptions: {
    secure: process.env.NODE_ENV === "production",
    sameSite: "lax",
    maxAge: 60 * 60, // 1 hour
  },
};

export async function getSession() {
  const session = await getIronSession<SessionData>(
    await cookies(),
    sessionOptions,
  );

  if (!session.isLoggedIn) {
    session.token = "";
    session.isLoggedIn = false;
    session.userData = {
      fistName: "",
      lastName: "",
      email: "",
      picture: "",
    };
  }

  return session;
}
