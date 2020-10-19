package singleton

import (
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	/*
		t.Run("single-thread", func(t *testing.T) {
			a := GetInstance()
			b := GetInstance()
			if a != b {
				t.Errorf("Got 2 instances: %v, %v, expect 1", a, b)
			}
			b.Inc()
			b.Inc()
			if a.Count != 2 {
				t.Errorf("Got count %d, expect 2", a.Count)
			}
		})
	*/
	t.Run("multi-thread", func(t *testing.T) {
		var wg sync.WaitGroup
		iters := 100000
		for i := 0; i < iters; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				a := GetInstance()
				a.Inc()
			}()
		}
		wg.Wait()
		a := GetInstance()
		if a.Count != iters {
			t.Errorf("Got count %d, expect %d", a.Count, iters)
		}
	})
}

func TestCreateInstanceOnce(t *testing.T) {
	t.Run("multi-thread", func(t *testing.T) {
		var wg sync.WaitGroup
		iters := 100000
		for i := 0; i < iters; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				a := CreateInstanceOnce()
				a.Inc()
			}()
		}
		wg.Wait()
		a := CreateInstanceOnce()
		if a.Count != iters {
			t.Errorf("Got count %d, expect %d", a.Count, iters)
		}
	})
}
