package main

import (
	"fmt"
	"log"
	"runtime"
	"sort"
)

func main() {
	fmt.Println(multiReturn())
	fmt.Println(multiReturn2())
	addr()
	multiParam([]int{1, 2, 3}...)
	testDefer()
	innerFunc()
	returnFunc(addr)(1, 1)
	sortFunc()
	mapFunc()
}

// 多返回值
func multiReturn() (int, int) {
	return 1, 2
}
func multiReturn2() (rst int, rst2 int) {
	rst = 1
	rst2 = 2
	return
}

// 指针
func addr() {
	normal := 1
	addr := &normal
	fmt.Println(addr)
	fmt.Println(*addr)
}

// 传递多值
func multiParam(i ...int) {
	fmt.Println("传递多值", i)
}

// defer 关键字
func testDefer() {
	fmt.Println(1)
	defer fmt.Println(3) // 类似php的finally
	fmt.Println(2)
}

// 内置函数
func innerFunc() {
	// close
	// len cap
	// new make ---new(int):返回类型的一个地址
	// copy append
	// panic recover
	// print println
	// complex real imag
}

// 函数作为参数/返回值
func returnFunc(addr func()) func(int1 int, int2 int) {
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d", file, line)
	addr()
	return func(int1 int, int2 int) {
		fmt.Println("函数作为参数/返回值", int1, int2)
	}
}

// 排序函数
func sortFunc() {
	k := []int{2, 3, 1, 2, 1}
	sort.Ints(k)
	fmt.Println("排序", k)
	fmt.Println("排序", sort.SearchInts(k, 3))
}

// map
func mapFunc() {
	m := map[int]string{}
	m[0] = "sdd"
	m[2] = "sdd"
	fmt.Println("map", m)
	k := make(map[int]int, 2)
	k[1] = 1
	fmt.Println("map", k)

}