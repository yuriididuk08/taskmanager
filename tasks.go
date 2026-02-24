package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var taskFile = "tasks.json"

func LoadTasks() ([]Task, error) {
	if _, err := os.Stat(taskFile); os.IsNotExist(err) {
		return []Task{}, nil
	}

	data, err := ioutil.ReadFile(taskFile)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func SaveTasks(tasks []Task) error {
	data, err := json.MarshalIndent(tasks, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(taskFile, data, 0644)
}

// Додавання завдання
func AddTask(title string) error {
	tasks, _ := LoadTasks()
	id := 1
	if len(tasks) > 0 {
		id = tasks[len(tasks)-1].ID + 1
	}
	tasks = append(tasks, Task{ID: id, Title: title})
	return SaveTasks(tasks)
}

func DeleteTask(id int) error {
	tasks, _ := LoadTasks()
	newTasks := []Task{}
	found := false
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		} else {
			found = true
		}
	}
	if !found {
		return errors.New("task not found")
	}
	return SaveTasks(newTasks)
}

func CompleteTask(id int) error {
	tasks, _ := LoadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Completed = true
			return SaveTasks(tasks)
		}
	}
	return errors.New("task not found")
}
