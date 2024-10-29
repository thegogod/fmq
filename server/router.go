package main

import (
	"regexp"
	"slices"
	"sync"
	"time"

	"math/rand"

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
	self.exprs[pattern] = nil

	if expr, err := regexp.Compile(pattern); err == nil {
		self.exprs[pattern] = expr
	}

	return self
}

func (self *Router) Next(topic string) protocol.Connection {
	self.mu.RLock()
	defer self.mu.RUnlock()

	for pattern, expr := range self.exprs {
		if expr == nil && pattern != topic {
			continue
		}

		if expr != nil && !expr.MatchString(topic) {
			continue
		}

		subs, exists := self.items[pattern]

		if !exists {
			subs = []protocol.Connection{}
		}

		if len(subs) > 0 {
			r := rand.New(rand.NewSource(time.Now().Unix()))
			return subs[r.Intn(len(subs))]
		}

		self.items[pattern] = subs
	}

	return nil
}
