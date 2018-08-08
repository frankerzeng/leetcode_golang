// 结构体
package main

import "fmt"

type test struct {
	speed int
}

func (test test) getSpeed() int {
	return test.speed
}
func main() {
	var t test
	t.speed = 1

	fmt.Print(t.getSpeed())
}
