package main

import (
	json2 "encoding/json"
	"fmt"
)

func main() {
	var s Say = &Stu{}
	fmt.Println(s)

	var a interface{}
	a = Point{3}
	var b Point
	b = a.(Point) // 类型断言
	fmt.Println(b)
	fmt.Println("0----")

	fmt.Println(Jsonf())
}

type Say interface {
	Say()
}

type Stu struct {
}

func (stu *Stu) Say() {
	fmt.Println(1)
}

type Point struct {
	x int
}

type JsonTT struct {
	Name   string
	Salary []int
	Age    int
}

func Jsonf() string {
	var j = JsonTT{"s", []int{2, 3}, 2}
	//j.age = 1
	//j.name = "zfl"
	jstring, _ := json2.Marshal(&j)
	fmt.Println(string(jstring))

	return string(jstring)
}
