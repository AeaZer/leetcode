package string_array

type LRUCache struct {
	capacity int
	cacheMap map[int]int
	keys     []int
}

func Constructor2(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		cacheMap: make(map[int]int, capacity),
		keys:     make([]int, 0, capacity),
	}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.cacheMap[key]
	if ok {
		upKeyPriority(key, c.keys)
		return v
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	_, ok := c.cacheMap[key]
	if ok {
		c.cacheMap[key] = value
		upKeyPriority(key, c.keys)
		return
	}

	if len(c.cacheMap) == c.capacity {
		delete(c.cacheMap, c.keys[0])
		c.keys = c.keys[1:]
	}
	c.cacheMap[key] = value
	c.keys = append(c.keys, key)
}

func upKeyPriority(key int, keys []int) {
	var index int
	for i, v := range keys {
		if v == key {
			index = i
			break
		}
	}
	for i := index; i < len(keys)-1; i++ {
		keys[i] = keys[i+1]
	}
	keys[len(keys)-1] = key
}
