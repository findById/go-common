package pool

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestNewWorkerPool(t *testing.T) {
	s := time.Now().UnixNano() / 1e6
	worker := NewWorkerPool(1000, 100)
	wg := sync.WaitGroup{}
	for i := 0; i <= 10000; i++ {
		wg.Add(1)
		go func(i int) {
			worker.Submit(func() {
				time.Sleep(time.Duration(10000) * 60)
			})
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println("=========", time.Now().UnixNano()/1e6)
	worker.Stop()
	fmt.Println("=========", time.Now().UnixNano()/1e6)
	fmt.Printf("used %vms\n", time.Now().UnixNano()/1e6-s)
}
