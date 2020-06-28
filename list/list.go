package bidirectionallist

import "github.com/serg1122/bidirectionallist/node"

type List struct {
	first *node.Node
	last  *node.Node
}

func New() *List {

	return &List{}
}

func (l *List) IsEmpty() bool {

	return l.first == nil
}

func (l *List) init(value interface{}) {

	n := node.New(nil, nil, value)

	l.first = n
	l.last = n
}

func (l *List) Prepend(value interface{}) {

	if l.IsEmpty() {
		l.init(value)
		return
	}

	newFirstNode := node.New(nil, l.first, value)
	l.first.SetPrev(newFirstNode)
	l.first = newFirstNode
}

func (l *List) Append(value interface{}) {

	if l.IsEmpty() {
		l.init(value)
		return
	}

	newLastNode := node.New(l.last, nil, value)
	l.last.SetLast(newLastNode)
	l.last = newLastNode
}

func (l List) GetFirstNode() *node.Node {

	return l.first
}

func (l List) GetLastNode() *node.Node {

	return l.last
}
