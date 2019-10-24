package tool

import (
	"fmt"
	"testing"
)

func Test_Task1(t *testing.T) {
	tasks := NewTask()
	tasks.AddTask(func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task1")
		return msg
	}).AddTask(func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task2")
		return msg
	}).AddTask(func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task3")
		return msg
	})

	result := tasks.Run(" run tasks")
	fmt.Println(result)
}

func Test_Task2(t *testing.T) {
	tasks := NewTask()

	task1 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task1")
		return msg
	}
	task2 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task2")
		return msg
	}
	task3 := func(data interface{}) interface{} {
		msg := fmt.Sprintf("%s\n %s", data, "i am task3")
		return msg
	}

	tasks.AddTasks(task1, task2, task3)
	result := tasks.Run(" run tasks")
	fmt.Println(result)
}
