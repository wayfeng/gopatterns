package singleton

import (
	"fmt"
	"sync"
)

type Counter struct {
	Count int
	mu    sync.Mutex
}

type Increase interface {
	Inc() int
}

func (c *Counter) Inc() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.Count++
	return c.Count
}

var instance *Counter

// Naive version doesn't work in multi-threaded program
func GetInstance() *Counter {
	if instance == nil {
		fmt.Println("Create new counter")
		instance = &Counter{Count: 0}
	}
	return instance
}

var mu sync.Mutex

func CreateInstanceMutex() *Counter {
	mu.Lock()
	defer mu.Unlock()
	if instance == nil {
		fmt.Println("Create new counter")
		instance = &Counter{Count: 0}
	}
	return instance
}

var once sync.Once

func CreateInstanceOnce() *Counter {
	once.Do(func() {
		instance = &Counter{Count: 0}
	})
	return instance
}
