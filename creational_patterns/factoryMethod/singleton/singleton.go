package main

import (
	"fmt"
	"sync"
)

var lock = &sync.Mutex{}

type single struct {
}

var singleInstance *single

func getInstance() *single {

	if singleInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if singleInstance == nil {
			fmt.Println("Creating single instnace now.")
			singleInstance = &single{}
		} else {
			fmt.Println("Using Single instnace already created.")
		}
	}
	return singleInstance
}

func main() {

	for i := 0; i < 10; i++ {
		go getInstance()
	}

	fmt.Scanln()
}
