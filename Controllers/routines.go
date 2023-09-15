package Controllers

import (
	"fmt"
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
