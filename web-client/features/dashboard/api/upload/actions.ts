"use server";

import { getSession } from "@/utils/session";
import { Courses, EventError } from "../../types";

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

export async function insertToGoogleCalendar(
  userTimezone: string,
  courses: Courses,
): Promise<void | EventError[]> {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const session = await getSession();

  const rep = await fetch(`${apiUrl}/api/event/save`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.token}`,
    },
    body: JSON.stringify({
      timezone: userTimezone,
      courses: courses.courses,
    }),
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    let parsedError: SyntaxError;

    try {
      parsedError = JSON.parse(errorText);
    } catch {
      throw new Error(`Server error (${rep.status}): ${errorText}`);
    }

    if (Array.isArray(parsedError)) {
      const eventErrors: EventError[] = parsedError.filter(
        (item) =>
          item.assignment_id !== undefined &&
          item.assignment_name !== undefined &&
          Array.isArray(item.errors),
      ) as EventError[];

      if (eventErrors.length > 0) {
        return eventErrors;
      }

      // No valid errors found
      throw new Error(`Server error (${rep.status}): ${errorText}`);
    } else {
      throw new Error(`Server error (${rep.status}): ${errorText}`);
    }
  }

  return;
}
