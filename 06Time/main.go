package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(2000, 12, 12, 13, 0, 0, 0, time.UTC)
	fmt.Println(t.Format("2006.January.02 03(15):04:05 MST Monday"))
	tStr, err := time.Parse("15:04:05", "23:30:00")
	fmt.Printf("%v %T %v \n", tStr, tStr, err)
	fmt.Println(time.Unix(1676626000, 0))
	fmt.Println(t.In(time.Local))
	fmt.Println(t.Add(-10 * time.Minute))
	fmt.Println(t.AddDate(1, 1, 20))
	fmt.Println(t)
	fmt.Println(t.Sub(time.Now()))

	for range time.Tick(time.Second) {
		fmt.Println(time.Now())
		//time.Sleep(time.Second)
	}
}
