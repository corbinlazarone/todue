"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";

import AssignmentCard from "./assingment-card";
import { toast } from "sonner";
import { useRef } from "react";
import {
  extractCourseData,
  extractTextFromPDF,
} from "../../api/upload/actions";

export function Upload() {
  const fileInputRef = useRef<HTMLInputElement>(null);

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
    } catch {
      toast.error("Failed to extract text from PDF. Try again or contact us.");
      return;
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
        <Button onClick={() => handleExtract()}>Extract</Button>
        <Button
          variant="outline"
          onClick={() => toast.info("Not implemented")}
        >
          Sync to Google Calendar
        </Button>
      </div>

      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-2xl font-semibold tracking-tight">
              Your Extracted Assignments
            </h2>
          </div>
          <Button
            variant="outline"
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
