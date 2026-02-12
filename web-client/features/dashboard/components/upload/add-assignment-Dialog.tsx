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

export function AddNewAssignmentDialog({
  trigger,
  onConfirm,
}: {
  trigger: ReactNode;
  onConfirm(data: Assignment): void;
}) {
  const [calOpen, setCalOpen] = useState(false);

  // Assignment value state
  const [allDay, setAllDay] = useState<boolean>(true);
  const [dueDate, setDueDate] = useState<Date | undefined>(
    new Date(),
  );
  const [color, setColor] = useState<string>("");
  const [reminder, setReminder] = useState<number>(0);
  const [name, setName] = useState<string>("");
  const [description, setDescription] = useState<string>("");
  const [startTime, setStartTime] = useState<string>("");
  const [endTime, setEndTime] = useState<string>("");

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();

    const newAssignment: Assignment = {
      id: 0, // corrected this in handle function 
      name: name,
      description: description,
      all_day: allDay,
      due_date: dueDate ? dueDate.toISOString().split("T")[0] : "",
      start_time: allDay ? "" : startTime,
      end_time: allDay ? "" : endTime,
      color: color,
      reminder: reminder,
    };

    onConfirm(newAssignment);
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
              <Input
                id="name"
                name="name"
                value={name}
                onChange={(e) => setName(e.target.value)}
              />
            </div>

            {/* Description section */}
            <div className="grid gap-3">
              <Label htmlFor="description">Description</Label>
              <Textarea
                id="description"
                maxLength={50}
                value={description}
                onChange={(e) => setDescription(e.target.value)}
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
                  value={startTime}
                  onChange={(e) => setStartTime(e.target.value)}
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
                  value={endTime}
                  onChange={(e) => setEndTime(e.target.value)}
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
