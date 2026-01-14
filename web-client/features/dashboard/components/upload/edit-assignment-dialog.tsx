"use client";

import {
  Dialog,
  DialogTrigger,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogClose,
  DialogDescription,
} from "@/shared/ui/dialog";
import { Button } from "@/shared/ui/button";
import { Label } from "@/shared/ui/label";
import { Input } from "@/shared/ui/input";
import { ReactNode } from "react";
import { Assignment } from "../../types";
import { Textarea } from "@/shared/ui/textarea";

export function EditAssignmentDialog({
  // trigger,
  // onConfirm,
  AssignmentData,
}: {
  // trigger: ReactNode;
  // onConfirm(e: React.FormEvent): void;
  AssignmentData: Assignment;
}) {
  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    // onConfirm(e);
  }

  // TODO: Disply the assignment data they are about to edit

  return (
    <Dialog open={true}>
      {/* <DialogTrigger asChild>{trigger}</DialogTrigger> */}
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Edit Assignment</DialogTitle>
          {/* <DialogDescription></DialogDescription> */}
        </DialogHeader>
        <div className="grid gap-4">
          {/* Name section */}
          <div className="grid gap-3">
            <Label htmlFor="name">Name</Label>
            <Input id="name" name="name" defaultValue={AssignmentData.name} />
          </div>

          {/* Description section */}
          <div className="grid gap-3">
            <Label htmlFor="description">Description</Label>
            <Textarea
              id="description"
              maxLength={50}
              defaultValue={AssignmentData.description}
              className="resize-none"
            />
          </div>

          {/* all day section */}
          <div className="flex items-center justify-between">
            <Label>All Day</Label>
            <div className="flex gap-3">
              <Button variant="outline" type="button">
                Yes
              </Button>
              <Button variant="outline" type="button">
                No
              </Button>
            </div>
          </div>

          {/* Due Date Section */}
          <div className="grid gap-3">
            <Label>Due Date</Label>
          </div>

          {/* Start Time Section */}
          <div className="grid gap-3">
            <Label>Start Time</Label>
          </div>

          {/* End Time Section */}
          <div className="grid gap-3">
            <Label>End Time</Label>
          </div>

          {/* Color Picker Section */}
          <div className="grid gap-3">
            <Label>Color</Label>
          </div>

          {/* Reminder Picker Section */}
          <div className="grid gap-3">
            <Label>Reminder</Label>
          </div>
        </div>
        <form onSubmit={handleSubmit}>
          <DialogFooter>
            <DialogClose asChild>
              <Button variant="outline" type="button">
                Cancel
              </Button>
            </DialogClose>
            <DialogClose asChild>
              <Button type="submit">Confirm</Button>
            </DialogClose>
          </DialogFooter>
        </form>
      </DialogContent>
    </Dialog>
  );
}
