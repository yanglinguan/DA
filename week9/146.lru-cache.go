type LRUCache struct {
	kv map[int]*list.Element
	data *list.List
	capacity int
	size int
}


func Constructor(capacity int) LRUCache {
	lru := LRUCache {
			kv: make(map[int]*list.Element),
			data: list.New(),
			capacity: capacity,
			size: 0,
	}
	return lru
}


func (this *LRUCache) Get(key int) int {
	if _, exist := this.kv[key]; !exist {
			return -1
	}
	e := this.kv[key]
	v := e.Value.([]int)[1]
	this.data.MoveToFront(e)
	return v
}


func (this *LRUCache) Put(key int, value int)  {
	if _, exist := this.kv[key]; exist {
			e := this.kv[key]
			e.Value.([]int)[1] = value
			this.data.MoveToFront(e)
			return
	}
	
	if this.capacity == this.size {
			last := this.data.Back()
			this.data.Remove(last)
			delete(this.kv, last.Value.([]int)[0])
			this.size--
	}
	e := this.data.PushFront([]int{key, value})
	this.kv[key] = e
	this.size++
}


/**
* Your LRUCache object will be instantiated and called as such:
* obj := Constructor(capacity);
* param_1 := obj.Get(key);
* obj.Put(key,value);
*/