package main

import (
	"fmt"
	"sync"
	"time"
)

type Host struct {
	C chan int
}

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	id                            int
	leftChopstick, rightChopstick *ChopStick
}

func (h *Host) start() {
	h.C <- 1
}

func (h *Host) finish() {
	<-h.C
}

func (philo Philosopher) eat(wg *sync.WaitGroup, host *Host) {

	defer wg.Done()

	for i := 0; i < 3; i++ {

		// the host here will block the process after send/receive & send/receive, after two processes :)

		host.start()

		philo.leftChopstick.Lock()
		philo.rightChopstick.Lock()

		fmt.Printf("==> [Starting] -- Philosopher (%d) is starting to eat. \n", philo.id)
		time.Sleep(time.Millisecond * 1000)
		fmt.Printf("==> [Finishing] -- Philosopher (%d) funished eating. \n", philo.id)

		philo.leftChopstick.Unlock()
		philo.rightChopstick.Unlock()

		host.finish()
	}
}

func main() {
	var wg sync.WaitGroup

	// initializations:
	tableAmount := 5
	chops := make([]*ChopStick, tableAmount)
	philos := make([]*Philosopher, tableAmount)
	host := &Host{
		C: make(chan int, 2),
	}

	for i := 0; i < tableAmount; i++ {
		chops[i] = new(ChopStick)
	}

	for i := 0; i < tableAmount; i++ {
		philos[i] = &Philosopher{i + 1, chops[i], chops[(i+1)%tableAmount]}
	}

	// create routines:
	for _, ph := range philos {

		wg.Add(1)
		go ph.eat(&wg, host)
	}

	wg.Wait()
}
