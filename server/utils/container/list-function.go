package container

import "container/list"

func Find(collection *list.List, fn func(value interface{}, index int64) bool) *list.Element {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		found := fn(e.Value, index)
		if found {
			return e
		}
		index++
	}
	return nil
}

func FindByElement(collection *list.List, fn func(element *list.Element, index int64) bool) *list.Element {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		found := fn(e, index)
		if found {
			return e
		}
		index++
	}
	return nil
}

func FindAfter(last *list.Element, fn func(value interface{}) bool) *list.Element {
	for e := last.Next(); e != nil; e = e.Next() {
		found := fn(e.Value)
		if found {
			return e
		}
	}
	return nil
}

func Map(collection *list.List, fn func(value interface{}, index int64) interface{}) *list.List {
	newList := list.New()
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		v := fn(e.Value, index)
		newList.PushBack(v)
		index++
	}
	return newList
}

func Filter(collection *list.List, fn func(value interface{}, index int64) bool) *list.List {
	newList := list.New()
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		ok := fn(e.Value, index)
		if ok {
			newList.PushBack(e.Value)
		}
		index++
	}
	return newList
}

func ForEach(collection *list.List, fn func(value interface{}, index int64)) {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		fn(e.Value, index)
		index++
	}
}

func ForEachElement(collection *list.List, fn func(element *list.Element, index int64)) {
	index := int64(0)
	for e := collection.Front(); e != nil; e = e.Next() {
		fn(e, index)
		index++
	}
}
