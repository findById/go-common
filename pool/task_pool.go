package pool

import (
	"context"
	"errors"
	"math/rand"
	"sync"
)

type WorkerPool struct {
	maxWorker  int // 最大队列数
	queueSize  int // 单个队列最大容量
	taskQueue  []chan func()
	ctx        context.Context
	cancelFunc func()
	wg         sync.WaitGroup
	stopped    bool
}

func NewWorkerPool(maxWorker, queueSize int) *WorkerPool {
	pool := &WorkerPool{
		maxWorker: maxWorker,
		queueSize: queueSize,
		taskQueue: make([]chan func(), maxWorker),
		wg:        sync.WaitGroup{},
	}
	ctx, cancelFunc := context.WithCancel(context.Background())
	pool.ctx = ctx
	pool.cancelFunc = cancelFunc
	pool.dispatch()
	return pool
}

func (p *WorkerPool) dispatch() {
	for i := 0; i < p.maxWorker; i++ {
		p.taskQueue[i] = make(chan func(), p.queueSize)
		p.startWorker(p.taskQueue[i])
	}
}

func (p *WorkerPool) startWorker(taskChan chan func()) {
	go func() {
		var task func()
		var ok bool
		for {
			select {
			case task, ok = <-taskChan:
				if !ok {
					continue
				}
				//defer p.wg.Done()
				task()
				p.wg.Done()
			case <-p.ctx.Done():
				close(taskChan)
				return
			}
		}
	}()
}

func (p *WorkerPool) Submit(task func()) error {
	if p.stopped {
		return errors.New("worker is stopped")
	}
	if task != nil {
		idx := rand.New(rand.NewSource(int64(p.maxWorker))).Int() % (p.maxWorker)
		p.wg.Add(1)
		p.taskQueue[idx] <- task
	}
	return nil
}

func (p *WorkerPool) Stop() {
	p.stopped = true
	p.wg.Wait()
	p.cancelFunc()
}
