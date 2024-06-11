package main

import (
	"bufio"
	"fmt"
	"os"
)

func todoActions() {
	fmt.Println("Choose an option below")
	fmt.Println("type (a to Add Todo, s to show All Todos, u to update a task, d to delete a task, e to Exit)")
	fmt.Print("(m to mark a task complete): ")
	reader := bufio.NewReader(os.Stdin)
	option, _ := userInput(reader)

	switch option {
	case "a":
		fmt.Print("Type your todo: ")
		taskDesc, _ := userInput(reader)
		newT := newTask(taskDesc)
		addTask(&newT)
		fmt.Println("Todo added successfully!")
		todoActions()
	case "s":
		displayTask()
		todoActions()
	case "u":
		updateTask()
		todoActions()
	case "d":
		deleteTodo()
		todoActions()
	case "m":
		markAsCompleted()
		todoActions()
	case "e":
		fmt.Println("Exited successfully! Have a nice day")
	default:
		fmt.Println("Not a valid option")
		todoActions()
	}

}

func main() {
	fmt.Println("Welcome to your Todo List App")
	todoActions()
}
