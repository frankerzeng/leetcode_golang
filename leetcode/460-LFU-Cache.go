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
	ReadTime       int // 读取时间，后读取的更大
	Cap            int
	K2V            map[int]int
	K2LastReadTime map[int]int
	K2NoRead       map[int]bool // 未读取的数
}

func Constructor(capacity int) LFUCache {
	return LFUCache{Cap: capacity, K2V: make(map[int]int), K2LastReadTime: make(map[int]int), K2NoRead: make(map[int]bool)}
}

func (this *LFUCache) Get(key int) int {
	if v, ok := this.K2V[key]; ok {
		this.ReadTime++
		this.K2LastReadTime[key] = this.ReadTime
		delete(this.K2NoRead, key)
		return v
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.Cap == 0 {
		return
	}
	if _, ok := this.K2V[key]; ok {
		this.K2LastReadTime[key] = 0
		this.K2V[key] = value
		delete(this.K2NoRead, key)
		this.K2NoRead[key] = true
	} else {
		if len(this.K2V) == this.Cap {
			vTime, kTime := 10000, -1
			if len(this.K2NoRead) > 0 {
				for k := range this.K2NoRead { // 确定淘汰的下标
					kTime = k
					break
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
		}
		this.K2V[key] = value
		this.K2LastReadTime[key] = 0
		this.K2NoRead[key] = true
	}
}

func main() {
	obj := Constructor(2)
	obj.Put(1, 1)
	fmt.Println(obj)

	obj.Put(2, 2)
	fmt.Println(obj)

	obj.Get(1)
	fmt.Println(obj)

	obj.Put(3, 3)
	fmt.Println(obj)

	fmt.Println("----------")

	obj.Get(2)
	fmt.Println(obj)

	obj.Get(3)
	fmt.Println(obj)

	obj.Put(4, 4)
	fmt.Println(obj)

	obj.Get(1)
	fmt.Println(obj)

}
