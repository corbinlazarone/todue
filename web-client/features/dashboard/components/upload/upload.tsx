"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";
import { Card } from "@/shared/ui/card";
import { Courses } from "../../types";
import { toast } from "sonner";
import { useRef, useState } from "react";
import { extractCourseData } from "../../api/upload/actions";
import { extractTextFromPDF } from "../../api/upload/helpers";
import { FileText } from "lucide-react";

import AssignmentCard from "./assingment-card";

export function Upload() {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [extractedCourseData, setExtractedCourseData] = useState<Courses>();
  const [disWhileExtract, setDisWhileExtract] = useState<boolean>(false);
  const [disBeforeExtract, setDisBeforeExtract] = useState<boolean>(true);

  async function handleExtract() {
    const file = fileInputRef.current?.files?.[0];

    if (!file) {
      toast.error("Please select a file first");
      return;
    }

    const maxSize = 5 * 1024 * 1024; // 5MB
    if (file.size > maxSize) {
      toast.error("File size must be less than 5MB");
      return;
    }

    try {
      setDisWhileExtract(true);

      const buffer = Buffer.from(await file.arrayBuffer());
      const text = await extractTextFromPDF(buffer);

      const courseDataPromise = extractCourseData(text);

      toast.promise(courseDataPromise, {
        loading: "Uploading...",
        success: () => `${file.name} has been uploaded!`,
        error: "Error",
      });

      const courseData = await courseDataPromise;

      console.log(JSON.stringify(courseData, null, 2));

      setExtractedCourseData(courseData);
    } catch {
      toast.error("Failed to extract text from PDF. Try again or contact us.");
      return;
    } finally {
      setDisBeforeExtract(false);
      setDisWhileExtract(false);
    }
  }

  return (
    <div className="space-y-8 w-full max-w-7xl mx-auto px-4 py-6">
      <div className="flex gap-2">
        <Input
          ref={fileInputRef}
          id="syllabus"
          type="file"
          accept=".pdf"
          className="cursor-pointer w-auto"
        />
        <Button
          disabled={disWhileExtract || disWhileExtract}
          onClick={() => handleExtract()}
        >
          Extract
        </Button>
        <Button
          variant="secondary"
          disabled={disBeforeExtract || disWhileExtract}
          onClick={() => toast.info("Not implemented")}
        >
          Sync to Google Calendar
        </Button>
      </div>

      {!disBeforeExtract && (
        <Card className="p-3 text-sm border-l-4 border-l-blue-500 bg-blue-50 dark:bg-blue-950 text-blue-800 dark:text-blue-200">
          Please review the extracted assignments below for any mistakes before
          syncing to your primary Google Calendar.
        </Card>
      )}

      <div className="space-y-4">
        {extractedCourseData &&
        extractedCourseData.courses &&
        extractedCourseData.courses.length > 0 ? (
          <>
            <div className="flex items-center justify-between">
              <div>
                <h2 className="text-2xl font-semibold tracking-tight">
                  {extractedCourseData.courses[0].course_name}
                </h2>
              </div>
              <Button
                variant="outline"
                disabled={disBeforeExtract || disWhileExtract}
                onClick={() => toast.info("Not implemented")}
              >
                Add new Assignment
              </Button>
            </div>
            <AssignmentCard
              assignments={extractedCourseData.courses[0].assignments}
            />
          </>
        ) : extractedCourseData ? (
          <div className="text-red-500 p-4 bg-red-50 dark:bg-red-950 border border-red-200 dark:border-red-800 rounded">
            Error: Failed to extract courses from the PDF. Please try again or
            contact support if the issue persists.
          </div>
        ) : (
          <Card className="p-8 text-center border-dashed border-2 border-gray-300 dark:border-gray-600">
            <div className="flex flex-col items-center space-y-4">
              <FileText className="w-12 h-12 text-gray-400" />
              <div>
                <h3 className="text-lg font-medium text-gray-900 dark:text-gray-100">
                  Ready to Extract Assignments
                </h3>
                <p className="text-sm text-gray-500 dark:text-gray-400 mt-1">
                  Select a PDF syllabus and click Extract to see your
                  assignments here
                </p>
              </div>
            </div>
          </Card>
        )}
      </div>
    </div>
  );
}
