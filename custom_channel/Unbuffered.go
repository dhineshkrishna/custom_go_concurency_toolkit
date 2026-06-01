package main

import (
	"fmt"
	"sync"
)

type Unbuffered[T any] struct {
	value   T
	mu      sync.Mutex
	cond    *sync.Cond
	hasData bool
}

func UnbufferedChan[T any]() *Unbuffered[T] {
	ub := &Unbuffered[T]{}
	ub.cond = sync.NewCond(&ub.mu)
	return ub
}

func (ub *Unbuffered[T]) send(i T) {
	ub.mu.Lock()
	defer ub.mu.Unlock()

	// wait until previous value is consumed
	for ub.hasData {
		ub.cond.Wait()
	}
	ub.value = i
	ub.hasData = true

	// for wakeup
	ub.cond.Signal()

	// wait until receiver takes it
	for ub.hasData {
		ub.cond.Wait()
	}
}

func (ub *Unbuffered[T]) receive() T {
	ub.mu.Lock()
	defer ub.mu.Unlock()

	for !ub.hasData {
		ub.cond.Wait()
	}
	v := ub.value
	ub.hasData = false

	ub.cond.Signal()

	return v

}

func main() {
	Ubb := UnbufferedChan[int]()
	go func() {
		Ubb.send(10)
		Ubb.send(20)
	}()
	kk := Ubb.receive()
	fmt.Println(kk,Ubb.receive())

}
