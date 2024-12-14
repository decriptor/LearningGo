package persistence

import "taskmanager2/internal/app"

type DataStore interface {
	LoadTasks() ([]app.Task, int, error)
	SaveTasks([]app.Task, int) error
}
