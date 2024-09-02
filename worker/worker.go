package worker

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/muntashir-islam/cube-orchestrator/task"
)

type Worker struct {
	Name      string
	Db        map[uuid.UUID]*task.Task
	Queue     queue.Queue
	TaskCount int
}

func (w *Worker) CollectStates() {
	fmt.Println("Collecting State")
}

func (w *Worker) RunTask() {
	fmt.Println("Start or stop a task")
}

func (w *Worker) StartTask() {
	fmt.Println("start a task")
}

func (w *Worker) StopTask() {
	fmt.Println("Stop Task")
}
