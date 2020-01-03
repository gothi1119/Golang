//고루틴 동기화
//뮤텍스 사용

package main

import (
	"fmt"
	"runtime"
	"sync"
)

type count struct {
	num    int
	mutext sync.Mutex
}

func (c *count) increment() {
	c.mutext.Lock()
	c.num += 1
	c.mutext.Unlock()
}

func (c *count) result() {
	fmt.Println(c.num)
}

func main() {
	//고루틴 동기화 예제
	//실행 흐름 제어 및 변수 동기화 가능
	//공유 데이터 보호가 가장 중요
	//뮤텍스 사용(Muxtex)
	//sync.Mutex 선언 후 Lock, 종료 후 Unlock 사용

	//동기화를 사용하지 않은 경우 예제
	//시스템 전체 cpu 사용
	runtime.GOMAXPROCS(runtime.NumCPU())

	c := count{num: 0}
	done := make(chan bool)

	for i := 1; i <= 10000; i++ {
		go func() {
			c.increment()
			done <- true
			runtime.Gosched() // 다른 CPU에게 양보
		}()
	}
	for i := 1; i <= 10000; i++ {
		<-done
	}
	c.result()
}
