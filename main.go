package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)


func main() {

	todolist:=TodoList{}
	

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Println("============== TODO LIST ===============")
		fmt.Println("1. Add Task")
		fmt.Println("2. View Taks")
		fmt.Println("3. Mark Task as Done")
		fmt.Println("1. Exit")
		fmt.Println("===========================================")
		fmt.Println("Enter Your choice: ")

		scanner.Scan()

		choice := scanner.Text()

		fmt.Println("Your choice :", choice)

		switch choice {
		case "1":
			fmt.Print("Enter task name :")
			scanner.Scan()
			taskname:= scanner.Text()
			todolist.AddTask(taskname)
		case "2":
			todolist.ViewTask()
		case "3":
			fmt.Print("Enter  Task ID")
			scanner.Scan()
			id:=scanner.Text()
			taskid,err :=strconv.Atoi(id)
			if err !=nil {
				fmt.Println("invalide task id")
				continue
				
			}

			todolist.MarkAsDone(taskid)
		case "4":
			fmt.Println("Exiting .....")
			return
		default:
			fmt.Println("Invalid choice Try agin")
		}
	}

}
