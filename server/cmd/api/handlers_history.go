package main

import (
	"encoding/json"
	"net/http"

	"github.com/corbinlazarone/todue/cmd/internals/types"
)

func (app *application) historyHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var rep types.Response

	userID, ok := ctx.Value("userID").(string)
	if !ok {
		app.errLog.Println("USER ID NOT FOUND IN REQUEST CONTEXT")
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	courseData, err := app.courses.GetAllCourseData(ctx, userID)
	if err != nil {
		app.errLog.Println(err)
		rep.WriteErrorResponse(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(courseData)
}
