package routines

import (
	"fmt"
	"github.com/golang/glog"
	"sync"
)

type EventChannel struct {
	total       chan int
	endCh       chan int
	muxGzBuffer sync.RWMutex
}

func (ch *EventChannel) start() {
	for{
		glog.Info("HERE")
		select {
		case <- c.endCh:
			c.flush()
			glog.Info("Three goroutines finished, printing results...")
			glog.Info(fmt.Sprintf("Total esecuted: %-5v, messages: %s", total, messages))
		case <- c.total:
			glog.Info(fmt.Sprintf("Goroutine finished, %v missing to reach total",3-total))
			glog.Info(fmt.Sprintf("Total esecuted: %-5v, messages: %s", total, messages))
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

var total int
var messages []string
var c EventChannel

func init()  {
	go c.start()
}

func ExecuteGoRoutine(message string) {
	glog.Info(fmt.Sprintf("Executing goroutine, iteration %v, message %s", (total+1), message))
	executeTask(message)
}

func executeTask(message string) EventChannel{
	go func() {
		total++
		messages = append(messages, message)
		if total>= 3{
			c.endCh <-0
			return
		}
		c.total <- total
	}()
	return c
}