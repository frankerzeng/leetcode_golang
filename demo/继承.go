package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
	"time"
)

func init() { // 初始化包
	fmt.Println("init")
}

type Test struct {
	Name string `json:"oame"` //  标签
	Kame string `json:"pame"` //  标签
}

func maketest(test *Test) {
	test.Name = "44"
}

type Grandfather struct {
	Very string
}
type Father struct {
	Name string
	age  int
	Grandfather
}

func (father *Father) getName() string {
	return father.Name
}

type son struct {
	Father
	Grandfather
	int // 匿名字段
}

type interfaceTest interface { // 接口
	test() int
}

func main() {
	sort.Sort()
	// 继承
	s := son{Father: Father{Name: "sf", age: 34}}
	s.age = 12
	fmt.Println(s.Name, s.Father.Name, s.age, s.Father.age)
	fmt.Println(s.getName())

	// 多重继承
	s.Very = "44"
	fmt.Println(s.Very)

	return

	// 继承

	tkk := &Test{Name: "sdf444", Kame: "sftt"}
	fmt.Println(tkk)
	rst, _ := json.Marshal(tkk)

	fmt.Println(rst)
	fmt.Println(string(rst))

	var mapVar map[string]string
	mapVar = make(map[string]string)
	mapVar["dd"] = "dsd"
	mapVarp := &mapVar
	(*mapVarp)["dd"] = "sdf"
	fmt.Println(*mapVarp)
	return
	var tvv Test
	tvv.Name = "ds"
	maketest(&tvv)
	fmt.Println(tvv)

	return

	fmt.Println("---")
	kkMap := make(map[string]string, 1)
	kkMap["d"] = "ddd"
	kkMap["dd"] = "ddd"
	fmt.Println(kkMap)
	fmt.Println("---")
	return
	kmap := map[string][]int{}
	kmap["s"] = []int{1, 2}
	fmt.Println(kmap)

	var newName name11 = 12
	fmt.Println(newName)

	kk()

	fmt.Println(tt())

	type person struct {
		name int `json:"namei"`
	}
	p := person{1}
	bytes, _ := json.Marshal(p)
	fmt.Println(string(bytes))

	fmt.Println((time.Now()).Format("2006-01-02 15:04:05"))

	for i := 0; i < 100; i++ {
		fmt.Println((time.Now()).UnixNano())
	}

	t := 1
	kk := fmt.Sprintf("%T", t)
	fmt.Printf(kk)
	fmt.Println()

	fmt.Println(reflect.TypeOf(t).String())

	ii := new(int)
	*ii = 100
	fmt.Println(ii, &ii)

	fmt.Println("---")
	slilll := []int{1, 2}
	slilll2 := slilll[0:1]
	slilll2[0] = 11
	fmt.Println(slilll)
	fmt.Println(slilll2)

	skk := [...]int{1, 2, 3}
	skkd := skk[0:1]
	skkd[1] = 11
	fmt.Println(skk)
	fmt.Println(skkd)
	fmt.Println("=====")

	var intArr [2]int
	fmt.Println(&intArr[0], &intArr[1])

	intSlice := []int{1, 2}
	intSlice = append(intSlice, 2, 3)
	intSlice = append(intSlice, 2, 3)
	fmt.Println(&intSlice[0], &intSlice[1], &intSlice[2], &intSlice[3])

}

func tst(i int, f func(kk int) int) int { // 函数作为形参
	fmt.Println(f(2))
	return 1
}

func kk() {
	var f = func(kk int) int {
		fmt.Println(11, kk)
		return 12
	}
	tst(1, f)
}

type name11 int // 取别名

var name string = "1"

func tt() (n int) {
	defer func(n *int) {
		*n++
		fmt.Println(*n)
	}(&n)

	for i := 1; i < 3; i++ {

	}

	fmt.Println(name)
	f := strings.Compare("，", "s")
	fmt.Println(f)

	var m map[string]int
	m = make(map[string]int) // 要初始化
	m["sd"] = 2

	// 数组容量
	mm := make([]int, 1, 8)
	mm = append(mm, 23)
	fmt.Println(len(mm), cap(mm), "---", mm)

	mmGuding := [...]int{12, 3, 2}
	mmGuding[2] = 22

	fmt.Println(mmGuding)

	return
}
