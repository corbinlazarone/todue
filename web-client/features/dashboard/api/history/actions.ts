"use server";

import { getSession } from "@/utils/session";
import { CourseData } from "../../types";

export async function fetchCourseHistory(): Promise<CourseData[]> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const session = await getSession();

  const rep = await fetch(`${apiUrl}/api/history/all`, {
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.token}`,
    },
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    console.error("Backend error:", rep.status, errorText);
    throw new Error();
  }

  const data: CourseData[] = await rep.json();

  return data;
}

