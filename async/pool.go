package async

import (
	"context"
	"sync"
)

type Task func() error

type Pool struct {
	count int
	tasks chan Task
	start sync.Once
	stop  sync.Once
	exit  chan struct{}
}

func New(count uint) *Pool {
	return &Pool{
		count: int(count),
		tasks: make(chan Task),
	}
}

func (self *Pool) Start() {
	self.StartWithContext(context.Background())
}

func (self *Pool) StartWithContext(ctx context.Context) {
	self.start.Do(func() {
		for i := 0; i < self.count; i++ {
			go self.listen(ctx)
		}
	})
}

func (self *Pool) Stop() {
	self.stop.Do(func() {
		close(self.exit)
	})
}

func (self *Pool) Push(task Task) {
	select {
	case self.tasks <- task:
	case <-self.exit:
	}
}

func (self *Pool) listen(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-self.exit:
			return
		case task, ok := <-self.tasks:
			if !ok {
				return
			}

			if err := task(); err != nil {
				return
			}
		}
	}
}
