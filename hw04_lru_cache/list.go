package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v any) *ListItem
	PushBack(v any) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value any
	Next  *ListItem
	Prev  *ListItem
}

func NewList() List {
	return &list{}
}

type list struct {
	head *ListItem
	tail *ListItem
	size int
}

func (l *list) Len() int {
	return l.size
}

func (l *list) Front() *ListItem {
	return l.head
}

func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v any) *ListItem {
	node := &ListItem{Value: v}

	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		node.Next = l.head
		l.head.Prev = node
		l.head = node
	}

	l.size++

	return node
}

func (l *list) PushBack(v any) *ListItem {
	node := &ListItem{Value: v}

	if l.size == 0 {
		l.head = node
		l.tail = node
	} else {
		node.Prev = l.tail
		l.tail.Next = node
		l.tail = node
	}

	l.size++

	return node
}

func (l *list) Remove(i *ListItem) {
	if i == nil {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	} else {
		l.head = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Next = nil
	i.Prev = nil
	l.size--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == nil || i == l.head {
		return
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.tail = i.Prev
	}

	i.Prev = nil
	i.Next = l.head
	l.head.Prev = i
	l.head = i
}
