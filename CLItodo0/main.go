package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . [add|list|done|delete]")
		return
	}

	command := os.Args[1]
	tasks, _ := LoadTasks()

	switch command {

	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: add <task>")
			return
		}

		task := Task{
			ID:    len(tasks) + 1,
			Title: os.Args[2],
			Done:  false,
		}

		tasks = append(tasks, task)
		SaveTasks(tasks)

		fmt.Println("Task added!")

	case "list":
		for _, t := range tasks {
			status := "❌"
			if t.Done {
				status = "✅"
			}
			fmt.Printf("%d. %s [%s]\n", t.ID, t.Title, status)
		}

	case "done":
		id, _ := strconv.Atoi(os.Args[2])

		for i := range tasks {
			if tasks[i].ID == id {
				tasks[i].Done = true
			}
		}

		SaveTasks(tasks)
		fmt.Println("Task marked as done!")

	case "delete":
		id, _ := strconv.Atoi(os.Args[2])

		var newTasks []Task
		for _, t := range tasks {
			if t.ID != id {
				newTasks = append(newTasks, t)
			}
		}

		SaveTasks(newTasks)
		fmt.Println("Task deleted!")

	default:
		fmt.Println("Unknown command")
	}
}
