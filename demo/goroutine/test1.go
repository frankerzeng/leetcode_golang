package main

import (
	"fmt"
	"time"
)

func main() {
	var ch = make(chan int, 3)
	ch <- 1
	ch <- 12
	ch <- 13
	close(ch)

	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)

	fmt.Println(1)
	go test()
	time.Sleep(time.Second * 8)
}
func test() {
	time.Sleep(2 * time.Second)
	fmt.Println(2)

}
