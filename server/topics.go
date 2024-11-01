package main

import (
	"sync"

	"github.com/thegogod/fmq/common/protocol"
)

type Topics struct {
	mu    sync.RWMutex
	items map[string]*Topic
}

func newTopics() *Topics {
	return &Topics{
		items: map[string]*Topic{},
	}
}

func (self *Topics) Count() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return len(self.items)
}

func (self *Topics) Subscribe(key string, conn protocol.Connection) {
	self.mu.Lock()
	defer self.mu.Unlock()
	topic, exists := self.items[key]

	if !exists {
		topic = newTopic()
		self.items[key] = topic
	}

	topic.Subscribe(conn)
}

func (self *Topics) UnSubscribe(key string, id string) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	topic, exists := self.items[key]

	if !exists {
		return
	}

	topic.UnSubscribe(id)
}

func (self *Topics) Publish(key string, packet *protocol.Publish) {
	self.mu.Lock()
	defer self.mu.Unlock()
	topic, exists := self.items[key]

	if !exists {
		topic = newTopic()
		self.items[key] = topic
	}

	topic.Publish(packet)
}
