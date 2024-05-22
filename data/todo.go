package data

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"
)

type Task struct {
	Title       string
	Description string
	Duedate     time.Time
	Created_at  time.Time
	Updated_at  time.Time
}

type TaskList struct {
	list map[string]Task
}

func (tl *TaskList) Add(task Task) {
	tl.list[task.Title] = task
}

func clear(duration time.Duration) {
	time.Sleep(time.Second * duration)
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	err := cmd.Run()

	if err != nil {
		fmt.Println("error clearing fails")
	}
}

func Menu() {
	var opt int
	var taskList TaskList
	taskList.list = make(map[string]Task)

	for opt != 4 {
		fmt.Println(`
		|***** TRELLO *****|
		|[1] add task *****|
		|[2] remove task **|
		|[3] list task ****|
		|[4] exit     *****|
		`)

		fmt.Print("choose a opt: ")
		fmt.Scanln(&opt)

		switch opt {
		case 1:
			addTask(taskList)
			clear(1)
		case 2:
			dropTask(taskList)
			clear(1)
		case 3:
			showTasks(taskList)
			clear(5)
		}

	}
}

func addTask(taskList TaskList) {
	task := new(Task)
	var duedate string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Write title: ")
	title, err := reader.ReadString('\n')

	if err != nil {
		task.Title = ""
	}

	fmt.Print("write description: ")
	description, err := reader.ReadString('\n')

	if err != nil {
		task.Description = ""
	}

	fmt.Print("write due date: ")
	fmt.Scanln(&duedate)

	parsedDate, err := time.Parse("2006-01-02", duedate)

	task.Duedate = parsedDate

	if err != nil {
		task.Duedate = time.Now().AddDate(0, 0, 3)
	}

	task.Title = title
	task.Description = description
	task.Created_at = time.Now()
	task.Updated_at = time.Now()

	taskList.Add(*task)
}

func dropTask(taskList TaskList) {
	var title string

	fmt.Print("Write the title of the task to eliminate: ")
	fmt.Scanln(&title)

	delete(taskList.list, title)
}

func showTasks(taskList TaskList) {
	json, err := json.MarshalIndent(taskList.list, "", "  ")

	if err != nil {
		fmt.Println("{}")
	}

	fmt.Println(string(json))
}
