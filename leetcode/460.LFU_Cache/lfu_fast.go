package lfucache

type cacheValue struct {
	key   int
	value int
	freq  int
	next  *cacheValue
	prev  *cacheValue
}

type LFUCachef struct {
	cap         int
	minf        int
	frequencies map[int]*linkedListf
	cache       map[int]*cacheValue
}

type linkedListf struct {
	head *cacheValue
	tail *cacheValue
}

func Constructorf(cap int) *LFUCachef {
	return &LFUCachef{
		cap:         cap,
		frequencies: make(map[int]*linkedListf),
		cache:       make(map[int]*cacheValue),
	}
}

func (c *LFUCachef) Get(key int) int {
	tmp, ok := c.cache[key]
	if !ok {
		return -1
	}

	c.frequencies[tmp.freq].remove(tmp)
	if c.frequencies[tmp.freq].head == nil {
		delete(c.frequencies, tmp.freq)
		if tmp.freq == c.minf {
			c.minf++
		}
	}

	tmp.freq++
	if _, ok := c.frequencies[tmp.freq]; !ok {
		c.frequencies[tmp.freq] = &linkedListf{}
	}

	c.frequencies[tmp.freq].insert(tmp)

	return tmp.value
}

func (c *LFUCachef) Put(key, value int) {
	if c.cap == 0 {
		return
	}

	if val := c.Get(key); val != -1 {
		c.cache[key].value = value
		return
	}

	if len(c.cache) == c.cap {
		delete(c.cache, c.frequencies[c.minf].head.key)

		c.frequencies[c.minf].remove(c.frequencies[c.minf].head)
		if c.frequencies[c.minf].head == nil {
			delete(c.frequencies, c.minf)
		}
	}

	c.minf = 1

	cv := &cacheValue{
		key:   key,
		value: value,
		freq:  1,
	}

	c.cache[key] = cv

	if _, ok := c.frequencies[cv.freq]; !ok {
		c.frequencies[cv.freq] = &linkedListf{}
	}

	c.frequencies[cv.freq].insert(cv)
}

func (l *linkedListf) remove(cv *cacheValue) {
	if cv.prev != nil {
		cv.prev.next = cv.next
	} else {
		l.head = cv.next
	}

	if cv.next != nil {
		cv.next.prev = cv.prev
	} else {
		l.tail = cv.prev
	}

	cv.next, cv.prev = nil, nil
}

func (l *linkedListf) insert(cv *cacheValue) {
	if l.head == nil {
		l.head = cv
	} else {
		l.tail.next = cv
		cv.prev = l.tail
	}
	l.tail = cv
}
