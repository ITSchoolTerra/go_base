package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	mtx := sync.Mutex{}
	j := 0
	for i := 0; i < 1000; i++ {

		wg.Add(1)
		go func(wg *sync.WaitGroup, i int) {
			mtx.Lock()
			j++
			mtx.Unlock()
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println(j)
}

//func test(wg *sync.WaitGroup) {
//	fmt.Println("hello")
//	wg.Done()
//}
