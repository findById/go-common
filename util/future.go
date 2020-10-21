package util

import "sync"

type Future struct {
	wg      sync.WaitGroup
	mutex   sync.RWMutex
	queue   []func()
	workers []func()
	started bool
}

func NewFuture() *Future {
	return &Future{
		wg: sync.WaitGroup{},
	}
}

func (f *Future) Then(worker func()) *Future {
	if f.started {
		panic("worker started")
	}
	f.mutex.Lock()
	f.queue = append(f.queue, worker)
	f.mutex.Unlock()
	return f
}

func (f *Future) ThenAll(workers ...func()) *Future {
	if f.started {
		panic("worker started")
	}
	f.mutex.Lock()
	f.workers = append(f.workers, workers...)
	f.mutex.Unlock()
	return f
}

func (f *Future) Do() {
	f.started = true
	defer func() {
		f.started = false
	}()
	if f.workers != nil && len(f.workers) > 0 {
		f.wg.Add(len(f.workers))
		for _, work := range f.workers {
			item := work
			go func() {
				defer f.wg.Done()
				item()
			}()
		}
		f.wg.Wait()
	}
	if f.queue != nil && len(f.queue) > 0 {
		for _, item := range f.queue {
			item()
		}
	}
}
