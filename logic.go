package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []singleTask

type singleTask struct {
	description string
	completed   bool
}

func userInput(r *bufio.Reader) (string, error) {
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func newTask(d string) singleTask {
	t := singleTask{
		description: d,
		completed:   false,
	}

	return t
}

func convToInt(n string) int {
	i, err := strconv.Atoi(n)

	if err != nil {
		panic(err)
	}

	return i - 1
}

func getItemByIndex(i int) *singleTask {
	return &tasks[i]
}

func addTask(t *singleTask) {
	tasks = append(tasks, (*t))
}

func displayTask() {
	if len(tasks) == 0 {
		fmt.Println("No task created yet!")
		todoActions()
		return
	}
	for idx, task := range tasks {
		desc := task.description
		number := idx + 1
		fmt.Printf("%v. %v", number, desc)
		if task.completed == true {
			fmt.Print(" -Completed\n")
		} else {
			fmt.Print(" -Not completed\n")
		}

	}
}

func updateTask() {
	if len(tasks) == 0 {
		fmt.Println("No task to update, add task!")
		todoActions()
		return
	}

	fmt.Print("Type the task number you want to update: ")
	reader := bufio.NewReader(os.Stdin)
	number, _ := userInput(reader)

	conIdx := convToInt(number)

	if conIdx < 0 || conIdx >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	val := getItemByIndex(conIdx)

	fmt.Print("Type your updated task: ")
	updated, _ := userInput(reader)

	val.description = updated

	fmt.Println("Task updated successfully!")
	displayTask()
}

func remove(s []singleTask, i int) {
	tasks = append(s[:i], s[i+1:]...)
}

func deleteTodo() {
	if len(tasks) == 0 {
		fmt.Println("No task to delete, add task!")
		todoActions()
		return
	}

	fmt.Print("Type the task number you want to delete: ")
	reader := bufio.NewReader(os.Stdin)
	num, _ := userInput(reader)
	conIdx := convToInt(num)

	if conIdx < 0 || conIdx >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	remove(tasks, conIdx)
	fmt.Println("Task deleted successfully!")
	displayTask()
}

func markAsCompleted() {
	if len(tasks) == 0 {
		fmt.Println("No task has been created!, add task!")
		todoActions()
		return
	}

	fmt.Print("Type the task number you want to mark completed: ")
	reader := bufio.NewReader(os.Stdin)
	number, _ := userInput(reader)

	conIdx := convToInt(number)

	if conIdx < 0 || conIdx >= len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	val := getItemByIndex(conIdx)

	val.completed = true

	displayTask()

}
