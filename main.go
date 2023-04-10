package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	task "github.com/devchrisar/gocli_frugal_stardust/tasks"
	"github.com/fatih/color"
	"github.com/rodaine/table"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.OpenFile("tasks.json", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var tasks []task.Task
	info, err := file.Stat()
	if err != nil {
		panic(err)
	}
	if info.Size() != 0 {
		bytes, err := io.ReadAll(file)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(bytes, &tasks)
		if err != nil {
			panic(err)
		}
	} else {
		tasks = []task.Task{}
	}

	if len(os.Args) < 2 {
		PrintUsage()
		return
	}
	switch os.Args[1] {
	case "list":
		task.ListTasks(tasks)
	case "add":
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Enter task name:")
		name, _ := reader.ReadString('\n')
		name = strings.TrimSpace(name)
		tasks = task.AddTask(tasks, name)
		task.SaveTasks(file, tasks)
	case "delete":
		if len(os.Args) < 3 {
			PrintUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("the second argument must be an integer")
			return
		}
		tasks = task.DeleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	case "complete":
		if len(os.Args) < 3 {
			PrintUsage()
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("the second argument must be an integer")
			return
		}
		tasks = task.CompleteTask(tasks, id)
		task.SaveTasks(file, tasks)
	default:
		PrintUsage()
	}
}

func PrintUsage() {
	fmt.Println(" ⭐ Welcome to the Go CLI! ⭐ ")
	fmt.Println("Usage: goCli [command] [arguments]")
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()
	tbl := table.New("Commands:")
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	tbl.AddRow("\t📌 list\t\t\t\tList all tasks")
	tbl.AddRow("\t📌 add\t\t\t\tAdd a new task")
	tbl.AddRow("\t📌 complete [arguments]\t\tComplete a task")
	tbl.AddRow("\t📌 delete [arguments]\t\tDelete a task")
	tbl.Print()
}
