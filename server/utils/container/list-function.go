package container

import "container/list"

func Some(list *list.List, fn func(element *list.Element, index int64) bool) *list.Element {
	index := int64(0)
	for e := list.Front(); e != nil; e = e.Next() {
		found := fn(e, index)
		if found {
			return e
		}
		index++
	}
	return nil
}
