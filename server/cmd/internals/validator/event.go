package validator

import (
	"errors"
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

// func (e *Event) validateReminder() error {
// 	return nil
// }

// func (e *Event) validateColor() error {
// 	return nil
// }

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

	// err = e.validateReminder()
	// if err != nil {
	// 	errs = append(errs, EventError{
	// 		AssignmentID:   e.AssignmentID,
	// 		AssignmentName: e.AssignmentName,
	// 		Errors:         []string{err.Error()},
	// 	})
	// }

	// err = e.validateColor()
	// if err != nil {
	// 	errs = append(errs, EventError{
	// 		AssignmentID:   e.AssignmentID,
	// 		AssignmentName: e.AssignmentName,
	// 		Errors:         []string{err.Error()},
	// 	})
	// }

	return errs
}

func checkFormat(current, expected string) bool {
	_, err := time.Parse(expected, current)
	if err != nil {
		return false
	}
	return true
}
