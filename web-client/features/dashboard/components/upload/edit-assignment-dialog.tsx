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
          <div className="grid gap-3">
            <Label htmlFor="name-1">Name</Label>
            <Input id="name-1" name="name" defaultValue={AssignmentData.name} />
          </div>
          <div className="grid gap-3">
            <Label htmlFor="username-1">Description</Label>
            <Input id="username-1" name="username" defaultValue="@peduarte" />
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
