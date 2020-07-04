package list

import (
	"github.com/serg1122/bidirectional"
)

type List struct {
	first *bidirectional.Node
	last  *bidirectional.Node
}

func New() *List {

	return &List{}
}

func (l *List) IsEmpty() bool {

	return l.first == nil
}

func (l *List) init(value interface{}) {

	n := bidirectional.CreateNode(nil, nil, value)

	l.first = n
	l.last = n
}

func (l *List) Prepend(value interface{}) {

	if l.IsEmpty() {
		l.init(value)
		return
	}

	newFirstNode := bidirectional.CreateNode(nil, l.first, value)
	l.first.SetPrev(newFirstNode)
	l.first = newFirstNode
}

func (l *List) Append(value interface{}) {

	if l.IsEmpty() {
		l.init(value)
		return
	}

	newLastNode := bidirectional.CreateNode(l.last, nil, value)
	l.last.SetLast(newLastNode)
	l.last = newLastNode
}

func (l List) GetFirstNode() *bidirectional.Node {

	return l.first
}

func (l List) GetLastNode() *bidirectional.Node {

	return l.last
}
