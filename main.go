package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func main() {
	fmt.Println("Hello!")

	const filename = "tasks.json"
	// scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(loadTasks(filename))

	for {
		break
	}
}

// Load tasks
func loadTasks(filename string) []task {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []task{}
	}
	var tasks []task
	json.Unmarshal(file, &tasks)
	return tasks
}

// Add
func addTask(tasks []task, name string) []task {

	return []task{}
}

// List
func listTasks(task []task) {

}

// Delete

// Mark done and no

// k
