export type Assignment = {
  id: number;
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
  course_id: number;
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
