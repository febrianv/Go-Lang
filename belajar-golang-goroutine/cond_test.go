package belajargolanggoroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int){
	defer group.Done()
	group.Add(1)

	// ini tetap mengizinkan sati-satu yang jalan
	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		go WaitCondition(i)
	}	

	// signal membolehkan 1 goroutine untuk jalan setalah di wait
	// go func()  {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Signal()
	// 	}
	// }()

	// broacast setelah dieksekusi membolehkan seluruh goroutine untuk jalan setelah wait
	go func()  {
		time.Sleep(1 * time.Second)
		cond.Broadcast()
	}()

	group.Wait()
}