"use client";

import { logout } from "../api/actions";

export function LogoutButton() {
  return <button onClick={async () => await logout()}>Logout</button>;
}
