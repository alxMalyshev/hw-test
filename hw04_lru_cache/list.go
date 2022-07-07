package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	// List ListItem
	len int
	front *ListItem
	back *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int { return l.len }

func (l *list) Front() *ListItem {
	
	return
}
func (l *list) Back() *ListItem {return}

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
	}
	if l.front == nil {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
		l.front = newItem
	}
	l.len ++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {}
func (l *list) Remove(i *ListItem) {}
func (l *list) MoveToFront(i *ListItem) {}
