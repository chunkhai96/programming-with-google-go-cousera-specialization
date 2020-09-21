package main

import (
	"fmt"
	"sync"
	"reflect"
)

type Philosopher struct {
	leftChopstick *sync.Mutex
	rightChopstick *sync.Mutex
	eatNum int
}

func (p *Philosopher) Eat(idx int) {
	p.leftChopstick.Lock()
	if reflect.ValueOf(p.rightChopstick).Elem().FieldByName("state").Int() > 0 {
		p.leftChopstick.Unlock()
	} else if p.eatNum >= 3{
		p.leftChopstick.Unlock()
	} else {
		p.rightChopstick.Lock()
		fmt.Printf("starting to eat %d\n", idx)
		fmt.Printf("finish eating %d\n", idx)
		p.eatNum++
		p.leftChopstick.Unlock()
		p.rightChopstick.Unlock()
	}
}

func main() {
	chopsticks := make([]sync.Mutex, 5)
	philos := make([]*Philosopher, 5)
	for i:=0; i<5; i++ {
		philos[i] = &Philosopher{&chopsticks[i], &chopsticks[(i+1)%5], 0}
	}

	for {
		allDone := true
		for i:=0; i<5; i++ {
			if philos[i].eatNum < 3 {
				allDone = false
				go philos[i].Eat(i+1)
			} else {
				fmt.Printf("Philosopher %d eat times: %d\n", i+1, philos[i].eatNum)
			}
		}
		if allDone {
			break
		}
	}
}