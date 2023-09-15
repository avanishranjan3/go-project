package Tasks

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

const TypeAbc = "tasks:abc"

func ABC(ctx context.Context, t *asynq.Task) error {
	fmt.Println("@ABC")
	var payload map[string]interface{}
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		fmt.Println("@ABC Err!", payload)
	}
	fmt.Println("@ABC info!", payload)
	return nil
}

func CreateTasks() *asynq.Task {
	payload, err := json.Marshal(map[string]interface{}{"clientId": "111"})
	if err != nil {
		fmt.Println(" @KraStatusJob Payload: ", payload)
		return nil
	}
	return asynq.NewTask(TypeAbc, payload)
}

func SetKraData() {
	t := CreateTasks()
	var redisClient = asynq.NewClient(asynq.RedisClientOpt{DB: 0, Addr: "127.0.0.1:6379", Password: ""})
	_, err := redisClient.Enqueue(t, asynq.Queue("my-task"), asynq.ProcessIn(1))
	if err != nil {
		fmt.Println("@SetKraData failed to set in redis: ", err)
		return
	}
	fmt.Println("@SetKraData job set successfully for ")
}
