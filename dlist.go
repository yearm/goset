package structure

import (
	"errors"
)

//节点
type DNode struct {
	Data interface{}
	Prve *DNode
	Next *DNode
}

//链表
type DList struct {
	Size int64
	Head *DNode
	Tail *DNode
}

func NewDList() *DList {
	list := &DList{}
	list.Size = 0
	list.Head = nil
	list.Tail = nil
	return list
}

func (d *DList) Append(data interface{}) (err error) {
	if data == nil {
		err = errors.New("data cannot be empty！")
	}
	newNode := &DNode{Data: data}
	if d.Size == 0 {
		d.Head = newNode
		d.Tail = newNode
		newNode.Prve = nil
		newNode.Next = nil
	} else {
		newNode.Prve = d.Tail
		newNode.Next = nil
		d.Tail.Next = newNode //将之前的尾节点的Next指向新的节点
		d.Tail = newNode      //更新链表尾节点
	}
	d.Size++
	return
}

func (d *DList) Insert(index int64, data interface{}) (err error) {
	switch {
	case index > d.Size-1 || index < 0:
		err = errors.New("index error！")
	case data == nil:
		err = errors.New("data cannot be empty！")
	case d.Size == 0 || index == d.Size-1:
		d.Append(data)
	default:
		node, _ := d.Get(index)
		newNode := &DNode{Data: data}
		newNode.Next = node.Next
		newNode.Prve = node
		node.Next = newNode
		newNode.Next.Prve = newNode
		d.Size++
	}
	return
}

func (d *DList) Get(index int64) (node *DNode, err error) {
	switch {
	case index > d.Size-1 || index < 0:
		node = nil
		err = errors.New("index error！")
	case d.Size == 0:
		node = nil
	case index == 0:
		node = d.Head
	case index == d.Size-1:
		node = d.Tail
	default:
		node = d.Head
		for i := 0; int64(i) < index; i++ {
			node = node.Next
		}
	}
	return
}

func (d *DList) Remove(index int64) (err error) {
	node, err := d.Get(index)
	if err != nil {
		return
	}
	switch {
	case d.Size == 0:
		err = errors.New("this list is nil")
	case index > d.Size-1 || index < 0:
		err = errors.New("index error！")
	case index == 0:
		node.Next.Prve = nil
		d.Head = node.Next
	case index == d.Size-1:
		node.Prve.Next = nil
		d.Tail = node.Prve
	default:
		node.Prve.Next = node.Next
		node.Next.Prve = node.Prve
	}
	d.Size--
	return
}

func (d *DList) RemoveAll() {
	for i := 0; int64(i) < d.Size; i++ {
		node := d.Head
		d.Head = node.Next
		node.Next = nil
		node.Prve = nil
	}
	d.Tail = nil
	d.Size = 0
}

func (d *DList) Update(index int64, data interface{}) (err error) {
	switch {
	case d.Size == 0:
		err = errors.New("this list is nil！")
	case index > d.Size-1 || index < 0:
		err = errors.New("index error！")
	case data == nil:
		err = errors.New("data cannot be empty！")
	default:
		node := d.Head
		for i := 0; int64(i) <= index; i++ {
			if int64(i) == index {
				node.Data = data
			}
			node = node.Next
		}
	}
	return
}
