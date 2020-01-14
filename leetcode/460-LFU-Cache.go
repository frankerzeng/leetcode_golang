/**
460. LFU Cache

Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.

Note that the number of times an item is used is the number of calls to the get and put functions for that item since it was inserted. This number is set to zero when the item is removed.
*/
package main

import (
	"fmt"
)

type LFUCache struct {
	ReadTime        int // 读取时间，后读取的更大
	AddTime         int // 读取时间，后读取的更大
	Cap             int
	K2V             map[int]int
	K2LastReadTime  map[int]int
	K2LastReadTimes map[int]int // 使用次数
	K2NoRead        map[int]int // 未读取的数
}

func Constructor(capacity int) LFUCache {
	return LFUCache{Cap: capacity, K2V: make(map[int]int), K2LastReadTime: make(map[int]int), K2LastReadTimes: make(map[int]int), K2NoRead: make(map[int]int)}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.K2V[key]; ok {
		this.ReadTime++
		this.K2LastReadTime[key] = this.ReadTime
		this.K2LastReadTimes[key]++
		delete(this.K2NoRead, key)
		return v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap == 0 {
		return
	}
	this.AddTime++
	if _, ok := this.K2V[key]; ok {
		this.K2V[key] = value
		if _, ok := this.K2NoRead[key]; ok {
			this.K2NoRead[key] = this.AddTime
		}
		if this.K2V[key] != value {
			delete(this.K2LastReadTime, key)
		}
	} else {
		if len(this.K2V) == this.Cap {
			vTime, kTime := 9999999, -1
			if len(this.K2NoRead) > 0 {
				for k, v := range this.K2NoRead { // 确定淘汰的下标
					if v < vTime {
						vTime = v
						kTime = k
					}
				}
			} else {
				for k, v := range this.K2LastReadTime { // 确定淘汰的下标
					if v < vTime {
						vTime = v
						kTime = k
					}
				}
			}
			delete(this.K2V, kTime)
			delete(this.K2LastReadTime, kTime)
			delete(this.K2LastReadTimes, kTime)
			delete(this.K2NoRead, kTime)
		}
		this.K2V[key] = value
		this.K2NoRead[key] = this.AddTime
		this.K2LastReadTimes[key] = 0
	}
}

/*
todo
*/
func main() {
	obj := Constructor(10)
	obj.Put(10, 13)
	obj.Put(3, 17)
	obj.Put(6, 11)
	obj.Put(10, 5)
	obj.Put(9, 10)
	obj.Get(13)
	fmt.Println(obj)

	obj.Put(2, 19)
	fmt.Println("1_--------")
	fmt.Println(obj)

	obj.Get(2)
	obj.Get(3)
	fmt.Println("2_--------")
	fmt.Println(obj)

	obj.Put(5, 25)
	obj.Get(8)
	fmt.Println("3_--------")
	fmt.Println(obj)

	obj.Put(9, 22)
	obj.Put(5, 5)
	obj.Put(1, 30)
	obj.Get(11)
	fmt.Println("452_--------")
	fmt.Println(obj)

	obj.Put(9, 12)
	obj.Get(7)
	obj.Get(5)
	obj.Get(8)
	obj.Get(9)
	fmt.Println("52_--------")
	fmt.Println(obj)

	obj.Put(4, 30)
	obj.Put(9, 3)
	obj.Get(9)
	obj.Get(10)
	obj.Get(10)
	fmt.Println("2343_--------")
	fmt.Println(obj)

	obj.Put(6, 14)
	obj.Put(3, 1)
	obj.Get(3)

	obj.Put(10, 11)
	obj.Get(8)
	fmt.Println("4542_--------")
	fmt.Println(obj)

	obj.Put(2, 14)
	obj.Get(1)
	obj.Get(5)
	obj.Get(4)
	fmt.Println("2435_--------")
	fmt.Println(obj)

	obj.Put(11, 4)
	obj.Put(12, 24)
	obj.Put(5, 18)
	obj.Get(13)
	fmt.Println("2ereer_--------")
	fmt.Println(obj)

	obj.Put(7, 23)
	obj.Get(8)
	obj.Get(12)
	fmt.Println("2_-3434-------")
	fmt.Println(obj)

	obj.Put(3, 27)
	obj.Put(2, 12)
	obj.Get(5)
	fmt.Println("2asdfsdf_--------")
	fmt.Println(obj)

	obj.Put(2, 9)
	obj.Put(13, 4)
	obj.Put(8, 18)
	obj.Put(1, 7)
	fmt.Println(obj.Get(6))
	fmt.Println("1212_--------")
	fmt.Println(obj)

}
