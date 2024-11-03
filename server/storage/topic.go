package storage

import (
	"encoding/json"
	"slices"
	"sync"
	"time"

	"github.com/thegogod/fmq/common/protocol"
)

type Topic struct {
	mu        sync.RWMutex
	index     int
	queue     Queue[*protocol.Publish]
	consumers []protocol.Connection
}

func NewTopic() *Topic {
	self := &Topic{
		index:     -1,
		queue:     make(Queue[*protocol.Publish], 10000),
		consumers: []protocol.Connection{},
	}

	go self.listen()
	return self
}

func (self *Topic) Count() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return len(self.consumers)
}

func (self *Topic) Next() (protocol.Connection, bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	if len(self.consumers) == 0 {
		return nil, false
	}

	i := self.index + 1

	if i > (len(self.consumers) - 1) {
		i = 0
	}

	self.index = i
	return self.consumers[i], true
}

func (self *Topic) Subscribe(conn protocol.Connection) {
	self.mu.RLock()
	exists := slices.ContainsFunc(self.consumers, func(c protocol.Connection) bool {
		return c.ID() == conn.ID()
	})

	self.mu.RUnlock()

	if exists {
		return
	}

	self.mu.Lock()
	defer self.mu.Unlock()
	self.consumers = append(self.consumers, conn)
}

func (self *Topic) UnSubscribe(id string) {
	self.mu.Lock()
	defer self.mu.Unlock()

	i := slices.IndexFunc(self.consumers, func(c protocol.Connection) bool {
		return c.ID() == id
	})

	if i == -1 {
		return
	}

	self.consumers = append(self.consumers[:i], self.consumers[i+1:]...)

	if self.index >= len(self.consumers)-1 {
		self.index = -1
	}
}

func (self *Topic) Publish(packet *protocol.Publish) {
	self.queue.Push(packet)
}

func (self *Topic) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"messages":  len(self.queue),
		"consumers": self.Count(),
	})
}

func (self *Topic) listen() {
	for {
		select {
		case packet := <-self.queue:
			next, exists := self.Next()

			if !exists {
				self.Publish(packet)
				time.Sleep(100 * time.Millisecond)
				break
			}

			if err := next.Write(packet); err != nil {
				next.Close()
				return
			}
		default:
			time.Sleep(100 * time.Millisecond)
		}
	}
}
