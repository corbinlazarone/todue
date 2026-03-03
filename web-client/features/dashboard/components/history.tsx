"use client";

import { useEffect, useState } from "react";
import {
  Calendar,
  Clock,
  ChevronDown,
  ChevronUp,
  AlertCircle,
  FileText,
} from "lucide-react";
import { Card, CardContent } from "@/shared/ui/card";
import { Skeleton } from "@/shared/ui/skeleton";
import { format, parseISO } from "date-fns";
import { toZonedTime } from "date-fns-tz";
import { toast } from "sonner";
import { fetchCourseHistory } from "../api/history/actions";
import { CourseData } from "../types";

export default function History() {
  const [timezone, setTimezone] = useState<string>("America/New_York");
  const [expandedCourse, setExpandedCourse] = useState<string | null>(null);
  const [loading, setLoading] = useState<boolean>(true);
  const [courseHistory, setCourseHistory] = useState<CourseData[]>([]);

  useEffect(() => {
    const detectedTimezone = Intl.DateTimeFormat().resolvedOptions().timeZone;
    setTimezone(detectedTimezone);
  }, []);

  // TODO: replace this flow with something like tanstack react query.
  useEffect(() => {
    const CourseHistory = async () => {
      try {
        const result = await fetchCourseHistory();
        setCourseHistory(result);
      } catch {
        toast.error("Failed to fetch course history. Try again or contact us.");
        return;
      } finally {
        setLoading(false);
      }
    }
    CourseHistory();
  }, []);

  const formatDate = (date: string) => {
    return new Date(date).toLocaleDateString("en-US", {
      month: "short",
      day: "numeric",
      year: "numeric",
    });
  };

  const formatTime = (time: string) => {
    if (!time) return "";
    let date: Date;
    if (time.includes("T")) {
      date = parseISO(time);
    } else {
      date = new Date(`2000-01-01T${time}`);
    }
    const zonedDate = toZonedTime(date, timezone);
    return format(zonedDate, "h:mm a");
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
    <div className="w-full max-w-7xl mx-auto px-4 py-6">
      {loading ? (
        <div className="space-y-4">
          {[1, 2, 3].map((i) => (
            <Card key={i} className="p-4">
              <div className="flex items-center justify-between">
                <div className="space-y-2">
                  <Skeleton className="h-5 w-48" />
                  <Skeleton className="h-4 w-24" />
                </div>
                <Skeleton className="h-5 w-5 rounded-full" />
              </div>
            </Card>
          ))}
        </div>
      ) : courseHistory.length === 0 ? (
        <Card className="flex flex-col items-center justify-center p-8">
          <FileText className="h-12 w-12 text-muted-foreground mb-3" />
          <h3 className="text-lg font-medium mb-1">No Upload History</h3>
          <p className="text-sm text-muted-foreground text-center">
            Your uploaded course history will appear here
          </p>
        </Card>
      ) : (
        <div className="space-y-4">
          {courseHistory.map((course) => (
            <div
              key={course.course_id}
            >
              <Card className="overflow-hidden">
                <div
                  className="p-4 cursor-pointer hover:bg-muted/50 transition-colors"
                  onClick={() =>
                    setExpandedCourse(
                      expandedCourse === course.course_id
                        ? null
                        : course.course_id
                    )
                  }
                >
                  <div className="flex items-center justify-between">
                    <div>
                      <h3 className="font-medium">{course.course_name}</h3>
                      <p className="text-sm text-muted-foreground">
                        {course.assignments.length} assignments
                      </p>
                    </div>
                    {expandedCourse === course.course_id ? (
                      <ChevronUp className="h-5 w-5 text-muted-foreground" />
                    ) : (
                      <ChevronDown className="h-5 w-5 text-muted-foreground" />
                    )}
                  </div>
                </div>

                {expandedCourse === course.course_id && (
                  <CardContent className="border-t p-4 grid grid-cols-1 md:grid-cols-2 xl:grid-cols-3 gap-4">
                    {course.assignments.map((assignment) => (
                      <Card key={assignment.id} className="p-4">
                        <div className="flex items-center space-x-2 mb-3">
                          <div
                            className="h-2.5 w-2.5 rounded-full"
                            style={{ backgroundColor: assignment.color }}
                          />
                          <h4 className="font-medium">{assignment.name}</h4>
                        </div>

                        {assignment.description && (
                          <p className="text-sm text-muted-foreground mb-3 line-clamp-2">
                            {assignment.description}
                          </p>
                        )}

                        <div className="space-y-1.5 text-sm">
                          <div className="flex items-center">
                            <Calendar className="w-3.5 h-3.5 mr-1.5 text-muted-foreground" />
                            {formatDate(assignment.due_date)}
                          </div>
                          <div className="flex items-center">
                            <Clock className="w-3.5 h-3.5 mr-1.5 text-muted-foreground" />
                            {formatTime(assignment.start_time)} -{" "}
                            {formatTime(assignment.end_time)}
                          </div>
                          <div className="flex items-center">
                            <AlertCircle className="w-3.5 h-3.5 mr-1.5 text-muted-foreground" />
                            {formatReminder(assignment.reminder)}
                          </div>
                        </div>
                      </Card>
                    ))}
                  </CardContent>
                )}
              </Card>
            </div>
          ))}
        </div>
      )}
    </div>
  );
}
