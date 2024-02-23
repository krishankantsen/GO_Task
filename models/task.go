package models

type Task struct {
	ID          int64
	Title       string
	Description string
	Status      string
	DueDate     string
}
