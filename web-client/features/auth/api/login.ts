"use server";

export async function login(googleJWT: string | undefined) {
  if (!googleJWT) {
    throw new Error("Google JWT undefined");
  }

  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  if (!apiUrl) {
    throw new Error("API URL not set");
  }

  const rep = await fetch(`${apiUrl}/auth/login`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify({
      googleJWT,
    }),
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    console.error("Backend error:", rep.status, errorText);
    throw new Error(`Failed to login: ${rep.status}`);
  }

  const data = await rep.json();
  console.log(JSON.stringify(data));
}
