package orderedmap

import (
	"github.com/serg1122/bidirectional"
	"github.com/serg1122/bidirectional/list"
)

type OrderedMap struct {
	list  *list.List
	pairs map[string]*bidirectional.Node
}

func New() *OrderedMap {

	return &OrderedMap{
		list:  list.New(),
		pairs: map[string]*bidirectional.Node{},
	}
}

func (om OrderedMap) Length() int {

	return len(om.pairs)
}

func (om OrderedMap) IsEmpty() bool {

	return om.Length() == 0
}

func (om OrderedMap) HasKey(key string) bool {

	_, hasKey := om.pairs[key]
	return hasKey
}

func (om *OrderedMap) Set(key string, value interface{}) {

	if om.HasKey(key) {

		om.pairs[key].SetValue(value)
		return
	}

	om.list.Append(value)
	om.pairs[key] = om.list.GetLastNode()
}

func (om OrderedMap) Get(key string) (interface{}, bool) {

	if om.HasKey(key) {

		return om.pairs[key], true
	}

	return nil, false
}

func (om *OrderedMap) Remove(key string) bool {

	if !om.HasKey(key) {

		return false
	}

	node := om.pairs[key]

	if om.Length() == 1 {
		om.list = list.New() // no way to fix om.list.First() and om.list.Last()
	} else {

		node := om.pairs[key]

		prevNode := node.GetPrev()
		nextNode := node.GetNext()

		if prevNode != nil {
			prevNode.SetNext(nextNode)
		}

		if nextNode != nil {
			nextNode.SetPrev(prevNode)
		}
	}

	node.SetPrev(nil) // Paranoia
	node.SetNext(nil) // Paranoia

	delete(om.pairs, key)

	return true
}

func (om OrderedMap) GetList() *list.List {

	return om.list
}
