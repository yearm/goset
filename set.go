package structure

import "sync"

//线程安全set
var Empty = struct{}{}

type Set struct {
	m map[interface{}]struct{}
	sync.RWMutex
}

func New(items ...interface{}) *Set {
	s := &Set{}
	s.m = make(map[interface{}]struct{})
	s.Add(items...)
	return s
}

func (this *Set) Add(items ...interface{}) {
	for _, item := range items {
		this.Lock()
		this.m[item] = Empty
		this.Unlock()
	}
}

func (this *Set) Remove(item interface{}) {
	this.Lock()
	delete(this.m, item)
	this.Unlock()
}

func (this *Set) Size() int {
	this.RLock()
	length := len(this.m)
	this.RUnlock()
	return length
}

func (this *Set) Contains(item interface{}) bool {
	this.RLock()
	_, ok := this.m[item]
	this.RUnlock()
	return ok
}

func (this *Set) Clear() {
	this.Lock()
	this.m = make(map[interface{}]struct{})
	this.Unlock()
}

func (this *Set) List(callback func(items interface{})) {
	this.RLock()
	for k, _ := range this.m {
		callback(k)
	}
	this.RUnlock()
}
