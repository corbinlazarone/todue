package validator

import (
	"errors"
	"slices"
	"time"
)

// Validator user's and AI's Google Calendar events.

type Event struct {
	AssignmentID   int    `json:"assignment_id"`
	AssignmentName string `json:"assignment_name"`
	DueDate        string `json:"due_date"`
	AllDay         bool   `json:"all_day"`
	StartTime      string `json:"start_time"`
	EndTime        string `json:"end_time"`
	Reminder       int    `json:"reminder"`
	Color          string `json:"color"`
}

type EventError struct {
	AssignmentID   int      `json:"assignment_id"`
	AssignmentName string   `json:"assignment_name"`
	Errors         []string `json:"errors"`
}

func (e *Event) validateDueDate() error {
	res := checkFormat(e.DueDate, time.DateOnly)

	if !res && e.AllDay {
		return errors.New("Due date does not match expected 'YYYY-MM-DD' format")
	}

	if res && !e.AllDay {
		return errors.New("All day event is not marked")
	}

	return nil
}

func (e *Event) validateDateTimes() error {
	resStart := checkFormat(e.StartTime, time.RFC3339)
	if !resStart {
		return errors.New("Start time does not match expected 'YYYY-MM-DDTHH:MM:SSZ' format")
	}

	resEnd := checkFormat(e.EndTime, time.RFC3339)
	if !resEnd {
		return errors.New("Start time does not match expected 'YYYY-MM-DDTHH:MM:SSZ' format")
	}

	return nil
}

func (e *Event) validateReminder() error {
	reminderOptions := []int{
		0,     // At time of event
		5,     // 5 minutes before
		10,    // 10 minutes before
		15,    // 15 minutes before
		30,    // 30 minutes before
		60,    // 1 hour before
		120,   // 2 hours before
		180,   // 3 hours before
		360,   // 6 hours before
		720,   // 12 hours before
		1440,  // 1 day before
		2880,  // 2 days before
		4320,  // 3 days before
		5760,  // 4 days before
		7200,  // 5 days before
		8640,  // 6 days before
		10080, // 1 week before
		20160, // 2 weeks before
		30240, // 3 weeks before
		40320, // 4 weeks before (max)
	}

	if !slices.Contains(reminderOptions, e.Reminder) {
		return errors.New("Reminder is not valid")
	}
	return nil
}

func (e *Event) validateColor() error {
	colorOptions := map[string]string{
		"#7986cb": "1",  // Blue
		"#33b679": "2",  // Green
		"#8e24aa": "3",  // Purple
		"#e67c73": "4",  // Red
		"#f6c026": "5",  // Yellow
		"#f5511d": "6",  // Orange
		"#039be5": "7",  // Turquoise
		"#616161": "8",  // Gray
		"#3f51b5": "9",  // Bold Blue
		"#0b8043": "10", // Bold Green
		"#d60000": "11", // Bold Red
	}

	_, ok := colorOptions[e.Color]
	if !ok {
		return errors.New("Color is not valid")
	}

	return nil
}

func (e *Event) Validate() []EventError {
	var errs []EventError

	err := e.validateDueDate()
	if err != nil {
		errs = append(errs, EventError{
			AssignmentID:   e.AssignmentID,
			AssignmentName: e.AssignmentName,
			Errors:         []string{err.Error()},
		})
	}

	err = e.validateDateTimes()
	if err != nil {
		errs = append(errs, EventError{
			AssignmentID:   e.AssignmentID,
			AssignmentName: e.AssignmentName,
			Errors:         []string{err.Error()},
		})
	}

	err = e.validateReminder()
	if err != nil {
		errs = append(errs, EventError{
			AssignmentID:   e.AssignmentID,
			AssignmentName: e.AssignmentName,
			Errors:         []string{err.Error()},
		})
	}

	err = e.validateColor()
	if err != nil {
		errs = append(errs, EventError{
			AssignmentID:   e.AssignmentID,
			AssignmentName: e.AssignmentName,
			Errors:         []string{err.Error()},
		})
	}

	return errs
}

func checkFormat(current, expected string) bool {
	_, err := time.Parse(expected, current)
	if err != nil {
		return false
	}
	return true
}
