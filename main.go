package main

import "fmt"

type TODOLISt struct {
	task     string
	priority string
}

var TodoList = []TODOLISt{}

func (t TODOLISt) CreateTask(task string, priority string) {

}

func main() {
	fmt.Println("--------------------------X---------------------------")
	fmt.Println("Welcome To the TODO LIST CLI APP!")
out:
	for {
		fmt.Println("Enter 1.Show Todo List 2.Add Task to List 3.Delete From Todo List 4.Update Task in Todo List 5.Exit Application")
		var choice int
		fmt.Scanln(&choice)
		switch choice {

		case 1:
			for index, el := range TodoList {
				fmt.Printf("%d:%+v\n", index+1, el)
			}

		case 2:
			fmt.Println("Enter the task name")
			var taskName string
			fmt.Scanln(&taskName)
			fmt.Println("Enter the priority")
			var priority string
			fmt.Scanln(&priority)
			newTask := TODOLISt{task: taskName, priority: priority}
			TodoList = append(TodoList, newTask)

		case 3:
			fmt.Println("Enter the Index of the task to be deleted")
			var index int
			fmt.Scanln(&index)
			TodoList = append(TodoList[:index], TodoList[index+1:]...)
			for _, el := range TodoList {
				fmt.Println(el)
			}
		case 4:
			fmt.Println("Enter the element index to be updated")
			var index int
			fmt.Scanln(&index)
			fmt.Println("Enter the new Task Name")
			var newTaskName string
			fmt.Scanln(&newTaskName)
			fmt.Println("Enter the new Priority")
			var newPriority string
			fmt.Scanln(&newPriority)
			TodoList[index] = TODOLISt{task: newTaskName, priority: newPriority}
			for _, el := range TodoList {
				fmt.Println(el)
			}
		case 5:
			break out
		}
	}

}
