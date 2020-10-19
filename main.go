package main

import (
	"fmt"
	"gopatterns/singleton"
	"sync"
)

func main() {
	a := singleton.GetInstance()
	a.Inc()
	fmt.Println(a.Count)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		a := singleton.GetInstance()
		a.Inc()
	}()
	wg.Wait()
	fmt.Println(a.Count)
}
