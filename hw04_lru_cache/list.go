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
	Value      interface{}
	Next, Prev *ListItem
}

type list struct {
	// List
	len         int
	front, back *ListItem
}

func NewList() List {
	return new(list)
}

func (l *list) Len() int         { return l.len }
func (l *list) Front() *ListItem { return l.front }
func (l *list) Back() *ListItem  { return l.back }

func (l *list) PushFront(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
	}
	if l.front == nil && l.back == nil {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Next = l.front
		l.front.Prev = newItem
		l.front = newItem
	}
	l.len++
	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	newItem := &ListItem{
		Value: v,
	}
	if l.back == nil && l.front == nil {
		l.front = newItem
		l.back = newItem
	} else {
		newItem.Prev = l.back
		l.back.Next = newItem
		l.back = newItem
	}
	l.len++
	return l.back
}

func (l *list) Remove(i *ListItem) {
	if i.Prev != i.Next {
		prevItem := i.Prev
		nextItem := i.Next

		switch {
		case l.front == i:
			l.front = nextItem
			nextItem.Prev = nil
		case l.back == i:
			l.back = prevItem
			prevItem.Next = nil
		default:
			prevItem.Next = nextItem
			nextItem.Prev = prevItem
		}
		// i = nil
		l.len--
	}
}

func (l *list) MoveToFront(i *ListItem) {
	// l.Remove(i)
	// l.PushFront(i.Value)
	if l.front != i {
		prevItem := i.Prev
		nextItem := i.Next
		currentFront := l.front

		if l.back == i {
			l.back = prevItem
			l.back.Next = nil
			if currentFront == prevItem {
				l.back.Prev = i
			}
		} else {
			prevItem.Next = nextItem.Prev
			nextItem.Prev = prevItem.Next
		}

		i.Prev = nil
		i.Next = currentFront
		l.front = i
		currentFront.Prev = l.front
	}
}
