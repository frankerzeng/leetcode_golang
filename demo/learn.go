package main

import (
	"./array4"
	"container/list"
	"encoding/gob"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"runtime"
	"sdf"
	"sort"
	"sync"
	"time"
	"unsafe"
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
	arrayAndSlice()
	listFunc()
	regFunc()
	// lockFunc()
	multiAccess()
	interfaceFunc()
	jsonFunc()
	channlFunc()
	selectFun()
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

// 数组 切片
func arrayAndSlice() {
	arr := [...]int{1, 2, 3} // 固定长度
	sli := []int{1, 2, 3}    // 可变长度
	fmt.Println("数组 切片", arr, sli)
}

// 链表
func listFunc() {
	l := list.New()
	l.PushBack(10)
	l.PushBackList(l)
	fmt.Println("链表", l.Front().Value)
}

// 正则
func regFunc() {
	searchStr := "ljljl 1023.33 and 109.4"
	pat := "\\d+.\\d+"
	ok, _ := regexp.Match(pat, []byte(searchStr))

	fmt.Println("正则", ok)

	reObj, _ := regexp.Compile(pat)
	str := reObj.ReplaceAllString(searchStr, "11.11")

	fmt.Println("正则", str)
}

type lockInfo struct {
	mu  sync.Mutex
	Str string
}

var m *sync.RWMutex

// 锁
func lockFunc() {
	var lock_info lockInfo
	lock_info.Str = "ddd"
	fmt.Println(lock_info)
	m = new(sync.RWMutex)

	// 多个同时读 独占锁
	go readLock(1)
	go readLock(2)

	time.Sleep(4 * time.Second)

	// 多个同时读 共享锁
	go readRLock(1)
	go readRLock(2)

	time.Sleep(4 * time.Second)
}

func readLock(i int) {
	println(i, "read start")
	time.Sleep(1 * time.Second)

	m.Lock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.Unlock()

	println(i, "read over")
}
func readRLock(i int) {
	println(i, "read start")
	time.Sleep(1 * time.Second)

	m.RLock()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	m.RUnlock()

	println(i, "read over")
}

type CatTest struct{}

func (c *CatTest) Mu() string {
	return "cat mu"
}

type DogTest struct{}

func (p *DogTest) Buf() string {
	return "dog buf"
}

type CameraPhone struct {
	CatTest
	DogTest
}

func multiAccess() {
	cp := new(CameraPhone)
	fmt.Println("多重继承", cp.Mu()) // 等效 fmt.Println(cp.CatTest.Mu())
	fmt.Println("多重继承", cp.Buf())
}

type Simple struct {
	name int
}

func (simple *Simple) Get() int {
	return simple.name
}
func (simple *Simple) Set(name int) {
	simple.name = name
}

type Simpler interface {
	Get() int
	Set(name int)
}

func interfaceFunc() {
	var sim Simpler = &Simple{1}
	fmt.Println("接口------")
	fmt.Println(sim.Get())
	sim.Set(3)
	fmt.Println(sim.Get())
	i := 1
	test1(&i)
}

func test1(ptr *int) {
	var t uint = 1
	s, _ := fmt.Printf("%T", t)
	fmt.Println(s)
	fmt.Printf("%T,%d\n", t, unsafe.Sizeof(t))
	var tt int = 'g'
	fmt.Println(tt)
	var i float64 = 23.3333333 + 20
	j := byte(i)
	var n int8 = 20
	n = n + 90
	fmt.Println("n=", n)
	fmt.Println(j)

	int1 := -2 & -3 // 运算都是补码运算
	fmt.Println(int1)

	if int1++; int1 > -10 { // if 的缩略写法
		fmt.Println(int1, "aa")
	}
	int1 = 1
	switch {
	case int1 < 101:
		fmt.Println(111)
		fallthrough
	case int1 < 10:
		fmt.Println(222)
	}

	var x interface{} // 变量类型
	var y int = 1
	x = y
	switch x.(type) {
	case bool:
		fmt.Println("dsd")
	case int:
		fmt.Println("int")

	}

	// rune类型 切片 类似int32
	st := "dssfd"
	st1 := []rune(st)
	fmt.Println(st1)
	count := 0
	for {
		rand.Seed(time.Now().UnixNano())
		ii := rand.Intn(100000) + 1
		count++
		if ii == 99 {
			break
		}
	}
	fmt.Println(count)

	count = 0
lab: // 跳出标签
	for {
		for {
			count++
			if count == 10 {
				break lab
			}
		}
	}
	fmt.Println(count)

	count = 1
	for i0 := 0; i0 < 3; i0++ {
		fmt.Println(1)
		for i00 := 0; i00 < 3; i00++ {
			fmt.Println(2)
			count++
			//if count == 10 {
			goto lab1
			//}
		}
		break
	}
lab1:
	fmt.Println(1)

	array4.ArrayFun()
}

type Address struct {
	Type    string
	City    string
	Country string
}

type VCard struct {
	FirstName string
	LastName  string
	Addresses []*Address
	Remark    string
}

func jsonFunc() {
	pa := &Address{"private", "Aartselaar", "Belgium"}
	wa := &Address{"work", "Boom", "Belgium"}
	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	// fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	// JSON format:
	js, _ := json.Marshal(vc)
	fmt.Printf("JSON format: %s", js)
	// using an encoder:
	file, _ := os.OpenFile("testFile/vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := gob.NewEncoder(file)
	err := enc.Encode(vc)
	if err != nil {
		log.Println("Error in encoding json")
	}
	go goFunc()
	fmt.Println("d")
}

func goFunc() {
	time.Sleep(2)
	fmt.Println("gofunc")
	time.Sleep(2)
}

func Productor(mychan chan int, data int, wait *sync.WaitGroup) {
	mychan <- data
	fmt.Println("product data：", data)
	wait.Done()
}
func Consumer(mychan chan int, wait *sync.WaitGroup) {
	a := <-mychan
	fmt.Println("consumer data：", a)
	wait.Done()
}

func channlFunc() {
	datachan := make(chan int, 100) // 通讯数据管道
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		go Productor(datachan, i, &wg) // 生产数据
		wg.Add(1)
	}
	for j := 0; j < 10; j++ {
		go Consumer(datachan, &wg) // 消费数据
		wg.Add(1)
	}
	wg.Wait()

	ch := make(chan int)
	go test(ch)
	for j := 0; j < 10; j++ {
		fmt.Println("get ", <-ch)
	}

	requests := make(chan int, 2)
	for i := 1; i < 2; i++ {
		requests <- i
	}
	close(requests)
	limiter := time.Tick(time.Second * 1) // 计时器
	for req := range requests {
		<-limiter
		fmt.Println("requets", req, time.Now()) // 执行到这里，需要隔1秒才继续往下执行，time.Tick(timer)上面已定义
	}

	channel := make(chan int, 1)
	channel <- 2
	// close(channel)

	for data := range channel {
		fmt.Println(data, "sd")
		break
	}

	c := make(chan int, 3)
	var send chan<- int = c // send-only
	send <- 2
	fmt.Println(send)

	// 接口，可定义任意类型的map
	mixedData2 := make(map[string]interface{})
	mixedData2["username"] = "YamiOdymel"
	mixedData2["time"] = 123456
	fmt.Println(mixedData2)
	fmt.Println(mixedData2["usedsrname"])
	username, exists := mixedData2["usedsrname"]
	fmt.Println(username, exists)

	goto b
	fmt.Println("goto 1")
b:
	fmt.Println("goto 2")
	t := ConfigOne{"daemon"}
	fmt.Println(t.String())

	fmt.Println(len("你好bj!"))
}

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c.Daemon)
}

