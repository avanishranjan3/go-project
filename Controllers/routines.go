package Controllers

import (
	"fmt"
	"net/http"
	"sync"
)

func ChannelsInGo() {
	var wg sync.WaitGroup
	channel := make(chan int)
	wg.Add(2)
	go output(channel, &wg)
	go Input(channel, &wg)

	wg.Wait()
	close(channel)

}
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Dockerized Web App!")
}
func output(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i <= 5; i++ {
		ch <- i
	}
}
func Input(ch <-chan int, wg *sync.WaitGroup) {
	for num := range ch {
		fmt.Println("Number :- ", num)
	}
}
