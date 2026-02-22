export type Assignment = {
  id: string;
  name: string;
  description: string;
  due_date: string;
  all_day: boolean;
  color: string;
  start_time: string;
  end_time: string;
  reminder: number;
};

export type CourseData = {
  course_id: string;
  course_name: string;
  assignments: Assignment[];
};

export type Courses = {
  courses: CourseData[];
};

export type EventError = {
  assignment_id: number;
  assignment_name: string;
  errors: string[];
};

export const reminderOptions = [
  { value: "0", label: "At time of event" },
  { value: "5", label: "5 minutes before" },
  { value: "10", label: "10 minutes before" },
  { value: "15", label: "15 minutes before" },
  { value: "30", label: "30 minutes before" },
  { value: "60", label: "1 hour before" },
  { value: "120", label: "2 hours before" },
  { value: "180", label: "3 hours before" },
  { value: "360", label: "6 hours before" },
  { value: "720", label: "12 hours before" },
  { value: "1440", label: "1 day before" },
  { value: "2880", label: "2 days before" },
  { value: "4320", label: "3 days before" },
  { value: "5760", label: "4 days before" },
  { value: "7200", label: "5 days before" },
  { value: "8640", label: "6 days before" },
  { value: "10080", label: "1 week before" },
  { value: "20160", label: "2 weeks before" },
  { value: "30240", label: "3 weeks before" },
  { value: "40320", label: "4 weeks before (max)" },
];

export const colorOptions = [
  { value: "#7986cb", label: "Blue" },
  { value: "#33b679", label: "Green" },
  { value: "#8e24aa", label: "Purple" },
  { value: "#e67c73", label: "Red" },
  { value: "#f6c026", label: "Yellow" },
  { value: "#f5511d", label: "Orange" },
  { value: "#039be5", label: "Turquoise" },
  { value: "#616161", label: "Gray" },
  { value: "#3f51b5", label: "Bold Blue" },
  { value: "#0b8043", label: "Bold Green" },
  { value: "#d60000", label: "Bold Red" },
];
