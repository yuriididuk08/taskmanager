package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: task [add|list|delete|done] ...")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task title")
			return
		}
		title := os.Args[2]
		if err := AddTask(title); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task added:", title)
		}
	case "list":
		tasks, _ := LoadTasks()
		for _, t := range tasks {
			status := "[ ]"
			if t.Completed {
				status = "[x]"
			}
			fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
		}
	case "delete":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID to delete")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := DeleteTask(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task deleted:", id)
		}
	case "done":
		if len(os.Args) < 3 {
			fmt.Println("Please provide task ID to mark as done")
			return
		}
		id, _ := strconv.Atoi(os.Args[2])
		if err := CompleteTask(id); err != nil {
			fmt.Println("Error:", err)
		} else {
			fmt.Println("Task completed:", id)
		}
	default:
		fmt.Println("Unknown command:", os.Args[1])
	}
}
