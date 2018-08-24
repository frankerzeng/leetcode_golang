package demo

import "fmt"

func main() {
	// 数组 ，固定长度
	var arr = [...]float32{1.0, 32.3}
	b := []int{2: 1, 3: 2}
	fmt.Print(b)
	fmt.Print(arr)

	// 切片，长度可变
	//var arrN = []float32{1.0, 32.3}
	var arrN []int
	arrN = make([]int, 3)
	arrN = append(arrN, 145)
	fmt.Println(arrN)
	fmt.Println(len(arrN))
	fmt.Println(cap(arrN))

	for key, value := range arrN {
		println(key)
		println(value)
	}

	// map 无序集合
	var mapN map[string]string
	mapN = make(map[string]string)
	mapN["sdf"] = "sdf"
	fmt.Println(mapN)

	mapName := map[string]string{"s": "sd", "34": "sdf"}
	fmt.Println(mapName)
}
