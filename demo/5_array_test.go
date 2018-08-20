package demo

import (
	"fmt"
)

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 7}
	s := a[3:6]
	s[0] = 8 // 切片->引用 注意
	fmt.Println(a)
}
