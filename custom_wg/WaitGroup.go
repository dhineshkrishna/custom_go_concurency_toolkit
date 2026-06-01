package main

import (
	"fmt"
	"sync"
)

type WaitGroup struct {
	counter int
	mu      sync.Mutex
	done    chan struct{}
}

func NewWaitGroup() *WaitGroup {
	return &WaitGroup{
		done: make(chan struct{}),
	}
}

func (wg *WaitGroup) add(i int) {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter += i
}
func (wg *WaitGroup) wait() {
	<-wg.done
}

func (wg *WaitGroup) Done() {
	wg.mu.Lock()
	defer wg.mu.Unlock()
	wg.counter--
	if wg.counter == 0 {
		close(wg.done)
	}
}

func main() {
	wg := NewWaitGroup()
	wg.add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Data firtsdt")
	}()
	wg.add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Data second")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()
	wg.wait()
	//time.Sleep(5 * time.Second)

}
