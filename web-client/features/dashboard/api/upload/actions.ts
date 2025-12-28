"use server";

import { getSession } from "@/utils/session";
import { Courses } from "../../types";

export async function extractCourseData(text: string): Promise<Courses> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const session = await getSession();

  const rep = await fetch(`${apiUrl}/api/ai/extract`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.token}`,
    },
    body: JSON.stringify({
      pdfText: text,
    }),
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    console.error("Backend error:", rep.status, errorText);
    throw new Error();
  }

  const data: Courses = await rep.json();

  return data;
}
