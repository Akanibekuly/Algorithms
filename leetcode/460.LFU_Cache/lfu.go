package lfucache

import "fmt"

type pair struct {
	value, frequency int
}

type linkedList struct {
	value int
	next  *linkedList
}

type LFUCache struct {
	cache          map[int]*pair
	frequencies    map[int]*linkedList
	capacity, minf int
}

func Constructor(capacity int) LFUCache {
	return LFUCache{
		cache:       make(map[int]*pair, capacity),
		frequencies: make(map[int]*linkedList),
		capacity:    capacity,
	}
}

func (c *LFUCache) insert(key, freq, value int) {

	if c.frequencies[freq] == nil {
		c.frequencies[freq] = &linkedList{value: key}
	} else {
		tail := c.frequencies[freq]
		for tail.next != nil {
			tail = tail.next
		}

		tail.next = &linkedList{value: key}
	}

	c.cache[key] = &pair{value: value, frequency: freq}
}

func (c *LFUCache) Get(key int) int {
	p, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.frequencies[p.frequency] = c.frequencies[p.frequency].remove(key)
	if p.frequency == c.minf && c.frequencies[p.frequency] == nil {
		c.minf++
	}

	c.insert(key, p.frequency+1, p.value)

	return p.value
}

func (c *LFUCache) Put(key int, value int) {
	if c.capacity <= 0 {
		return
	}

	if p, ok := c.cache[key]; ok {
		p.value = value

		c.Get(key)
	} else {
		if len(c.cache) == c.capacity {
			lfu := c.frequencies[c.minf]
			c.frequencies[c.minf] = c.frequencies[c.minf].remove(lfu.value)
			delete(c.cache, lfu.value)
		}

		c.insert(key, 1, value)
		c.minf = 1
	}
}

func (c *LFUCache) print() {
	if c == nil {
		fmt.Println("(empty)")
		return
	}

	fmt.Print("cahce: ")
	for key, p := range c.cache {
		fmt.Printf("{key: %d, val: %d, freq: %d}, ", key, p.value, p.frequency)
	}
	fmt.Println()

	fmt.Println("frequencies")
	for freq, l := range c.frequencies {
		fmt.Printf("freq: %d, list: ", freq)
		l.print()
	}
}

func (l *linkedList) remove(key int) *linkedList {
	if l == nil {
		return nil
	}

	if l.value == key {
		l = l.next
	} else {
		for t := l; t.next != nil; t = t.next {
			if t.next.value == key {
				t.next = t.next.next
				break
			}
		}
	}

	return l
}

func (l *linkedList) print() {
	if l == nil {
		fmt.Println("(empty)")
		return
	}

	for t := l; t != nil; t = t.next {
		fmt.Print(t.value)
		if t.next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
