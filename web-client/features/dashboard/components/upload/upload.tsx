"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";

import AssignmentCard from "./assingment-card";
import { toast } from "sonner";
import { useRef } from "react";
import { extractTextFromPDF } from "../../api/upload/actions";

export function Upload() {
  const fileInputRef = useRef<HTMLInputElement>(null);

  async function handleExtract() {
    const file = fileInputRef.current?.files?.[0];

    if (!file) {
      toast.error("Please select a file first");
      return;
    }

    try {
      const buffer = Buffer.from(await file.arrayBuffer());
      console.log("Buffer:", buffer);

      await extractTextFromPDF(buffer);

      // TODO: send to go api to claude

      toast.success("Extracted text from PDF");
    } catch (error) {
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
            onClick={() => toast.info("Not implmented")}
          >
            Sync to Google Calendar
          </Button>
        </div>
        <AssignmentCard />
      </div>
    </div>
  );
}
