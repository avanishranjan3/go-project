package main

import (
	"root/Controllers"
	"root/Tasks"

	"github.com/hibiken/asynq"
)

type TestStruct struct {
	Name string `json:"name"`
}

func main() {
	// Controllers.GormDb()
	// Controllers.FetchDb()
	Controllers.ChannelsInGo()
	Tasks.SetKraData()

	//Redis connection
	redisCli := asynq.RedisClientOpt{
		DB:       0,
		Addr:     "127.0.0.1:6379",
		Password: "",
	}
	//new server
	worker := asynq.NewServer(redisCli, asynq.Config{
		Concurrency: 12,
		Queues: map[string]int{
			"my-task": 12,
		},
	})
	//register task handler function with worker
	mux := asynq.NewServeMux()
	mux.HandleFunc(Tasks.TypeAbc, Tasks.ABC)
	worker.Run(mux)

}
