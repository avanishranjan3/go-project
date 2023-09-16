package main

import (
	"root/Controllers"
)

type TestStruct struct {
	Name string `json:"name"`
}

func main() {
	// Controllers.GormDb()
	// Controllers.FetchDb()
	Controllers.ChannelsInGo()
	Controllers.SendJobInQueue()

}
