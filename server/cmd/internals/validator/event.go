package validator

// Validator user's and AI's Google Calendar events.

type Event struct {
	AssignmentID   int    `json:"assignment_id"`
	AssignmentName string `json:"assignment_name"`
	DueDate        string `json:"due_date"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Reminder       int    `json:"reminder"`
	Color          string `json:"color"`
}

type EventError struct {
	AssignmentID   int    `json:"assignment_id"`
	AssignmentName string `json:"assignment_name"`
	Error          string `json:"error"`
}

// TODO: validate that due date is valid date time string to match google calendar API requirements
// TODO: validate that start time is valid time string to match google calendar API requirements
// TODO: validate that end time is valid time string to match google calendar API requirements
// TODO: validate that reminder is valid integer to match google calendar API requirements
// TODO: validate that color is valid hex color to match google calendar API requirements

func (e *Event) Validate() []EventError {
	return []EventError{
		{
			AssignmentID:   1,
			AssignmentName: "Assignment 1",
			Error:          "Due date is invalid",
		},
		{
			AssignmentID:   2,
			AssignmentName: "Assignment 2",
			Error:          "Start time is invalid",
		},
		{
			AssignmentID:   3,
			AssignmentName: "Assignment 3",
			Error:          "End time is invalid",
		},
		{
			AssignmentID:   4,
			AssignmentName: "Assignment 4",
			Error:          "Reminder is invalid",
		},
		{
			AssignmentID:   5,
			AssignmentName: "Assignment 5",
			Error:          "Color is invalid",
		},
	}
}
