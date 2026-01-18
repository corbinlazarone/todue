"use client";

import { Button } from "@/shared/ui/button";
import { Label } from "@/shared/ui/label";
import { Input } from "@/shared/ui/input";
import { Calendar } from "@/shared/ui/calendar";
import { Assignment, colorOptions, reminderOptions } from "../../types";
import { Textarea } from "@/shared/ui/textarea";
import { Popover, PopoverContent, PopoverTrigger } from "@/shared/ui/popover";
import { ChevronDown } from "lucide-react";
import { ReactNode, useState } from "react";
import { DialogTrigger } from "@/shared/ui/dialog";

import { format } from "date-fns";
import { toZonedTime } from "date-fns-tz";
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
  trigger,
  timezone,
  onConfirm,
  AssignmentData,
}: {
  trigger: ReactNode;
  timezone: string;
  onConfirm(data: Assignment): void;
  AssignmentData: Assignment;
}) {
  const [allDay, setAllDay] = useState(AssignmentData.all_day);
  const [calOpen, setCalOpen] = useState(false);
  const [dueDate, setDueDate] = useState<Date | undefined>(
    new Date(AssignmentData.due_date),
  );
  const [color, setColor] = useState(AssignmentData.color);
  const [reminder, setReminder] = useState(AssignmentData.reminder);

  function combineToRFC3339WithTimezone(date: Date, timeStr: string): string {
    const dateStr = format(date, "yyyy-MM-dd");
    const dateTimeStr = `${dateStr}T${timeStr}:00`;
    const localDate = new Date(dateTimeStr);
    const zonedDate = toZonedTime(localDate, timezone);
    return format(zonedDate, "yyyy-MM-dd'T'HH:mm:ssXXX");
  }

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    const formData = new FormData(e.target as HTMLFormElement);

    const updatedAssignment: Assignment = {
      id: AssignmentData.id,
      name: formData.get("name") as string,
      description: formData.get("description") as string,
      all_day: allDay,
      due_date: dueDate ? dueDate.toISOString().split("T")[0] : "",
      start_time: allDay
        ? ""
        : combineToRFC3339WithTimezone(
            dueDate!,
            formData.get("start-time-picker") as string,
          ),
      end_time: allDay
        ? ""
        : combineToRFC3339WithTimezone(
            dueDate!,
            formData.get("end-time-picker") as string,
          ),
      color: color,
      reminder: reminder,
    };

    console.log(
      "UPDATED ASSIGNMENT IN DIALOG: ",
      JSON.stringify(updatedAssignment, null, 2),
    );

    onConfirm(updatedAssignment);
  }

  return (
    <Dialog>
      <DialogTrigger asChild>{trigger}</DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Edit Assignment</DialogTitle>
        </DialogHeader>
        <form onSubmit={handleSubmit}>
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
            <div className="grid gap-3">
              <Label>All Day</Label>
              <div className="flex gap-3">
                <Button
                  variant={allDay ? "default" : "outline"}
                  type="button"
                  onClick={() => setAllDay(true)}
                >
                  Yes
                </Button>
                <Button
                  variant={!allDay ? "default" : "outline"}
                  type="button"
                  onClick={() => setAllDay(false)}
                >
                  No
                </Button>
              </div>
            </div>

            {/* due date */}
            <div className="flex gap-4">
              <div className="flex flex-col gap-3">
                <Label htmlFor="due-date">Due Date</Label>
                <Popover open={calOpen} onOpenChange={setCalOpen}>
                  <PopoverTrigger asChild>
                    <Button
                      variant="outline"
                      id="due-date-picker"
                      disabled={false}
                    >
                      {dueDate ? dueDate.toLocaleDateString() : "Select date"}
                      <ChevronDown />
                    </Button>
                  </PopoverTrigger>
                  <PopoverContent>
                    <Calendar
                      mode="single"
                      selected={dueDate}
                      captionLayout="dropdown"
                      onSelect={(date: Date | undefined) => {
                        setDueDate(date);
                        setCalOpen(false);
                      }}
                    />
                  </PopoverContent>{" "}
                </Popover>
              </div>
            </div>

            {/* start time */}
            <div className="flex gap-4">
              <div className="flex flex-col gap-3">
                <Label htmlFor="start-time-picker">Start Time</Label>
                <Input
                  type="time"
                  id="start-time-picker"
                  name="start-time-picker"
                  step="1"
                  defaultValue={AssignmentData.start_time}
                  disabled={allDay}
                  className="bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden [&::-webkit-calendar-picker-indicator]:appearance-none"
                />
              </div>
            </div>

            {/* end date and time */}
            <div className="flex gap-4">
              <div className="flex flex-col gap-3">
                <Label htmlFor="end-time-picker">End Time</Label>
                <Input
                  type="time"
                  id="end-time-picker"
                  name="end-time-picker"
                  step="1"
                  defaultValue={AssignmentData.end_time}
                  disabled={allDay}
                  className="bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden [&::-webkit-calendar-picker-indicator]:appearance-none"
                />
              </div>
            </div>

            {/* Color Picker Section */}
            <div className="grid gap-3">
              <Label>Color</Label>
              <Select name="color" value={color} onValueChange={setColor}>
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
              <Select
                name="reminder"
                value={reminder.toString()}
                onValueChange={(value) => setReminder(parseInt(value, 10))}
              >
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
