package main

import (
	"fmt"
	"time"
)

func workerPool(n int, jobCh <-chan int, retCh chan<- string) {
	for i := 0; i < n; i++ {
		go worker(i, jobCh, retCh)
	}
}

func worker(id int, jobCh <-chan int, retCh chan<- string) {
	cnt := 0
	for job := range jobCh {
		cnt++
		ret := fmt.Sprintf("worker %d processed job: %d, it's the %dth processed by me.", id, job, cnt)
		retCh <- ret
	}
}
func main() {

	m := map[int]int{1: 1}
	for i := 0; i < 100; i++ {
		go func(m map[int]int, i int) {
			m[1] = i
		}(m, i)
	}
	b := make([]int, 1024)
	b = append(b, 99)
	fmt.Println("len:", len(b), "cap:", cap(b))
	return
	fmt.Println(m)
	jobCh := genJob(10000)
	retCh := make(chan string, 10000)
	workerPool(5, jobCh, retCh)

	time.Sleep(time.Second)
	close(retCh)
	for ret := range retCh {
		fmt.Println(ret)
	}
}

func genJob(n int) <-chan int {
	jobCh := make(chan int, 2000)
	go func() {
		for i := 0; i < n; i++ {
			jobCh <- i
		}
		close(jobCh)
	}()

	return jobCh
}
