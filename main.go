package main

import (
	"fmt"

	"github.com/shzy2012/common/tool"
)

func main() {
	tasks := tool.NewTask()

	task1 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "exec=> i am task1")
		return msg
	}
	task2 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "exec=> i am task2")
		return msg
	}
	task3 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "exec=> i am task3")
		return msg
	}

	tasks.AddTasks(task1, task2, task3)
	result := tasks.Run(" run tasks")
	fmt.Println(result)
}
