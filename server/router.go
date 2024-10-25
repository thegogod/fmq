package main

import (
	"regexp"
	"slices"
	"sync"

	"github.com/thegogod/fmq/common/protocol"
)

type Router struct {
	mu    sync.RWMutex
	items map[string][]protocol.Connection
	exprs map[string]*regexp.Regexp
}

func NewRouter() *Router {
	return &Router{
		items: map[string][]protocol.Connection{},
		exprs: map[string]*regexp.Regexp{},
	}
}

func (self *Router) On(pattern string, conn protocol.Connection) *Router {
	self.mu.Lock()
	defer self.mu.Unlock()

	subs, exists := self.items[pattern]

	if !exists {
		subs = []protocol.Connection{}
	}

	if slices.ContainsFunc(subs, func(c protocol.Connection) bool {
		return c.ID() == conn.ID()
	}) {
		return self
	}

	subs = append(subs, conn)
	self.items[pattern] = subs
	self.exprs[pattern] = regexp.MustCompile(pattern)

	return self
}

func (self *Router) Push(packet *protocol.Publish) {
	self.mu.RLock()
	defer self.mu.RUnlock()

	for pattern, expr := range self.exprs {
		if expr.MatchString(packet.Topic) {
			subs, exists := self.items[pattern]

			if !exists {
				subs = []protocol.Connection{}
			}

			for _, conn := range subs {
				conn.Write(packet)
			}

			self.items[pattern] = subs
		}
	}
}
