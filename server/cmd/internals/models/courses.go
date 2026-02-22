package models

import (
	"context"

	"github.com/corbinlazarone/todue/cmd/internals/types"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CourseModel struct {
	DB *pgxpool.Pool
}

// Creates a course entry and returns the created UUID
func (c *CourseModel) CreateNewCourseEntry(ctx context.Context, courseName string) (string, error) {
	statement := `INSERT INTO courses (course_name) VALUES ($1) RETURNING id`

	var id string
	err := c.DB.QueryRow(ctx, statement, courseName).Scan(&id)

	if err != nil {
		return "", err
	}

	return id, nil
}

// Creates a new assignment for the course matching the provided course id
func (c *CourseModel) CreateNewAssignment(ctx context.Context, courseId string, assignment types.Assignment) error {
	statement :=
		`INSERT INTO assignments (
		course_id,
		name,
		description,
		due_date,
		all_day,
		color,
		start_time,
		end_time,
		reminder
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5,
		$6,
		$7,
		$8,
		$9
	)`

	_, err := c.DB.Exec(ctx, statement,
		courseId,
		assignment.Name,
		assignment.Description,
		assignment.DueDate,
		assignment.AllDay,
		assignment.Color,
		assignment.StartTime,
		assignment.EndTime,
		assignment.Reminder,
	)

	if err != nil {
		return err
	}

	return nil
}
