package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	go func() {
		for i := 0; i < 10; i++ {
			if i < 6 {

				ch <- i
				fmt.Println("sent:", i)
			}
			if i == 5 {

				close(ch)
			}
		}
	}()
	for i := 0; i < 10; i++ {
		fmt.Scanln()
		j, ok := <-ch
		fmt.Println("received:", j, ok)
	}
}
