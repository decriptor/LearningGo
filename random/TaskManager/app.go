package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a single task
type Task struct {
	ID        int
	Title     string
	Completed bool
}

// TaskManage manages a list of tasks
type TaskManager struct {
	Tasks  []Task
	NextID int
}

// AddTask adds a new task to the manager
func (tm *TaskManager) AddTask(title string) {
	task := Task{
		ID:        tm.NextID,
		Title:     title,
		Completed: false,
	}
	tm.Tasks = append(tm.Tasks, task)
	tm.NextID++
	fmt.Printf("Task added: %q\n", title)
}

// ListTasks lists all tasks
func (tm *TaskManager) ListTasks() {
	if len(tm.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range tm.Tasks {
		status := "Incomplete"
		if task.Completed {
			status = "Complete"
		}
		fmt.Printf("[%d] %s - %s\n", task.ID, task.Title, status)
	}
}

// CompleteTask marks a task as completed
func (tm *TaskManager) CompleteTask(id int) {
	for i, task := range tm.Tasks {
		if task.ID == id {
			tm.Tasks[i].Completed = true
			fmt.Printf("Task %q marked as complete.=n", task.Title)
			return
		}
	}
	fmt.Printf("Task with ID %d not found.=n", id)
}

// ShowMenu displays the main menu
func ShowMenu() {
	fmt.Println("\nTask Manager")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Complete Task")
	fmt.Println("4. Quit")
	fmt.Print("Choose An option: ")
}

func main() {
	taskManager := TaskManager{}
	reader := bufio.NewReader(os.Stdin)

	for {
		ShowMenu()
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter task title: ")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			taskManager.AddTask(title)
		case "2":
			taskManager.ListTasks()
		case "3":
			fmt.Print("Enter task ID to mark as complete: ")
			var id int
			_, err := fmt.Scanf("%d\n", &id)
			if err != nil {
				fmt.Println("Invalid input. Please enter a numeric ID.")
				continue
			}
			taskManager.CompleteTask(id)
		case "4":
			fmt.Println("Exiting Task Manager. Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please choose a valid option.")
		}
	}
}
