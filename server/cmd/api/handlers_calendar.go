package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/types"
	"github.com/corbinlazarone/Todue-Actual/cmd/internals/validator"
)

func (app *application) insertCourseDataHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var rep types.Response

	type parameters struct {
		UserTimeZone string       `json:"timezone"`
		Courses      []CourseData `json:"courses"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	if params.UserTimeZone == "" {
		rep.WriteErrorResponse(w, http.StatusBadRequest, "timezone is required")
		return
	}

	if err := ValidateTimeZone(params.UserTimeZone); err != nil {
		rep.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}

	var allErrors []any

	for _, val := range params.Courses {
		for _, assVal := range val.Assignments {
			v := validator.Event{
				AssignmentID:   assVal.ID,
				AssignmentName: assVal.Name,
				DueDate:        assVal.DueDate,
				AllDay:         assVal.AllDay,
				StartTime:      assVal.StartTime,
				EndTime:        assVal.EndTime,
				Reminder:       assVal.Reminder,
				Color:          assVal.Color,
			}

			validationErrs := v.Validate()
			for _, validationErr := range validationErrs {
				allErrors = append(allErrors, validationErr)
			}

			if assVal.AllDay {
				_, err := AddOneDay(assVal.DueDate)
				if err != nil {
					allErrors = append(allErrors, types.TimeConversionError{
						AssignmentID:   assVal.ID,
						AssignmentName: assVal.Name,
						Error:          fmt.Sprintf("invalid due date format: %v", err),
					})
				}
			} else {
				_, err := ToRFC3339(assVal.DueDate, assVal.StartTime, params.UserTimeZone)
				if err != nil {
					allErrors = append(allErrors, types.TimeConversionError{
						AssignmentID:   assVal.ID,
						AssignmentName: assVal.Name,
						Error:          fmt.Sprintf("start time conversion failed: %v", err),
					})
				}

				_, err = ToRFC3339(assVal.DueDate, assVal.EndTime, params.UserTimeZone)
				if err != nil {
					allErrors = append(allErrors, types.TimeConversionError{
						AssignmentID:   assVal.ID,
						AssignmentName: assVal.Name,
						Error:          fmt.Sprintf("end time conversion failed: %v", err),
					})
				}
			}
		}
	}

	if len(allErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(allErrors)
		return
	}

	for _, val := range params.Courses {
		for _, assVal := range val.Assignments {
			event := types.NewCalendarEvent()
			event.Summary = assVal.Name
			event.Description = assVal.Description

			if assVal.AllDay {
				event.Start.Date = assVal.DueDate
				endDate, _ := AddOneDay(assVal.DueDate)
				event.End.Date = endDate
			} else {
				startTime, _ := ToRFC3339(assVal.DueDate, assVal.StartTime, params.UserTimeZone)
				event.Start.DateTime = startTime
				event.Start.TimeZone = params.UserTimeZone

				endTime, _ := ToRFC3339(assVal.DueDate, assVal.EndTime, params.UserTimeZone)
				event.End.DateTime = endTime
				event.End.TimeZone = params.UserTimeZone
			}

			event.ColorId = assVal.Color
			event.Reminders.Overrides[0].Minutes = assVal.Reminder

			err := addEventToCalendar(event)
			if err != nil {
				app.errLog.Println(err)
				rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
	}

	rep.WriteSuccessResponse(w, "Course data has been inserted successfully", http.StatusOK)
}

func addEventToCalendar(event *types.CalendarEvent) error {
	url := "https://www.googleapis.com/calendar/v3/calendars/primary/events"

	accessToken := ""

	jsonData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("error marshaling event: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		return fmt.Errorf("error creating request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error making request: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response: %w", err)
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("API error (status %d): %s", resp.StatusCode, string(body))
	}

	return nil
}
