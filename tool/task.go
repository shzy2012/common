//author: joey.zhou
//date: 2019-10-24
//desc: 多个任务，形成任务链执行

/*
	asks := NewTask()

	//创建任务task1
	task1 := func(data interface{}) interface{} {
		return data
	}

	//创建任务task2
	task2 := func(data interface{}) interface{} {
		return data
	}

	//创建任务task3
	task3 := func(data interface{}) interface{} {
		return data
	}

	//创建任务链
	tasks.AddTasks(task1, task2, task3)

	//执行任务链
	tasks.Run(" run tasks")
*/

package tool

type taskFunc func(data interface{}) interface{}

//Task task链
type Task struct {
	tasks []taskFunc
}

//NewTask 创建task链
func NewTask() *Task {
	return &Task{
		tasks: make([]taskFunc, 0),
	}
}

//AddTask 添加task
func (t *Task) AddTask(task taskFunc) *Task {
	t.tasks = append(t.tasks, task)
	return t
}

//AddTasks 添加多个task
func (t *Task) AddTasks(tasks ...taskFunc) *Task {
	t.tasks = append(t.tasks, tasks...)
	return t
}

//Run 运行task
func (t *Task) Run(data interface{}) interface{} {
	for _, task := range t.tasks {
		data = task(data)
	}

	return data
}
