package main

import (
	"storage/list"
	"storage/mp"
	"fmt"
)

type storage interface {
	NewStorage() interface{}
	Add(value interface{}) int64
	Len() int64
	RemoveByIndex(id int64)
	RemoveByValue(value interface{})
	RemoveAllByValue(value interface{})
	GetByIndex(id int64) (interface{}, bool)
	GetByValue(value interface{}) (int64, bool)
	GetAllByValue(value interface{}) (ids []int64, ok bool)
	GetAll() (values []interface{}, ok bool)
	Clear()
	Print()
}

func main() {
	var l storage = &list.List{}
	l.NewStorage()
	l.Add(12)
	l.Add(54)
	l.Add(3)
	l.Add(8)
	l.Add(12)
	l.Add(14)
	l.Add(428)
	l.Add(2123)
	fmt.Println(l.GetAllByValue(14))
	fmt.Println(l.GetAllByValue(13))
	fmt.Println(l.GetAllByValue(12))
	fmt.Println(l.GetAllByValue(11))

	fmt.Println(l.GetAll())
	fmt.Println(l.GetByValue(12))
	fmt.Println(l.GetByIndex(-4))
	l.RemoveAllByValue(12)
	l.RemoveByIndex(4)
	l.RemoveByValue(4)
	l.Print()
	var m storage = &mp.Map{}
	m.NewStorage()
	m.Add(5)
	m.Add(2)
	m.Add(1)
	m.Add(5)
	m.Add(1)
	m.Add(4)
	m.RemoveByIndex(3)
	m.RemoveByIndex(3)
	m.RemoveByIndex(0)
	fmt.Println(m.GetAllByValue(5))
	fmt.Println(m.GetByValue(5))
	go fmt.Println(m.GetByIndex(1))
	go fmt.Println(m.GetByIndex(1))
	go fmt.Println(m.GetByIndex(1))
	fmt.Println(m.GetAll())
	m.Print()
	m.Clear()
	m.Print()
}
