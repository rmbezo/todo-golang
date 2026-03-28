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

	user := []task{}

	// Переводим нашу структуру в json
	jsonData, err := json.Marshal(user)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))

	// Создаем config.json
	file, err := os.Create("tasks.json")
	if err != nil {
		return
	}
	defer file.Close()

	// Проверяем ли не создан файл уже
	// Создаем енкодер для созданного файла
	encoder := json.NewEncoder(file)
	// Делаем json beauty
	encoder.SetIndent("", "  ")

	// Input Json
	err = encoder.Encode(user)
	if err != nil {
		panic(err)
	}
}

// Загружаем уже созданные задачи
func loadTasks(filename string) task {

	return task{}
}

// Add new task
func addTask(name string, id int, done bool) {}

// Delete existing task
func deleteTask(id int) {}

// Marking that task is done
func markDone(id int) {}

// Marking that task is not done yet
func unmarkDone(id int) {}

// List all existing tasks
func readAll() {}

// Changing already existing task and saving ID
func changeTask(id int) {}
