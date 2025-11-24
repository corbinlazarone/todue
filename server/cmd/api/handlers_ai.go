package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/cmovie/cmd/internals/models"
)

type Assignment struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
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

func (app *application) ExtractCourseData(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	// Limit request body size to 5MB
	r.Body = http.MaxBytesReader(w, r.Body, 5<<20)

	var rep models.Response

	type parameters struct {
		PDFText string `json:"pdfText"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		app.errLog.Println(err)
		if err.Error() == "http: request body too large" {
			rep.WriteErrorResponse(w, http.StatusRequestEntityTooLarge, "Request body exceeds 5MB limit")
			return
		}
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Invalid request body")
		return
	}

	// TODO: do claude extraction
}
