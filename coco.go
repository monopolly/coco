package coco

import (
	"bytes"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/monopolly/coco/internal/cc"
)

// min is queue batch add 50, 1000,5000, etc
func New(count int) (a *Engine) {
	a = new(Engine)
	a.min = 1000

	a.db = cc.NewFilter(uint(count))
	go a.deamon()
	return
}

// min is queue batch add 50, 1000,5000, etc
func NewFromFile(filename string) (a *Engine, err error) {
	a = new(Engine)
	a.min = 1000
	err = a.Load(filename)
	if err != nil {
		return
	}
	go a.deamon()
	return
}

type Engine struct {
	db        *cc.Filter
	mu        sync.RWMutex
	lowercase bool
	min       int
	queue     [][]byte
}

func (a *Engine) Load(filename string) (err error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return
	}

	c, err := cc.Decode(data)
	if err != nil {
		return
	}

	a.db = c
	return
}

func (a *Engine) LoadData(data []byte) (err error) {
	c, err := cc.Decode(data)
	if err != nil {
		return
	}

	a.db = c
	return
}

func (a *Engine) Add(item []byte) {
	a.queue = append(a.queue, item)
}

func (a *Engine) Adds(item ...[]byte) {
	a.queue = append(a.queue, item...)
}

func (a *Engine) AddString(item string) {
	a.queue = append(a.queue, []byte(item))
}

func (a *Engine) AddStrings(item ...string) {
	for _, x := range item {
		a.queue = append(a.queue, []byte(x))
	}
}

func (a *Engine) Has(item []byte) (res bool) {
	if a.lowercase {
		item = bytes.ToLower(item)
	}
	a.mu.RLock()
	res = a.db.Lookup(item)
	a.mu.RUnlock()
	return
}

func (a *Engine) Hass(item string) (res bool) {
	if a.lowercase {
		item = strings.ToLower(item)
	}
	a.mu.RLock()
	res = a.db.Lookup([]byte(item))
	a.mu.RUnlock()
	return
}

func (a *Engine) Count() (res int) {
	return int(a.db.Count())
}

func (a *Engine) Data() (res []byte) {
	a.Flush()
	a.mu.RLock()
	res = a.db.Encode()
	a.mu.RUnlock()
	return
}

func (a *Engine) Save(filename string) (err error) {
	a.Flush()
	b := a.Data()
	return os.WriteFile(filename, b, os.ModePerm)
}

// minimum items in queue, default 1000
func (a *Engine) SetMin(min int) {
	a.min = min
}

func (a *Engine) deamon() {
	for {
		time.Sleep(time.Minute)
		if len(a.queue) < a.min {
			continue
		}
		a.Flush()
	}
}

// push queue to db
func (a *Engine) Flush() {
	a.mu.Lock()
	for _, x := range a.queue {
		if a.lowercase {
			x = bytes.ToLower(x)
		}
		if a.db.Lookup(x) {
			continue
		}
		a.db.Insert(x)
	}
	a.queue = nil
	a.mu.Unlock()
}

// add and check in lowercase
func (a *Engine) Lowercase(v bool) {
	a.lowercase = v
}
