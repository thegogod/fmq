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
	listeners []protocol.Connection
}

func NewTopic() *Topic {
	self := &Topic{
		index:     -1,
		queue:     make(Queue[*protocol.Publish], 10000),
		listeners: []protocol.Connection{},
	}

	go self.listen()
	return self
}

func (self *Topic) Count() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return len(self.listeners)
}

func (self *Topic) Next() (protocol.Connection, bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	if len(self.listeners) == 0 {
		return nil, false
	}

	i := self.index + 1

	if i > (len(self.listeners) - 1) {
		i = 0
	}

	self.index = i
	return self.listeners[i], true
}

func (self *Topic) Subscribe(conn protocol.Connection) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	exists := slices.ContainsFunc(self.listeners, func(c protocol.Connection) bool {
		return c.ID() == conn.ID()
	})

	if exists {
		return
	}

	self.listeners = append(self.listeners, conn)
}

func (self *Topic) UnSubscribe(id string) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	i := slices.IndexFunc(self.listeners, func(c protocol.Connection) bool {
		return c.ID() == id
	})

	if i == -1 {
		return
	}

	self.listeners = append(self.listeners[:i], self.listeners[i+1:]...)

	if self.index >= len(self.listeners)-1 {
		self.index = -1
	}
}

func (self *Topic) Publish(packet *protocol.Publish) {
	self.queue.Push(packet)
}

func (self *Topic) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"messages":  len(self.queue),
		"listeners": self.Count(),
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
