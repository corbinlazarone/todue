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
func (c *CourseModel) CreateNewCourseEntry(ctx context.Context, courseName string, userID string) (string, error) {
	statement := `INSERT INTO courses (course_name, user_id) VALUES ($1, $2) RETURNING id`

	var id string
	err := c.DB.QueryRow(ctx, statement, courseName, userID).Scan(&id)

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

func (c *CourseModel) GetAllCourseData(ctx context.Context, userID string) ([]types.CourseData, error) {
	coursesQuery := `SELECT id, course_name FROM courses WHERE user_id = $1`
	courseRows, err := c.DB.Query(ctx, coursesQuery, userID)
	if err != nil {
		return nil, err
	}
	defer courseRows.Close()

	var courses []types.CourseData
	for courseRows.Next() {
		var course types.CourseData
		var courseID string
		if err := courseRows.Scan(&courseID, &course.CourseName); err != nil {
			return nil, err
		}
		course.CourseID = courseID

		assignmentsQuery := `SELECT id, name, description, due_date::text, all_day, color, start_time, end_time, reminder 
			FROM assignments WHERE course_id = $1`
		assignmentRows, err := c.DB.Query(ctx, assignmentsQuery, courseID)
		if err != nil {
			return nil, err
		}

		var assignments []types.Assignment
		for assignmentRows.Next() {
			var a types.Assignment
			if err := assignmentRows.Scan(
				&a.ID,
				&a.Name,
				&a.Description,
				&a.DueDate,
				&a.AllDay,
				&a.Color,
				&a.StartTime,
				&a.EndTime,
				&a.Reminder,
			); err != nil {
				assignmentRows.Close()
				return nil, err
			}
			assignments = append(assignments, a)
		}
		assignmentRows.Close()

		course.Assignments = assignments
		courses = append(courses, course)
	}

	if err := courseRows.Err(); err != nil {
		return nil, err
	}

	return courses, nil
}
