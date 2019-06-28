package main

import (
	"fmt"
	"sync"
)

var counter = struct {
	sync.RWMutex
	m map[string]int
}{m: make(map[string]int)}

func main() {

	counter.Lock()
	counter.m["gwr"] = 100
	counter.Unlock()

	counter.RLock()
	defer counter.RUnlock()
	n := counter.m["gwr"]

	fmt.Println(n)

	var map2 sync.Map

	map2.Store("a", "b")
	res, _ := map2.Load("a")



	fmt.Println(res)

}
