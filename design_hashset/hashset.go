package main

type MyHashSet struct {
	set map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{
		set: make(map[int]struct{}),
	}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.set, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.set[key]
	return ok
}
