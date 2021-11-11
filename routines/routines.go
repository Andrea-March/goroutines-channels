package routines

import (
	"fmt"
	"github.com/golang/glog"
	"sync"
	"time"
)

type EventChannel struct {
	total       chan int
	endCh       chan int
	message 	string
	muxGzBuffer sync.RWMutex
}

func (c *EventChannel) start() {
	for{
		select {
		case <- c.endCh:
			glog.Info(fmt.Sprintf("Goroutine %s finished", c.message))
			//c.flush()
		case total := <-c.total:
			glog.Info(fmt.Sprintf("Iteration finished for routine %s, %v iteration(s) missing to reach total",c.message, 3-total))
		}
	}
}

func (c *EventChannel) reset(){
	c.total <- 0
}

func (c *EventChannel) flush() {
	c.muxGzBuffer.Lock()
	defer c.muxGzBuffer.Unlock()
	defer c.reset()
}

func ExecuteGoRoutine(message string) {
	glog.Info(fmt.Sprintf("Executing goroutine: %s", message))
	executeTask(message)
}

func executeTask(message string) {
	c := EventChannel{
		total: make(chan int),
		endCh: make(chan int),
		message: message,
	}
	defer c.start()
	total := 0
	go func() {
		for i:= 0; i<3; i++{
			time.Sleep(5*time.Second)
			total++
			c.total <- total
		}
		c.endCh <- 0
	}()
}