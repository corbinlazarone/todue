"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";
import { Card } from "@/shared/ui/card";
import { toast } from "sonner";
import { useRef, useState } from "react";
import { extractCourseData } from "../../api/upload/actions";
import { extractTextFromPDF } from "../../api/upload/helpers";

import AssignmentCard from "./assingment-card";

export function Upload() {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [disWhileExtract, setDisWhileExtract] = useState(false);
  const [disBeforeExtract, setDisBeforeExtract] = useState(true);

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

      // NOTE: console log
      console.log(JSON.stringify(courseData, null, 2));
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
          {/* TODO: Tell the user that we will use their primary google calendar.
             show them the name of that calendar before sumbiting.*/}
          Please review the extracted assignments below for any mistakes before
          syncing to your Google Calendar.
        </Card>
      )}

      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-2xl font-semibold tracking-tight">
              Your Extracted Assignments
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
        <AssignmentCard />
      </div>
    </div>
  );
}
