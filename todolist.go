 package main

import "fmt"
 

type TodoList struct {
 task []Task 
}

func (td *TodoList) AddTask(taskname string) {
	fmt.Println("Task added")
	taskid := len(td.task)+1
	task:= NewTask(taskid,taskname)

	td.task =append(td.task, task)
	fmt.Println("Task Added Successfulljy!")
	
}

func (td *TodoList) ViewTask() {
	fmt.Println("Task are ", td.task)

}
func (td *TodoList) MarkAsDone(taskId int) {
	


}
