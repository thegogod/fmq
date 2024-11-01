package storage

import (
	"encoding/json"
	"sync"

	"github.com/thegogod/fmq/common/protocol"
)

type Topics struct {
	mu    sync.RWMutex
	items map[string]*Topic
}

func New() *Topics {
	return &Topics{
		items: map[string]*Topic{},
	}
}

func (self *Topics) Count() int {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return len(self.items)
}

func (self *Topics) Keys() []string {
	self.mu.RLock()
	defer self.mu.RUnlock()

	keys := make([]string, len(self.items))
	i := 0

	for key := range self.items {
		keys[i] = key
		i++
	}

	return keys
}

func (self *Topics) Get(key string) (*Topic, bool) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	topic, exists := self.items[key]
	return topic, exists
}

func (self *Topics) Subscribe(key string, conn protocol.Connection) {
	self.mu.RLock()
	topic, exists := self.items[key]
	self.mu.RUnlock()

	if !exists {
		topic = NewTopic()
		self.mu.Lock()
		self.items[key] = topic
		self.mu.Unlock()
	}

	topic.Subscribe(conn)
}

func (self *Topics) UnSubscribe(key string, id string) {
	self.mu.RLock()
	topic, exists := self.items[key]
	self.mu.RUnlock()

	if !exists {
		return
	}

	topic.UnSubscribe(id)
}

func (self *Topics) Publish(key string, packet *protocol.Publish) {
	self.mu.RLock()
	topic, exists := self.items[key]
	self.mu.RUnlock()

	if !exists {
		topic = NewTopic()
		self.mu.Lock()
		self.items[key] = topic
		self.mu.Unlock()
	}

	topic.Publish(packet)
}

func (self *Topics) MarshalJSON() ([]byte, error) {
	self.mu.RLock()
	defer self.mu.RUnlock()
	return json.Marshal(self.items)
}
