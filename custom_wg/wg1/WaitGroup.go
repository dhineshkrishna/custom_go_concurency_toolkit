package main

import (
	"fmt"
	"sync"
)

type WaitGrouper struct {
	counter int
	mu      sync.Mutex
	cond    *sync.Cond
}

func NewWait() *WaitGrouper {
	wgL := &WaitGrouper{}
	wgL.cond = sync.NewCond(&wgL.mu)
	return wgL
}
func (wgL *WaitGrouper) add(i int) {
	wgL.mu.Lock()
	defer wgL.mu.Unlock()
	wgL.counter += i
}
func (wgL *WaitGrouper) done() {
	wgL.mu.Lock()
	defer wgL.mu.Unlock()
	wgL.counter--
	if wgL.counter == 0 {
		wgL.cond.Broadcast()
	}
}
func (wgL *WaitGrouper) wait() {
	wgL.mu.Lock()
	defer wgL.mu.Unlock()
	for wgL.counter > 0 {
		wgL.cond.Wait()
	}
}
func main() {
	wg := NewWait()
	wg.add(1)
	go func() {
		defer wg.done()
		fmt.Println("Data firtsdt")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()
	wg.add(1)
	go func() {
		defer wg.done()
		fmt.Println("Data second")
		for i := 0; i < 5; i++ {
			fmt.Println(i)
		}
	}()
	wg.wait()
}
