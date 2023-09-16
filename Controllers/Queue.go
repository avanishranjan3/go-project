package Controllers

import (
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const TaskType = "task:user"

func CreateJob() *asynq.Task {
	var payload []map[string]interface{}
	payload = append(payload, map[string]interface{}{"Name: ": "Avanish", "Age: ": "24", "Proff: ": "Eng"})
	payload = append(payload, map[string]interface{}{"Name: ": "Aryan", "Age: ": "26", "Proff: ": "Mng"})
	task, err := json.Marshal(payload)
	if err != nil {
		fmt.Println("Error ! ", err)
		return nil
	}
	return asynq.NewTask(TaskType, task)
}

func SendJobInQueue() {
	task := CreateJob()
	rediscli := asynq.NewClient(asynq.RedisClientOpt{DB: 0, Addr: "127.0.0.1:6379", Password: ""})
	_, err := rediscli.Enqueue(task, asynq.Queue("user-detail"), asynq.ProcessIn(1))

	if err != nil {
		fmt.Println("Error in Queuing job!", err)
		return
	}
	fmt.Println("successfully queued!")
}
