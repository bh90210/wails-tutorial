package main

import "fmt"
import "time"
import "sync"

func routine() {
	for {
		select {
		case <-pause:
			fmt.Println("pause")
			select {
			case <-play:
				fmt.Println("play")
			case <-quit:
				wg.Done()
				return
			}
		case <-quit:
			wg.Done()
			return
		default:
			work()
		}
	}
}

func main() {
	wg.Add(1)
	go routine()

	time.Sleep(1 * time.Second)
	pause <- struct{}{}

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	pause <- struct{}{}

	time.Sleep(1 * time.Second)
	play <- struct{}{}

	time.Sleep(1 * time.Second)
	close(quit)

	wg.Wait()
	fmt.Println("done")
}

func work() {
	time.Sleep(250 * time.Millisecond)
	i++
	fmt.Println(i)
}

var play = make(chan struct{})
var pause = make(chan struct{})
var quit = make(chan struct{})
var wg sync.WaitGroup
var i = 0
