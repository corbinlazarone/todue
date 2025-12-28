"use client";

import { Button } from "@/shared/ui/button";
import { Assignment } from "../../types";
import {
  Card,
  CardAction,
  CardContent,
  CardDescription,
  CardFooter,
  CardHeader,
  CardTitle,
} from "@/shared/ui/card";
import { AlertCircle, Calendar, Clock, Edit2, Trash2 } from "lucide-react";

export default function AssignmentCard({
  assignments,
  onDelete,
}: {
  assignments: Assignment[];
  onDelete(id: number): void;
}) {
  const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  };

  const formatTime = (time: string) => {
    return new Date(`2000-01-01T${time}`).toLocaleTimeString("en-US", {
      hour: "numeric",
      minute: "2-digit",
      hour12: true,
    });
  };

  const formatReminder = (minutes: number) => {
    if (minutes === 0) {
      return "At time of event";
    }
    if (minutes < 60) {
      return `${minutes} minute${minutes > 1 ? "s" : ""} before`;
    }
    if (minutes < 1440) {
      const hours = minutes / 60;
      return `${hours} hour${hours > 1 ? "s" : ""} before`;
    }
    const days = minutes / 1440;
    return `${days} day${days > 1 ? "s" : ""} before`;
  };

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 xl:grid-cols-4 gap-3">
      {assignments.length > 0 ? (
        assignments.map((assignment) => (
          <div key={assignment.id}>
            <Card className="hover:shadow-md transition-shadow h-full">
              <CardHeader className="pb-3">
                <div className="flex items-center space-x-2">
                  <div
                    className="h-2 w-2 rounded-full flex-shrink-0"
                    style={{ backgroundColor: assignment.color }}
                  />
                  <CardTitle className="text-sm leading-tight line-clamp-1">
                    {assignment.name}
                  </CardTitle>
                </div>
                <CardAction>
                  <div className="flex items-center space-x-0.5">
                    <Button
                      /* onClick={() => handleEdit(assignment)} */
                      variant="ghost"
                      size="icon-sm"
                      className="text-gray-400 hover:text-indigo-600 h-7 w-7"
                    >
                      <Edit2 className="h-3 w-3" />
                    </Button>
                    <Button
                      onClick={() => onDelete(assignment.id)}
                      variant="ghost"
                      size="icon-sm"
                      className="text-gray-400 hover:text-red-600 h-7 w-7"
                    >
                      <Trash2 className="h-3 w-3" />
                    </Button>
                  </div>
                </CardAction>
              </CardHeader>

              {assignment.description && (
                <CardContent className="pt-0 pb-3">
                  <CardDescription className="line-clamp-2 text-xs">
                    {assignment.description}
                  </CardDescription>
                </CardContent>
              )}

              <CardFooter className="flex-col items-start space-y-1 text-xs pt-0">
                <div className="flex items-center text-gray-700">
                  <Calendar className="w-3 h-3 mr-1.5 flex-shrink-0" />
                  <span>{formatDate(assignment.due_date)}</span>
                </div>
                <div className="flex items-center text-gray-700">
                  <Clock className="w-3 h-3 mr-1.5 flex-shrink-0" />
                  <span className="text-xs">
                    {formatTime(assignment.start_time)} -{" "}
                    {formatTime(assignment.end_time)}
                  </span>
                </div>
                <div className="flex items-center text-gray-700">
                  <AlertCircle className="w-3 h-3 mr-1.5 flex-shrink-0" />
                  <span>{formatReminder(assignment.reminder)}</span>
                </div>
              </CardFooter>
            </Card>
          </div>
        ))
      ) : (
        <Card className="p-6 text-center border-dashed border-2 border-gray-300 dark:border-gray-600 col-span-full">
          <div className="flex flex-col items-center space-y-3">
            <Calendar className="w-8 h-8 text-gray-400" />
            <div>
              <h4 className="text-sm font-medium text-gray-900 dark:text-gray-100">
                No Assignments Found
              </h4>
              <p className="text-xs text-gray-500 dark:text-gray-400 mt-1">
                No assignments were extracted from this course. You can add your
                own assignments manually.
              </p>
            </div>
          </div>
        </Card>
      )}
    </div>
  );
}
