package main

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/corbinlazarone/todue/cmd/internals/models"
	"github.com/corbinlazarone/todue/cmd/internals/services"
	"github.com/corbinlazarone/todue/cmd/internals/types"
	"github.com/corbinlazarone/todue/cmd/internals/validator"
	"google.golang.org/api/calendar/v3"
)

func (app *application) insertCourseDataHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	defer r.Body.Close()

	var rep types.Response

	type parameters struct {
		UserTimeZone string             `json:"timezone"`
		Courses      []types.CourseData `json:"courses"`
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
						Error:          []string{err.Error()},
					})
				}
			} else {
				_, err := ToRFC3339(assVal.DueDate, assVal.StartTime, params.UserTimeZone)
				if err != nil {
					allErrors = append(allErrors, types.TimeConversionError{
						AssignmentID:   assVal.ID,
						AssignmentName: assVal.Name,
						Error:          []string{err.Error()},
					})
				}

				_, err = ToRFC3339(assVal.DueDate, assVal.EndTime, params.UserTimeZone)
				if err != nil {
					allErrors = append(allErrors, types.TimeConversionError{
						AssignmentID:   assVal.ID,
						AssignmentName: assVal.Name,
						Error:          []string{err.Error()},
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

			err := addEventToCalendar(
				r.Context(),
				app.users,
				event,
				assVal.AllDay,
			)
			if err != nil {
				app.errLog.Println(err)
				rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
	}

	for _, val := range params.Courses {
		id, err := app.courses.CreateNewCourseEntry(ctx, val.CourseName)
		if err != nil {
			app.errLog.Println(err)
			rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		for _, assVal := range val.Assignments {
			err := app.courses.CreateNewAssignment(ctx, id, assVal)
			if err != nil {
				app.errLog.Println(err)
				rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
				return
			}
		}
	}

	rep.WriteSuccessResponse(w, "Course data has been inserted successfully", http.StatusOK)
}

func addEventToCalendar(
	ctx context.Context,
	userModel *models.UserModel,
	event *types.CalendarEvent,
	allDay bool,
) error {

	userID := ctx.Value("userID").(string)

	user, err := userModel.GetByID(ctx, userID)
	if err != nil {
		return err
	}

	if user.GoogleRefreshToken == "" {
		return errors.New("user's google refresh token is empty")
	}

	calendarSrv, err := services.GoogleCalendarService(
		ctx,
		userModel,
		userID,
		user.GoogleRefreshToken,
	)
	if err != nil {
		return err
	}

	var newEvent *calendar.Event

	if !allDay {
		newEvent = &calendar.Event{
			Summary: event.Summary,

			Description: event.Description,

			Start: &calendar.EventDateTime{
				DateTime: event.Start.DateTime,

				TimeZone: event.Start.TimeZone,
			},
			End: &calendar.EventDateTime{
				DateTime: event.End.DateTime,
				TimeZone: event.End.TimeZone,
			},
			ColorId: ConvertToColorID(event.ColorId),
			Reminders: &calendar.EventReminders{
				UseDefault: false,
				Overrides: []*calendar.EventReminder{
					{
						Method:  "email",
						Minutes: int64(event.Reminders.Overrides[0].Minutes),
					},
				},
				ForceSendFields: []string{"UseDefault"},
			},
		}
	} else {
		newEvent = &calendar.Event{
			Summary:     event.Summary,
			Description: event.Description,
			Start: &calendar.EventDateTime{
				Date: event.Start.Date,
			},

			End: &calendar.EventDateTime{
				Date: event.End.Date,
			},
			ColorId: ConvertToColorID(event.ColorId),
			Reminders: &calendar.EventReminders{
				UseDefault: false,
				Overrides: []*calendar.EventReminder{
					{
						Method:  "email",
						Minutes: int64(event.Reminders.Overrides[0].Minutes),
					},
				},
				ForceSendFields: []string{"UseDefault"},
			},
		}
	}

	_, err = calendarSrv.Events.Insert("primary", newEvent).
		Context(ctx).
		Do()

	if err != nil {
		return err
	}

	return nil
}
