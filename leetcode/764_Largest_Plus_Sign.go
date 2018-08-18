package main

import (
	"fmt"
	"time"
)

func main() {
	var list_ [][]int
	N := 5

	list_ = append(list_, []int{0, 2})
	list_ = append(list_, []int{1, 0})
	list_ = append(list_, []int{2, 0})

	N = 500
	list_ = [][]int{{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4}, {0, 5}, {0, 6}, {0, 8}, {0, 9}, {0, 10}, {0, 11}, {0, 12}, {0, 13}, {0, 14}, {0, 16}, {0, 17}, {0, 18}, {0, 19},
		{1, 1}, {1, 2}, {1, 3}, {1, 4}, {1, 6}, {1, 8}, {1, 9}, {1, 10}, {1, 11}, {1, 12}, {1, 13}, {1, 14}, {1, 15}, {1, 17}, {1, 19},
		{2, 1}, {2, 4}, {2, 5}, {2, 6}, {2, 7}, {2, 8}, {2, 9}, {2, 11}, {2, 12}, {2, 13}, {2, 14}, {2, 15}, {2, 16}, {2, 17}, {2, 18}, {2, 19},
		{3, 0}, {3, 1}, {3, 3}, {3, 4}, {3, 5}, {3, 6}, {3, 7}, {3, 8}, {3, 11}, {3, 13}, {3, 14}, {3, 15}, {3, 16}, {3, 17}, {3, 18}, {3, 19},
		{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 5}, {4, 7}, {4, 8}, {4, 9}, {4, 10}, {4, 11}, {4, 12}, {4, 13}, {4, 14}, {4, 16}, {4, 17}, {4, 18}, {4, 19},
		{5, 0}, {5, 1}, {5, 3}, {5, 4}, {5, 5}, {5, 6}, {5, 7}, {5, 8}, {5, 9}, {5, 10}, {5, 11}, {5, 13}, {5, 14}, {5, 15}, {5, 16}, {5, 17}, {5, 19},
		{6, 0}, {6, 1}, {6, 2}, {6, 4}, {6, 5}, {6, 6}, {6, 7}, {6, 8}, {6, 9}, {6, 12}, {6, 13}, {6, 15}, {6, 16}, {6, 17}, {6, 18}, {6, 19}, {7, 1},
		{7, 3}, {7, 5}, {7, 6}, {7, 7}, {7, 8}, {7, 9}, {7, 10}, {7, 11}, {7, 12}, {7, 13}, {7, 14}, {7, 16}, {7, 17}, {7, 18}, {7, 19},
		{8, 0}, {8, 3}, {8, 6}, {8, 7}, {8, 8}, {8, 10}, {8, 11}, {8, 13}, {8, 15}, {8, 16}, {8, 17}, {8, 18}, {8, 19},
		{9, 0}, {9, 1}, {9, 3}, {9, 4}, {9, 5}, {9, 7}, {9, 9}, {9, 10}, {9, 11}, {9, 12}, {9, 14}, {9, 15}, {9, 16}, {9, 17}, {9, 18}, {9, 19},
		{10, 0}, {10, 2}, {10, 3}, {10, 4}, {10, 5}, {10, 6}, {10, 7}, {10, 8}, {10, 9}, {10, 10}, {10, 11}, {10, 12}, {10, 13}, {10, 14}, {10, 15}, {10, 16}, {10, 19},
		{11, 1}, {11, 2}, {11, 3}, {11, 4}, {11, 6}, {11, 7}, {11, 9}, {11, 11}, {11, 12}, {11, 13}, {11, 16}, {11, 18}, {11, 19},
		{12, 0}, {12, 1}, {12, 2}, {12, 3}, {12, 5}, {12, 7}, {12, 8}, {12, 9}, {12, 10}, {12, 11}, {12, 12}, {12, 13}, {12, 14}, {12, 15}, {12, 16}, {12, 17}, {12, 18}, {12, 19},
		{13, 0}, {13, 1}, {13, 3}, {13, 4}, {13, 5}, {13, 6}, {13, 7}, {13, 8}, {13, 12}, {13, 13}, {13, 14}, {13, 15}, {13, 16}, {13, 18}, {13, 19},
		{14, 0}, {14, 1}, {14, 2}, {14, 3}, {14, 4}, {14, 5}, {14, 6}, {14, 7}, {14, 8}, {14, 11}, {14, 12}, {14, 13}, {14, 14}, {14, 15}, {14, 16}, {14, 17}, {14, 18}, {14, 19},
		{15, 0}, {15, 1}, {15, 3}, {15, 4}, {15, 5}, {15, 7}, {15, 8}, {15, 9}, {15, 10}, {15, 13}, {15, 14}, {15, 15}, {15, 16}, {15, 17}, {15, 18}, {15, 19},
		{16, 0}, {16, 1}, {16, 2}, {16, 3}, {16, 6}, {16, 7}, {16, 8}, {16, 9}, {16, 10}, {16, 11}, {16, 13}, {16, 14}, {16, 15}, {16, 16}, {16, 17}, {16, 18}, {16, 19},
		{17, 0}, {17, 1}, {17, 2}, {17, 4}, {17, 6}, {17, 8}, {17, 9}, {17, 10}, {17, 12}, {17, 13}, {17, 14}, {17, 16}, {17, 17}, {17, 18},
		{18, 0}, {18, 3}, {18, 4}, {18, 5}, {18, 6}, {18, 8}, {18, 9}, {18, 11}, {18, 12}, {18, 13}, {18, 14}, {18, 15}, {18, 16}, {18, 18}, {18, 19},
		{19, 0}, {19, 1}, {19, 2}, {19, 3}, {19, 4}, {19, 5}, {19, 6}, {19, 7}, {19, 8}, {19, 9}, {19, 10}, {19, 12}, {19, 13}, {19, 14}, {19, 15}, {19, 17}, {19, 18}}
	t1 := time.Now()
	fmt.Println(orderOfLargestPlusSign(N, list_))
	elapsed := time.Since(t1)
	fmt.Println("App elapsed: ", elapsed)
}

