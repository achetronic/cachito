/*
Copyright 2025.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package lru

import (
	"container/list"
	"errors"
	"sync"
)

// Entry represents a key-value pair stored in the cache.
// It is passed to the user-defined handlers.
type Entry struct {
	Key   string
	Value any
}

// LRU implements a thread-safe LRU cache with support for
// user-defined handlers and custom metadata.
type LRU[MetaT any] struct {
	mu       sync.RWMutex
	index    map[string]*list.Element
	list     *list.List
	Metadata MetaT // User-defined metadata available in all handlers

	// User-defined hooks
	onInsertHandler    func(metadata *MetaT, entry Entry) error
	onDeleteHandler    func(metadata *MetaT, entry Entry) error
	onAccessHandler    func(metadata *MetaT, entry Entry) error
	shouldEvictHandler func(metadata *MetaT, entry Entry) bool
}

// New creates a new LRU structure. The `metadata` object can be any value,
// and is accessible in all handler functions.
func New[MetaT any](metadata MetaT) *LRU[MetaT] {
	return &LRU[MetaT]{
		index:    make(map[string]*list.Element),
		list:     list.New(),
		Metadata: metadata,
	}
}

// OnInsert sets a handler to be called when a new entry is created
func (c *LRU[MetaT]) OnInsert(handler func(metadata *MetaT, entry Entry) error) {
	c.onInsertHandler = handler
}

// OnDelete sets a handler to be called when an entry is removed from the cache.
func (c *LRU[MetaT]) OnDelete(handler func(metadata *MetaT, entry Entry) error) {
	c.onDeleteHandler = handler
}

// OnAccess sets a handler to be called when an entry is accessed.
func (c *LRU[MetaT]) OnAccess(handler func(metadata *MetaT, entry Entry) error) {
	c.onAccessHandler = handler
}

// ShouldEvict sets a handler that decides whether eviction should occur.
// It should return true if the cache should evict the least recently used entry.
func (c *LRU[MetaT]) ShouldEvict(handler func(metadata *MetaT, entry Entry) bool) {
	c.shouldEvictHandler = handler
}

// CreateElement inserts or updates an entry in the cache.
// If eviction is needed, the least recently used entries are removed
// before the new one is inserted.
// Eviction conditions are managed by the user defining OnEvict
func (c *LRU[MetaT]) CreateElement(key string, value any) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	entry := Entry{Key: key, Value: value}
	element, exists := c.index[key]

	if exists {
		// Update existing element
		element.Value = entry
	} else {
		// Run eviction loop before inserting new element
		for c.shouldEvictHandler != nil && c.shouldEvictHandler(&c.Metadata, entry) {
			if err := c.deleteLastElementUnsafe(); err != nil {
				return err
			}
		}
		// Insert new element at the front
		element = c.list.PushFront(entry)
		c.index[key] = element
	}

	// Run create handler if present
	if c.onInsertHandler != nil {
		return c.onInsertHandler(&c.Metadata, entry)
	}
	return nil
}

// GetElement returns the value associated with the given key and
// moves it to the front (most recently used).
func (c *LRU[MetaT]) GetElement(key string) (any, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	element, found := c.index[key]
	if !found {
		return nil, nil
	}

	// Move to front (recent use)
	c.list.MoveToFront(element)
	entry := element.Value.(Entry)

	// Run get handler if present
	if c.onAccessHandler != nil {
		if err := c.onAccessHandler(&c.Metadata, entry); err != nil {
			return nil, err
		}
	}
	return entry.Value, nil
}

// DeleteElement removes an entry by key from the LRU.
func (c *LRU[MetaT]) DeleteElement(key string) error {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.deleteElementUnsafe(key)
}

// deleteElementUnsafe removes an entry by key from the LRU without locking it
func (c *LRU[MetaT]) deleteElementUnsafe(key string) error {

	element, found := c.index[key]
	if !found {
		return nil
	}

	entry := element.Value.(Entry)

	// Run delete handler if present
	if c.onDeleteHandler != nil {
		if err := c.onDeleteHandler(&c.Metadata, entry); err != nil {
			return err
		}
	}

	// Remove from map and list
	delete(c.index, key)
	c.list.Remove(element)
	return nil
}

// deleteLastElement removes the least recently used element from the LRU without locking it
func (c *LRU[MetaT]) deleteLastElementUnsafe() error {
	element := c.list.Back()
	if element == nil {
		return errors.New("cannot evict: cache is empty")
	}
	entry := element.Value.(Entry)
	return c.deleteElementUnsafe(entry.Key)
}
