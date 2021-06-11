package observer

import (
	"fmt"
	"testing"
	"time"
)

func TestObserver() {

	// Initialize a new notifer
	// n := equals var n typeofN = value
	n := EventNotifier{
		Observers: map[Observer]struct{}{},
	}

	observerlist := []int{1, 2, 3, 4, 5, 6, 7, 8}

	for _, ob := range observerlist {
		n.Register(&EventObserver{Id: ob})
	}

	stop := time.NewTimer(20 * time.Second).C
	tick := time.NewTicker(time.Second).C

	for {
		select {
		case <-stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}

}