type student struct {
	Name string
	Age  int
}

func test(c chan int) int {
	for i := 0; i < 10; i++ {
		time.Sleep(1)
		fmt.Println("send ", i)
		c <- i
	}

	m := make(map[string]*student)
	stus := []student{
		{Name: "zhou", Age: 24},
		{Name: "li", Age: 23},
		{Name: "wang", Age: 22},
	}
	for _, stu := range stus {
		fmt.Println(stu)
		m[stu.Name] = &stu
	}

	runtime.GOMAXPROCS(1)

	count := 10
	wg := sync.WaitGroup{}
	wg.Add(count * 2)
	for i := 0; i < count; i++ {
		fmt.Println(i, "+")
		go func() {
			fmt.Printf("[%d]", i)
			wg.Done()
		}()
	}
	for i := 0; i < count; i++ {
		fmt.Println(i, "-")
		go func(i int) {
			fmt.Printf("-%d-", i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	return 2
}

var (
	B = 2
)

func selectFun() {
	start := time.Now()
	c := make(chan interface{})

	go func() {
		time.Sleep(2 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}

	i := 1
	switch i {
	case 1:
		fmt.Println(1)
	default:
		fmt.Println(2)
	}
	fmt.Println(B)

	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))

	const (
		aa = iota
		bb = 1
		dd
	)
	fmt.Println(aa, bb, dd)

	runtime.GOMAXPROCS(2)
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	int_chan <- 1
	string_chan <- "hello"
	select {
	case value := <-int_chan:
		fmt.Println(value, "22")
	case value := <-string_chan:
		fmt.Println(value, "11")
	}

	mapArr := []int{1, 2, 3, 4, 33333}
	fmt.Println(&mapArr[0])
	fmt.Println(&mapArr[1])
	fmt.Println(&mapArr[2])
	fmt.Println(&mapArr[3])
	fmt.Println(&mapArr[4], "====================================")

	slice := []int{0, 1, 2, 3}
	myMap := make(map[int]*int)

	for index, value := range slice {
		fmt.Println(&value, value)
		myMap[index] = &value
	}
	fmt.Println(slice)

	fmt.Println("=====new map=====")
	fmt.Println(myMap[0])
	fmt.Println(myMap[1])
	fmt.Println(myMap[2])
	fmt.Println(myMap[3])

}

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		var result string
		result, err = tryTheThing()
		fmt.Println(err)
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}
