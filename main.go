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
	const filename = "tasks.json"
	tasks := loadTasks(filename)
	for {
		fmt.Println("\n--- TODO MENU ---")
		fmt.Println("[1] Add new task")
		fmt.Println("[2] List all tasks")
		fmt.Println("[3] Mark task as Done")
		fmt.Println("[4] Unmark task")
		fmt.Println("[5] Delete task")
		fmt.Println("[6] Edit task text")
		fmt.Println("[0] Exit")
		fmt.Print("Choose an option: ")

		var choice int
		fmt.Scanln(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter task name: ")
			var name string
			// Use fmt.Scanln to catch the name (note: space-separated words need more care)
			fmt.Scanln(&name)
			tasks = addTask(name, tasks)
			saveTasks(filename, tasks)
		case 2:
			readAll(tasks)
		case 3:
			fmt.Print("Enter ID to mark done: ")
			var id int
			fmt.Scanln(&id)
			markDone(tasks, id)
			saveTasks(filename, tasks)
		case 4:
			fmt.Print("Enter ID to unmark: ")
			var id int
			fmt.Scanln(&id)
			unmarkDone(tasks, id)
			saveTasks(filename, tasks)
		case 5:
			fmt.Print("Enter ID to delete: ")
			var id int
			fmt.Scanln(&id)
			tasks = deleteTask(tasks, id)
			tasks = reindexTasks(tasks)
			saveTasks(filename, tasks)
		case 6:
			fmt.Print("Enter ID to edit: ")
			var id int
			fmt.Scanln(&id)
			fmt.Print("Enter new text: ")
			var newText string
			fmt.Scanln(&newText)
			changeTask(tasks, id, newText)
			saveTasks(filename, tasks)
		case 0:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice, try again.")
		}
	}

	// Second verison
	// const filename = "tasks.json"
	// tasks := loadTasks(filename)

	// tasks = addTask("Buy milk", tasks)
	// tasks = addTask("Finish Go project", tasks)

	// markDone(tasks, 1)

	// saveTasks(filename, tasks)

	// readAll(tasks)

	// Firts verison
	// user := []task{}

	// // Переводим нашу структуру в json
	// jsonData, err := json.Marshal(user)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(string(jsonData))

	// // Создаем config.json
	// file, err := os.Create("tasks.json")
	// if err != nil {
	// 	return
	// }
	// defer file.Close()

	// // Проверяем ли не создан файл уже
	// // Создаем енкодер для созданного файла
	// encoder := json.NewEncoder(file)
	// // Делаем json beauty
	// encoder.SetIndent("", "  ")

	// // Input Json
	// err = encoder.Encode(user)
	// if err != nil {
	// 	panic(err)
	// }
}

// Changing already existing task text
func changeTask(tasks []task, id int, newText string) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Task = newText
			break
		}
	}
}

// Загружаем уже созданные задачи
func loadTasks(filename string) []task {
	file, err := os.ReadFile(filename)
	if err != nil {
		return []task{}
	}
	var tasks []task
	json.Unmarshal(file, &tasks)
	return tasks
}

// Add new task
func addTask(name string, tasks []task) []task {
	newId := len(tasks) + 1
	newTask := task{ID: newId, Task: name, Done: false}
	return append(tasks, newTask)
}

func saveTasks(filename string, tasks []task) {
	file, _ := os.Create(filename)
	defer file.Close()
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	encoder.Encode(tasks)
}

// Delete existing task
func deleteTask(tasks []task, id int) []task {
	for i, t := range tasks {
		if t.ID == id {
			return append(tasks[:i], tasks[i+1:]...)
		}
	}
	fmt.Printf("Task with ID %d not found.\n", id)
	return tasks
}
func reindexTasks(tasks []task) []task {
	for i := range tasks {
		tasks[i].ID = i + 1
	}
	return tasks
}

// Marking that task is done
func markDone(tasks []task, id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			break
		}
	}
}

// Marking that task is not done yet
func unmarkDone(tasks []task, id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = false
			break
		}
	}
}

// List all existing tasks
func readAll(tasks []task) {
	fmt.Println("\n--- MY TODO LIST ---")
	if len(tasks) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[✔]"
		}

		fmt.Printf("%d. %s %s\n", t.ID, status, t.Task)
	}
	fmt.Println("--------------------")
}
