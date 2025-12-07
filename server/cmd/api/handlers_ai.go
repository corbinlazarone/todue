package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/corbinlazarone/Todue-Actual/cmd/internals/models"
	"github.com/corbinlazarone/Todue-Actual/cmd/internals/opencode"
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

	if len(params.PDFText) == 0 {
		app.errLog.Println("PDF text is required")
		rep.WriteErrorResponse(w, http.StatusBadRequest, "PDF text is required")
		return
	}

	if err := sanitizePDFText(params.PDFText); err != nil {
		app.errLog.Printf("Input sanitization failed: %v", err)
		rep.WriteErrorResponse(w, http.StatusBadRequest, "Invalid input content")
		return
	}

	aiReq := opencode.OpenCodeRequest{
		MaxTokens: 4000,
		SystemMessage: `You are a precise course information extraction assistant. Extract only explicitly stated information from the syllabus. Follow these rules:
1. Only extract assignments, exams, and deadlines that have specific dates
2. Ensure dates are in YYYY-MM-DD format
3. Convert all times to 24-hour format (HH:mm)
4. If no specific time is mentioned, use "23:59" as the default
5. If no specific end time is mentioned, make it the same as start time
6. Assign appropriate colors from this list: #7986cb, #33b679, #8e24aa, #e67c73, #f6c026, #f5511d, #039be5, #3f51b5, #0b8043, #d60000
7. Set default reminder to 1440 (1 day) if not specified
8. Ensure each assignment has a unique ID
Do not infer or generate any data not directly present in the source text.`,
		Message: []opencode.Message{
			{
				Role: "user",
				Content: fmt.Sprintf(`Extract course information and assignments from this syllabus in this exact format:

{
  "courses": [
    {
      "course_id": 1,
      "course_name": "Course Name",
      "assignments": [
        {
          "id": 1,
          "name": "Assignment Name",
          "description": "Description",
          "due_date": "YYYY-MM-DD",
          "color": "#hexcolor",
          "start_time": "HH:mm",
          "end_time": "HH:mm",
          "reminder": 1440
        }
      ]
    }
  ]
}

Syllabus text:
%s`, params.PDFText),
			},
		},
	}

	aiResp, err := app.opencodeClient.CreateMessage(aiReq)
	if err != nil {
		app.errLog.Printf("AI extraction failed: %v", err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to extract course data")
		return
	}

	// Parse the JSON response into our expected format
	var courseResponse struct {
		Courses []CourseData `json:"courses"`
	}

	err = json.Unmarshal([]byte(aiResp.JSONResponse), &courseResponse)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Failed to parse AI response")
		return
	}

	// Return the extracted course data
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courseResponse)
}
