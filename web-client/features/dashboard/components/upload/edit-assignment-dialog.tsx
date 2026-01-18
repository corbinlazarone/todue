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

export function EditAssignmentDialog({
  trigger,
  onConfirm,
  AssignmentData,
}: {
  trigger: ReactNode;
  onConfirm(e: React.FormEvent): void;
  AssignmentData: Assignment;
}) {
  const [startCalOpen, setStartCalOpen] = useState(false);
  const [startDate, setStartDate] = useState<Date | undefined>(undefined);

  const [endCalOpen, setEndCalOpen] = useState(false);
  const [endDate, setEndDate] = useState<Date | undefined>(undefined);

  const [allDay, setAllDay] = useState(AssignmentData.all_day);

  function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    onConfirm(e);
  }

  // NOTE: going to use this method for converting user's chose date and time to their specfic timezone.
  // import { format } from 'date-fns';
  // import { utcToZonedTime, zonedTimeToUtc } from 'date-fns-tz';
  // export function combineToRFC3339WithTimezone(dateStr: string, timeStr: string, timezone: string): string {
  //   const dateTimeStr = `${dateStr}T${timeStr}:00`;
  //   const localDate = new Date(dateTimeStr);
  //   const zonedDate = utcToZonedTime(localDate, timezone);
  //   return format(zonedDate, "yyyy-MM-dd'T'HH:mm:ssXXX", { timeZone: timezone });
  // }

  return (
    <Dialog open={true}>
      <DialogTrigger asChild>{trigger}</DialogTrigger>
      <DialogContent className="sm:max-w-[425px]">
        <DialogHeader>
          <DialogTitle>Edit Assignment</DialogTitle>
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

          {/* start date and time */}
          <div className="flex gap-4">
            <div className="flex flex-col gap-3">
              <Label htmlFor="start-date-picker">Start Date</Label>
              <Popover open={startCalOpen} onOpenChange={setStartCalOpen}>
                <PopoverTrigger asChild>
                  <Button
                    variant="outline"
                    id="start-date-picker"
                    disabled={allDay}
                  >
                    {startDate ? startDate.toLocaleDateString() : "Select date"}
                    <ChevronDown />
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar
                    mode="single"
                    selected={startDate}
                    captionLayout="dropdown"
                    onSelect={(date: Date | undefined) => {
                      setStartDate(date);
                      setStartCalOpen(false);
                    }}
                  />
                </PopoverContent>
              </Popover>
            </div>
            <div className="flex flex-col gap-3">
              <Label htmlFor="start-time-picker">Start Time</Label>
              <Input
                type="time"
                id="start-time-picker"
                step="1"
                defaultValue="10:30:00"
                disabled={allDay}
                className="bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden [&::-webkit-calendar-picker-indicator]:appearance-none"
              />
            </div>
          </div>

          {/* end date and time */}
          <div className="flex gap-4">
            <div className="flex flex-col gap-3">
              <Label htmlFor="end-date-picker">End Date</Label>
              <Popover open={endCalOpen} onOpenChange={setEndCalOpen}>
                <PopoverTrigger asChild>
                  <Button
                    variant="outline"
                    id="end-date-picker"
                    disabled={allDay}
                  >
                    {endDate ? endDate.toLocaleDateString() : "Select date"}
                    <ChevronDown />
                  </Button>
                </PopoverTrigger>
                <PopoverContent>
                  <Calendar
                    mode="single"
                    selected={startDate}
                    captionLayout="dropdown"
                    onSelect={(date: Date | undefined) => {
                      setEndDate(date);
                      setEndCalOpen(false);
                    }}
                  />
                </PopoverContent>
              </Popover>
            </div>
            <div className="flex flex-col gap-3">
              <Label htmlFor="end-time-picker">End Time</Label>
              <Input
                type="time"
                id="end-time-picker"
                step="1"
                defaultValue="10:30:00"
                disabled={allDay}
                className="bg-background appearance-none [&::-webkit-calendar-picker-indicator]:hidden [&::-webkit-calendar-picker-indicator]:appearance-none"
              />
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
