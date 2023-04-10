package task

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Completed bool   `json:"completed"`
}

func ListTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks to list! Why not take a vacation?")
		return
	}
	for _, task := range tasks {
		status := ""
		if task.Completed {
			status = "✅"
		}
		fmt.Printf("[%s] %d %s\n", status, task.ID, task.Name)
	}
}

func AddTask(tasks []Task, name string) []Task {
	newTask := Task{
		ID:        GetNextID(tasks),
		Name:      name,
		Completed: false,
	}
	return append(tasks, newTask)
}

func DeleteTask(tasks []Task, id int) []Task {
	var newTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func CompleteTask(tasks []Task, id int) []Task {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			break
		}
	}
	return tasks
}

func SaveTasks(file *os.File, tasks []Task) {
	bytes, err := json.Marshal(tasks)
	if err != nil {
		panic(err)
	}
	_, err = file.Seek(0, 0)
	if err != nil {
		panic(err)
	}
	err = file.Truncate(0)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(file)
	_, err = writer.Write(bytes)
	if err != nil {
		panic(err)
	}
	err = writer.Flush()
	if err != nil {
		panic(err)
	}
}

func GetNextID(tasks []Task) int {
	var nextID int
	for _, task := range tasks {
		if task.ID > nextID {
			nextID = task.ID
		}
	}
	return nextID + 1
}
