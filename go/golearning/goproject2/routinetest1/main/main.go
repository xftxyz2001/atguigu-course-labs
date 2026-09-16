package main

import (
	"fmt"
	"sort"
	"sync"
)

var (
	m    map[int]uint64 = make(map[int]uint64, 10)
	lock sync.Mutex
)

// 计算阶乘
func factorial(n int) {
	var res uint64 = 1
	for i := n; i > 0; i-- {
		res *= uint64(i)
	}
	lock.Lock()
	m[n] = res
	lock.Unlock()

}

func main() {
	for i := 0; i < 200; i++ {
		go factorial(i) // concurrent map writes
	}

	lock.Lock()
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	lock.Unlock()

	sort.Ints(keys)
	for _, k := range keys {
		fmt.Printf("m[%v] = %v\n", k, m[k])
	}

}
