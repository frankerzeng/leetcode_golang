package demo

import "fmt"

//Name dd
var Name int

const const_name = 1

func ll() int {
	return 1
}
func locate(namd *int) int {
	*namd = 1
	return *namd
}
func main() {
	var _ = "sdf"
	var name = "sdf"
	name1 := &name
	name = "123"
	fmt.Print(name)
	fmt.Print(name1)
	/* 定义局部变量 */
	var a int = 10

	/* 循环 */
LOOP:
	for a < 20 {
		if a == 15 {
			/* 跳过迭代 */
			a = a + 1
			goto LOOP
		}
		fmt.Printf("a的值为 : %d\n", a)
		a++
	}
	fmt.Println(ll())
	namd := 9
	locate(&namd)
	fmt.Println(&namd)
	fmt.Println(namd)

	// 空指针
	var ptr *int
	fmt.Println(ptr)

	ptr = &namd
	var pptr **int = &ptr

	fmt.Println(pptr)
}
