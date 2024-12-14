package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"taskmanager2/internal/app"
	"taskmanager2/internal/persistence"
)

func main() {
	store := persistence.NewJSONStore("tasks.json")
	tasks, nextID, err := store.LoadTasks()
	if err != nil {
		fmt.Printf("Error loading tasks: %v\n", err)
		return
	}

	taskManager := &app.TaskManager{Tasks: tasks, NextID: nextID}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\nTask Manager")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Complete Task")
		fmt.Println("4. Quit")
		fmt.Print("Choose An option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":

		}
	}
}
