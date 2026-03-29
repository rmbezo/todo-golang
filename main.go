package main

import (
	"encoding/json"
	"fmt"
	"log"
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

	taskList := loadTasks(filename)

	fmt.Println(taskList)
	taskList = addTask(taskList, "Buy wallpapers")

	saveTasks(taskList, filename)

	listTasks(taskList)
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
	newId := len(tasks) + 1
	newTask := task{ID: newId, Task: name, Done: false}
	return append(tasks, newTask)
}

// List
func listTasks(task []task) {
	for i := 0; i < len(task); i++ {
		mark := "[ ]"
		if task[i].Done == true {
			mark = "[✔]"
		}
		fmt.Printf("%v. %s %s\n", task[i].ID, mark, task[i].Task)
	}
}

// Save tasks
func saveTasks(tasks []task, filename string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)
}

// Delete

// Mark done and no
func markDone(tasks []task, id int) {

}

// k
