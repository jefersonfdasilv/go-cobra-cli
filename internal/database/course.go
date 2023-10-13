package database

import (
	"database/sql"
	"github.com/google/uuid"
)

type Course struct {
	db          *sql.DB
	ID          string
	Name        string
	Description string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{
		db: db,
	}
}

func (c *Course) CreateCourse(name, description string, categoryId string) (Course, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO courses (id, name, description, category_id) VALUES ($1, $2, $3, $4)", id, name, description, categoryId)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description}, nil
}

func (c *Course) GetCourse(id string) (Course, error) {
	var course Course
	err := c.db.QueryRow("SELECT id, name, description FROM courses WHERE id = $1", id).
		Scan(&course.ID, &course.Name, &course.Description)

	if err != nil {
		return Course{}, err
	}

	return course, nil
}

func (c *Course) GetCourses() ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses")

	if err != nil {
		return nil, err
	}

	var courses []Course

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description)

		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func (c *Course) GetCoursesByCategoryID(categoryId string) ([]Course, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM courses WHERE category_id = $1", categoryId)

	if err != nil {
		return nil, err
	}

	var courses []Course

	for rows.Next() {
		var course Course
		err := rows.Scan(&course.ID, &course.Name, &course.Description)

		if err != nil {
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}
