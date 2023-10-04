package designhashmap

type Node struct {
	key, value int
	Next       *Node
}

type MyHashMap struct {
	store []*Node
	size  int
}

func Constructor() MyHashMap {
	return MyHashMap{
		store: make([]*Node, 10000),
		size:  10000,
	}
}

func (m *MyHashMap) Put(key int, value int) {
	idx := key % m.size
	if m.store[idx] == nil {
		m.store[idx] = &Node{
			key:   key,
			value: value,
		}
		return
	}

	prev, curr := m.store[idx], m.store[idx]
	for curr != nil {
		if curr.key == key {
			curr.value = value
			return
		}
		prev = curr
		curr = curr.Next
	}
	prev.Next = &Node{
		key:   key,
		value: value,
	}
}

func (m *MyHashMap) Get(key int) int {
	curr := m.store[key%m.size]
	for curr != nil {
		if curr.key == key {
			return curr.value
		}
		curr = curr.Next
	}

	return -1
}

func (m *MyHashMap) Remove(key int) {
	idx := key % m.size
	if m.store[idx] == nil {
		return
	}

	if m.store[idx].key == key {
		m.store[idx] = m.store[idx].Next
	}

	prev, curr := m.store[idx], m.store[idx]
	for curr != nil {
		if curr.key == key {
			prev.Next = curr.Next
			return
		}
		prev = curr
		curr = curr.Next
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
