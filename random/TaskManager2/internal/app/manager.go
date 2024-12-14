package app

import "fmt"

type TaskManager struct {
	Tasks  []Task
	NextID int
}

func NewTaskManager() *TaskManager {
	return &TaskManager{}
}

func (tm *TaskManager) AddTask(title string) Task {
	task := Task{
		ID:        tm.NextID,
		Title:     title,
		Completed: false,
	}
	tm.Tasks = append(tm.Tasks, task)
	tm.NextID++
	return task
}

func (tm *TaskManager) ListTasks() []Task {
	return tm.Tasks
}

func (tm *TaskManager) CompleteTask(id int) error {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i].Completed = true
			return nil
		}
	}
	return fmt.Errorf("task with ID %d not found", id)
}
