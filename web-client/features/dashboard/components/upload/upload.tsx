"use client";

import { Button } from "@/shared/ui/button";
import { Input } from "@/shared/ui/input";
import { Card, CardContent } from "@/shared/ui/card";

import AssignmentCard from "./assingment-card";

export function Upload() {
  return (
    <div className="space-y-8 w-full max-w-7xl mx-auto px-4 py-6">
      {/* Upload and Sync Section */}
      <div className="grid grid-cols-1 lg:grid-cols-2 gap-4">
        {/* Upload Syllabus */}
        <Card>
          <CardContent className="pt-6">
            <div className="flex gap-2">
              <Input
                id="syllabus"
                type="file"
                accept=".pdf,.doc,.docx"
                className="cursor-pointer"
              />
              <Button>Extract</Button>
            </div>
          </CardContent>
        </Card>

        {/* Google Calendar Sync */}
        <Card>
          <CardContent className="pt-6">
            <Button variant="outline" className="w-full">
              Sync to Google Calendar
            </Button>
          </CardContent>
        </Card>
      </div>

      {/* Assignments Section */}
      <div className="space-y-4">
        <div className="flex items-center justify-between">
          <div>
            <h2 className="text-2xl font-semibold tracking-tight">
              Your Assignments
            </h2>
            <p className="text-sm text-muted-foreground">
              Manage and track all your course assignments
            </p>
          </div>
        </div>
        <AssignmentCard />
      </div>
    </div>
  );
}
