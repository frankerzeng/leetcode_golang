package main

import (
	"fmt"
	"time"
)

func main() {

	var ch1 = make(chan int, 10)

	go func(ch chan int) {
		ch <- 10
		time.Sleep(time.Second * 2)
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}(ch1)

	for v := range ch1 {
		fmt.Println(v)
	}

}
