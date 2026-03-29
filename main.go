package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
)

type task struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
	Done bool   `json:"done"`
}

func main() {
	fmt.Println("Hello!")

	const filename = "tasks.json"
	scanner := bufio.NewScanner(os.Stdin)
	// scanner := bufio.NewScanner(os.Stdin)

	// taskList := loadTasks(filename)

	// fmt.Println(taskList)
	// taskList = addTask(taskList, "Buy wallpapers")

	// saveTasks(taskList, filename)

	// listTasks(taskList)

	fmt.Println("--- To do menu ---")
	fmt.Println("1 - List all tasks")
	fmt.Println("2 - Add task")
	fmt.Println("3 - Delete task")
	fmt.Println("4 - Mark done")
	fmt.Println("5 - Unmark")
	fmt.Println("0, ENTER - Exit")
	fmt.Println("help or 9 - to do menu")
	fmt.Println(" ")
	for {

		fmt.Print("Type option: ")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Exiting...")
			return
		}
		choice := scanner.Text()
		if choice == "" {
			fmt.Println("Exiting.")
			return
		}
		tasks := loadTasks(filename)
		switch choice {
		case "1":
			listTasks(tasks)
		case "2":
			fmt.Print("Type your new task: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Empty input, reload..")
				continue
			}
			newTask := scanner.Text()
			tasks = addTask(tasks, newTask)
			saveTasks(tasks, filename)
		case "3":
			fmt.Print("Type id of the task: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Empty input, reload..")
				continue
			}
			deleteID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Failed to convert input to integer!")
				continue
			}
			tasks = deleteTask(tasks, deleteID)
			saveTasks(tasks, filename)
		case "4":
			fmt.Print("Type id to mark done: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Empty input, reload")
				continue
			}
			markID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Failed to convert input to integer!")
				continue
			}
			tasks = markDone(tasks, markID)
			saveTasks(tasks, filename)
		case "5":
			fmt.Print("Type id to mark done: ")
			if ok := scanner.Scan(); !ok {
				fmt.Println("Empty input, reload")
				continue
			}
			unMarkID, err := strconv.Atoi(scanner.Text())
			if err != nil {
				fmt.Println("Failed to convert input into integer!")
				continue
			}
			tasks = unMarkDone(tasks, unMarkID)
			saveTasks(tasks, filename)
		case "0":
			saveTasks(tasks, filename)
			fmt.Println("Saving file and exiting..")
			return
		case "help", "9":
			fmt.Println("--- To do menu ---")
			fmt.Println("1 - List all tasks")
			fmt.Println("2 - Add task")
			fmt.Println("3 - Delete task")
			fmt.Println("4 - Mark done")
			fmt.Println("5 - Unmark")
			fmt.Println("0 - Exit")
			fmt.Println("help or 9 - to do menu")
		default:
			fmt.Println("Not a choice.")
			continue
		}
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
func listTasks(tasks []task) {
	for i := 0; i < len(tasks); i++ {
		mark := "[ ]"
		if tasks[i].Done == true {
			mark = "[✔]"
		}
		fmt.Printf("%v. %s %s\n", tasks[i].ID, mark, tasks[i].Task)
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

// Delete + reindex tasks
func deleteTask(tasks []task, id int) []task {
	for i, t := range tasks {
		if t.ID == id {
			tasks := append(tasks[:i], tasks[:i+1]...)
			for j := range tasks {
				tasks[j].ID = j + 1
			}
			return tasks
		}
	}
	fmt.Printf("Not finded task %v.\n", id)
	return tasks
}

// Mark done and no
func markDone(tasks []task, id int) []task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return tasks
		}
	}
	fmt.Printf("%v not founded.", id)
	return tasks
}
func unMarkDone(tasks []task, id int) []task {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = false
			return tasks
		}
	}
	fmt.Printf("%v not founded.", id)
	return tasks
}

// k
