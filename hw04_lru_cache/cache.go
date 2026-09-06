package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value any) bool
	Get(key Key) (any, bool)
	Clear()
}

type lruCache struct {
	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value any
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value any) bool {
	node, exists := c.items[key]

	if !exists {
		newItem := &cacheItem{key, value}
		newNode := c.queue.PushFront(newItem)
		c.items[key] = newNode

		if c.queue.Len() > c.capacity {
			oldNode := c.queue.Back()
			oldItem := oldNode.Value.(*cacheItem)

			delete(c.items, oldItem.key)
			c.queue.Remove(oldNode)
		}
	} else {
		item := node.Value.(*cacheItem)
		item.value = value

		c.queue.MoveToFront(node)
	}

	return exists
}

func (c *lruCache) Get(key Key) (any, bool) {
	node, exists := c.items[key]

	if !exists {
		return nil, false
	}

	c.queue.MoveToFront(node)

	return node.Value.(*cacheItem).value, true
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*ListItem, c.capacity)
}
