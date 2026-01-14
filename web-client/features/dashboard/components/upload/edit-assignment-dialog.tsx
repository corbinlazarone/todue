"use client";

import { Button } from "@/shared/ui/button";
import { Label } from "@/shared/ui/label";
import { Input } from "@/shared/ui/input";

import { Assignment, colorOptions, reminderOptions } from "../../types";
import { Textarea } from "@/shared/ui/textarea";

import {
  Dialog,
  DialogContent,
  DialogHeader,
  DialogTitle,
  DialogFooter,
  DialogClose,
} from "@/shared/ui/dialog";

import {
  Select,
  SelectTrigger,
  SelectValue,
  SelectContent,
  SelectGroup,
  SelectItem,
} from "@/shared/ui/select";

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

          {/* Color Picker Section */}
          <div className="grid gap-3">
            <Label>Color</Label>
            <Select>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select a color" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  {colorOptions.map((c) => (
                    <SelectItem key={c.value} value={c.value}>
                      <div className="flex items-center gap-2">
                        <div
                          className="w-4 h-4 rounded-full"
                          style={{ backgroundColor: c.value }}
                          aria-label={c.label}
                          title={c.label}
                        />
                        <span className="text-sm">{c.label}</span>
                      </div>
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
          </div>

          {/* Reminder Picker Section */}
          <div className="grid gap-3">
            <Label>Reminder</Label>
            <Select>
              <SelectTrigger className="w-full">
                <SelectValue placeholder="Select a reminder time" />
              </SelectTrigger>
              <SelectContent>
                <SelectGroup>
                  {reminderOptions.map((r) => (
                    <SelectItem key={r.value} value={r.value}>
                      {r.label}
                    </SelectItem>
                  ))}
                </SelectGroup>
              </SelectContent>
            </Select>
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
