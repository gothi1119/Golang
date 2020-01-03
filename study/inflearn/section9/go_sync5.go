//고루틴 동기화 기초(5)

package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	//고루틴 동기화 상태
	//동기화 상태(조건) 메소드 사용
	//Wait,notify, notifyAll : 기타 언어
	//Wait, Signal, Broadcast

	// 시스템 전체 CPU 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	var mutex = new(sync.Mutex)
	var condition = sync.NewCond(mutex)

	c := make(chan int, 5) //비동기

	for i := 0; i < 5; i++ {
		go func(n int) {
			mutex.Lock()
			c <- 777
			fmt.Println("Goroutine. Waiting:", n)
			condition.Wait()
			fmt.Println("Waiting End", n)
			mutex.Unlock()
		}(i)
	}
	for i := 0; i < 5; i++ {
		<-c
		//fmt.Println("receive:", <-c)
	}
	for i := 0; i < 5; i++ {
		mutex.Lock()
		fmt.Println("Wake Goroutine(Signal) : ", i)
		condition.Signal() //한개씩 깨움
		mutex.Unlock()
	}
	time.Sleep(2 * time.Second)
}
