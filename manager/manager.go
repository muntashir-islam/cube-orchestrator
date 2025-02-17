package manager

import (
	"fmt"
	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/muntashir-islam/cube-orchestrator/task"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[string][]*task.Task
	EventDb       map[string][]*task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("Selecting appropriate worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("Updating Task")
}

func (m *Manager) SendWork() {
	fmt.Println("Sending task")
}
