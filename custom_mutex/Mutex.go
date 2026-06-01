package main

import (
	"fmt"
	"sync"
)

type Mutex struct {
	ch chan struct{}
}

func NewMutex() *Mutex {
	mu := &Mutex{ch: make(chan struct{}, 1)}
	mu.ch <- struct{}{}
	return mu
}
func (mu *Mutex) Lock() {
	<-mu.ch
}

func (mu *Mutex) UnLock() {
	mu.ch <- struct{}{}
}

func main() {
	m := NewMutex()
	var wg sync.WaitGroup
	counter := 0
	for i := 1; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			m.Lock()
			defer m.UnLock()
			counter += i

		}(i)
	}

	wg.Wait()
	fmt.Println(counter)
}
