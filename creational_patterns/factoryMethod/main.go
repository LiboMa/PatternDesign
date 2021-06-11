package main

import (
	"fmt"
	"github.com/LiboMa/patterns-demo/components"
	//ob "github.com/LiboMa/patterns-demo/observer"
	. "github.com/LiboMa/patterns-demo/observer"
	//_ "github.com/LiboMa/patterns-demo/observer"
	"time"
)

func testObserver() {

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

func main() {
	fmt.Println("vim-go")
	s := components.New()
	s["this"] = "that"
	s2 := components.New()
	fmt.Println("This is", s2["this"])

	testObserver()

}
