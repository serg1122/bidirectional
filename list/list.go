package list

import (
	"sync"

	"github.com/serg1122/bidirectional"
)

type List struct {
	first *bidirectional.Node
	last  *bidirectional.Node

	mu   sync.Mutex
	once sync.Once
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

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsEmpty() {
		l.once.Do(func() { l.init(value) })
		return
	}

	newFirstNode := bidirectional.CreateNode(nil, l.first, value)
	l.first.SetPrev(newFirstNode)
	l.first = newFirstNode
}

func (l *List) Append(value interface{}) {

	l.mu.Lock()
	defer l.mu.Unlock()

	if l.IsEmpty() {
		l.once.Do(func() { l.init(value) })
		return
	}

	newLastNode := bidirectional.CreateNode(l.last, nil, value)
	l.last.SetNext(newLastNode)
	l.last = newLastNode
}

func (l List) GetFirstNode() *bidirectional.Node {

	return l.first
}

func (l List) GetLastNode() *bidirectional.Node {

	return l.last
}
