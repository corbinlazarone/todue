"use server";

import { getSession } from "@/utils/session";
import PDFParser from "pdf2json";

function safeDecodeURIComponent(str: string): string {
  try {
    return decodeURIComponent(str);
  } catch (error) {
    console.warn("Failed to decode URI component:", str);
    return str;
  }
}

export async function extractTextFromPDF(buffer: any): Promise<string> {
  const actualBuffer = Buffer.from(buffer.data || buffer);
  console.log("Buffer length:", actualBuffer.length);

  return new Promise((resolve, reject) => {
    const pdfParser = new PDFParser(null, true);

    pdfParser.on("pdfParser_dataReady", (pdfData) => {
      try {
        const text = pdfData.Pages.map((page) =>
          page.Texts.map((text) => safeDecodeURIComponent(text.R[0].T)).join(
            " ",
          ),
        ).join("\n");
        resolve(text);
      } catch (error) {
        console.error("Failed to parse PDF content:", error);
        reject(new Error("Failed to parse PDF content"));
      }
    });

    pdfParser.on("pdfParser_dataError", (error) => {
      console.error("PDF parsing error:", error);
      reject(new Error(`PDF parsing error: ${error}`));
    });

    try {
      pdfParser.parseBuffer(actualBuffer);
    } catch (error) {
      console.error("Failed to parse PDF buffer:", error);
      reject(new Error("Failed to parse PDF buffer"));
    }
  });
}

export async function extractCourseData(test: string) {
  const apiUrl = process.env.NEXT_PUBLIC_API_URL;

  const session = await getSession();

  const rep = await fetch(`${apiUrl}/api/ai/extract`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${session.token}`,
    },
    body: JSON.stringify({
      pdfText: test,
    }),
  });

  if (!rep.ok) {
    const errorText = await rep.text();
    console.error("Backend error:", rep.status, errorText);
    throw new Error(`Failed to extract course data: ${rep.status}`);
  }

  const data = await rep.json(); // expeted type of Courses: []CourseData

  return data;
}