//  方法一 Time Limit Exceeded
func orderOfLargestPlusSign_0(N int, mines [][]int) int {
	var Num = 0
	var NumTmp = 0
	var NumIndex = 0 // 向周围延伸的触手长度
	// 判断坐标是否是0
	var isNoZero = func(x int, y int) bool {
		for _, val := range mines {
			if val[0] == x && val[1] == y {
				return false
			}
		}
		return true
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			NumTmp = 0
			if isNoZero(i, j) {
				NumTmp = 1
				NumIndex = 0
				for true {
					NumIndex++
					// 边界条件
					if i < NumIndex || j < NumIndex || i+NumIndex >= N || j+NumIndex >= N {
						break
					}

					// 四个触手都不是0
					if isNoZero(i, j+NumIndex) && isNoZero(i, j-NumIndex) &&
						isNoZero(i-NumIndex, j) && isNoZero(i+NumIndex, j) {
						NumTmp++
					} else {
						break
					}
				}
			}
			if NumTmp > Num {
				Num = NumTmp
			}
		}
	}

	return Num
}

// 方法二  Time Limit Exceeded
func orderOfLargestPlusSign_1(N int, mines [][]int) int {
	var Num = 0
	var NumTmp = 0
	var NumIndexMax = 0 // 可延伸的最大值

	// cell value = 0
	ZeroMap := map[int]int{}
	for _, val := range mines {
		ZeroMap[val[0]*N+val[1]] = 1
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			NumTmp = 0
			if ZeroMap[j+i*N] == 0 {
				NumTmp = 1
				NumIndexMax = j + 1
				if i < j {
					NumIndexMax = i + 1
				}
				if N-i < NumIndexMax {
					NumIndexMax = N - i
				}
				if N-j < NumIndexMax {
					NumIndexMax = N - j
				}
				for ii := 1; ii < NumIndexMax; ii++ {
					// 四个触手都不是0
					if ZeroMap[i*N+(j+ii)] == 0 && ZeroMap[i*N+(j-ii)] == 0 &&
						ZeroMap[(i-ii)*N+j] == 0 && ZeroMap[(i+ii)*N+j] == 0 {
						NumTmp++
					} else {
						break
					}
				}
			}
			if NumTmp > Num {
				Num = NumTmp
			}
		}
	}

	return Num
}

func orderOfLargestPlusSign(N int, mines [][]int) int {
	index := 0
	Num := 0
	// cell value = 0
	ZeroMap := map[int]int{}
	for _, val := range mines {
		ZeroMap[val[0]*N+val[1]] = 1
	}

	// every cell value
	ValueMap := map[int]int{}

	// horizontal
	for i := 0; i < N; i++ {
		index = 0
		for j := 0; j < N; j++ {
			if ZeroMap[i*N+j] == 0 {
				index++
				ValueMap[i*N+j] = index
			} else {
				index = 0
				ValueMap[i*N+j] = 0
			}
		}
		index = 0
		for j := N - 1; j >= 0; j-- {
			if ZeroMap[i*N+j] == 0 {
				index++
				if index < ValueMap[i*N+j] {
					ValueMap[i*N+j] = index
				}
			} else {
				index = 0
				ValueMap[i*N+j] = 0
			}
		}
	}

	// vertical
	for i := 0; i < N; i++ {
		index = 0
		for j := 0; j < N; j++ {
			if ZeroMap[i+j*N] == 0 {
				index++
				if index < ValueMap[i+j*N] {
					ValueMap[i+j*N] = index
				}
			} else {
				index = 0
				ValueMap[i+j*N] = 0
			}
		}
		index = 0
		for j := N - 1; j >= 0; j-- {
			if ZeroMap[i+j*N] == 0 {
				index++
				if index < ValueMap[i+j*N] {
					ValueMap[i+j*N] = index
				}
			} else {
				index = 0
				ValueMap[i+j*N] = 0
			}

			// ValueMap[i+j*N] is final value of i,j
			if ValueMap[i+j*N] > Num {
				Num = ValueMap[i+j*N]
			}
		}
	}

	return Num
}
