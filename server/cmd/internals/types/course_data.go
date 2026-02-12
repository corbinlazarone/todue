package types

type Assignment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	AllDay      bool   `json:"all_day"`
	Color       string `json:"color"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Reminder    int    `json:"reminder"`
}

type CourseData struct {
	CourseID    int          `json:"course_id"`
	CourseName  string       `json:"course_name"`
	Assignments []Assignment `json:"assignments"`
}
