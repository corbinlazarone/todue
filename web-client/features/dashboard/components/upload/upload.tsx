"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";

import AssignmentCard from "./assingment-card";
import { toast } from "sonner";

export function Upload() {
  return (
    <div className="space-y-8 w-full max-w-7xl mx-auto px-4 py-6">
      <div className="flex gap-2">
        <Input
          id="syllabus"
          type="file"
          accept=".pdf"
          className="cursor-pointer w-auto"
        />
        <Button onClick={() => toast.info("Not implmented")}>Extract</Button>
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
