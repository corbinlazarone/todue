"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";
import { Card } from "@/shared/ui/card";
import { Assignment, Courses, EventError } from "../../types";
import { toast } from "sonner";
import { useEffect, useRef, useState } from "react";
import {
  extractCourseData,
  insertToGoogleCalendar,
} from "../../api/upload/actions";
import { extractTextFromPDF } from "../../api/upload/helpers";
import { useDashboard } from "../../context";
import { FileText } from "lucide-react";

import AssignmentCard from "./assingment-card";

export function Upload() {
  const fileInputRef = useRef<HTMLInputElement>(null);
  const [userTimezone, setTimezone] = useState<string>("America/New_York");
  const [extractedCourseData, setExtractedCourseData] = useState<Courses>();
  const [disWhileExtract, setDisWhileExtract] = useState<boolean>(false);
  const [disBeforeExtract, setDisBeforeExtract] = useState<boolean>(true);
  const [syncErrors, setSyncErrors] = useState<EventError[] | null>(null);
  const { setSidebarDisabled } = useDashboard();

  useEffect(() => {
    const detectedTimezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    setTimezone(detectedTimezone);
  }, []);

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
      setSidebarDisabled(true);

      const buffer = Buffer.from(await file.arrayBuffer());
      const text = await extractTextFromPDF(buffer);

      const courseDataPromise = extractCourseData(text);

      toast.promise(courseDataPromise, {
        loading: "Uploading...",
        success: () => `${file.name} has been uploaded!`,
      });

      const courseData = await courseDataPromise;

      setExtractedCourseData(courseData);

      setDisBeforeExtract(false);
    } catch {
      toast.error("Failed to extract text from PDF. Try again or contact us.");
      return;
    } finally {
      setDisWhileExtract(false);
      setSidebarDisabled(false);
    }
  }

  async function handleGoogleCalSync() {
    setSyncErrors(null);
    try {
      if (!extractedCourseData || !userTimezone) {
        toast.error("An error occured. Try again or contact us.");
        return;
      }

      const result = await insertToGoogleCalendar(
        userTimezone,
        extractedCourseData,
      );

      if (Array.isArray(result)) {
        setSyncErrors(result);
      } else {
        toast.success("Assignments have been synced!");
      }
    } catch {
      toast.error(
        "Failed to upload to Google Calendar. Try again or contact us.",
      );
    }
  }

  function handleAssignmentDelete(id: number) {
    setExtractedCourseData((prev) => {
      if (!prev || !prev.courses || prev.courses.length === 0) return prev;

      return {
        ...prev,
        courses: prev.courses.map((course) => ({
          ...course,
          assignments: course.assignments.filter(
            (assignment) => assignment.id !== id,
          ),
        })),
      };
    });
  }

  function handleAssignmentEdit(data: Assignment) {
    setExtractedCourseData((prev) => {
      if (!prev || !prev.courses) return prev;

      return {
        ...prev,
        courses: prev.courses.map((course) => ({
          ...course,
          assignments: course.assignments.map((assignment) =>
            assignment.id === data.id ? data : assignment,
          ),
        })),
      };
    });
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
          onClick={() => handleGoogleCalSync()}
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
              timezone={userTimezone}
              onDelete={handleAssignmentDelete}
              onEdit={handleAssignmentEdit}
              syncErrors={syncErrors}
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
