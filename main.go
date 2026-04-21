package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

)

type Task struct {
	Title string
	Done  bool
}

var tasks []Task
const fileName = "tasks.json"

// Load tasks from file
func loadTasks() {
	file, err := ioutil.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &tasks)
	}
}

// Save tasks to file
func saveTasks() {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	ioutil.WriteFile(fileName, data, 0644)
}

// Add task
func addTask(title string) {
	tasks = append(tasks, Task{Title: title, Done: false})
	saveTasks()
}

// List tasks
func listTasks() {
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		status := "❌"
		if task.Done {
			status = "✅"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Title, status)
	}
}

// Mark task as done
func markDone(index int) {
	if index >= 0 && index < len(tasks) {
		tasks[index].Done = true
		saveTasks()
	}
}

// Delete task
func deleteTask(index int) {
	if index >= 0 && index < len(tasks) {
		tasks = append(tasks[:index], tasks[index+1:]...)
		saveTasks()
	}
}

func main() {
	loadTasks()

	var choice int

	for {
		fmt.Println("\n==== TO-DO MENU ====")
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Mark Done")
		fmt.Println("4. Delete Task")
		fmt.Println("5. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var title string
			fmt.Print("Enter task: ")
			fmt.Scanln(&title)
			addTask(title)

		case 2:
			listTasks()

		case 3:
			var index int
			fmt.Print("Enter task number: ")
			fmt.Scan(&index)
			markDone(index - 1)

		case 4:
			var index int
			fmt.Print("Enter task number: ")
			fmt.Scan(&index)
			deleteTask(index - 1)

		case 5:
			fmt.Println("Goodbye 👋")
			return

		default:
			fmt.Println("Invalid choice")
		}
	}
}