package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/models"
	"github.com/corbinlazarone/Todue-Actual/cmd/internals/validator"
)

func (app *application) insertCourseDataHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var rep models.Response

	type parameters struct {
		Courses []CourseData `json:"courses"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	var allValidationErrors []validator.EventError

	for _, val := range params.Courses {
		for _, val := range val.Assignments {
			v := validator.Event{
				AssignmentID:   val.ID,
				AssignmentName: val.Name,
				DueDate:        val.DueDate,
				AllDay:         val.AllDay,
				StartTime:      val.StartTime,
				EndTime:        val.EndTime,
				Reminder:       val.Reminder,
				Color:          val.Color,
			}

			errs := v.Validate()
			allValidationErrors = append(allValidationErrors, errs...)
		}
	}

	if len(allValidationErrors) > 0 {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(allValidationErrors)
		return
	}

	rep.WriteSuccessResponse(w, "Course data has been inserted successfully", http.StatusOK)
}

// TODO: Add function to add event to calendar
func AddEventToCalendar() {}
