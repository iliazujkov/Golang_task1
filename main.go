package main

import (
	"MachineLearning_1/storage/list"
	"MachineLearning_1/storage/mp"
	"fmt"
)

type Storage interface {
	Len() int64
	Add(value int64) int64
	RemoveByIndex(id int64)
	RemoveByValue(value int64)
	RemoveAllByValue(value int64)
	GetByIndex(id int64) (int64, bool)
	GetByValue(value int64) (int64, bool)
	GetAllByValue(value int64) (ids []int64, ok bool)
	GetAll() (values []int64, ok bool)
	Clear()
	Print()
}

func main() {
	l := list.NewList()
	l.Add(12)
	l.Add(54)
	l.Add(3)
	l.Add(8)
	l.Add(12)
	l.Add(14)
	l.Add(428)
	l.Add(2123)
	fmt.Println(l.GetAllByValue(14))
	fmt.Println(l.GetAll())
	fmt.Println(l.GetByValue(12))
	fmt.Println(l.GetByIndex(-4))
	l.RemoveAllByValue(12)
	l.RemoveByIndex(4)
	l.RemoveByValue(4)
	l.Print()
	m := mp.NewMap()
	m.Add(5)
	m.Add(2)
	m.Add(1)
	m.Add(5)
	m.Add(1)
	m.Add(4)
	//m.RemoveByIndex(3)
	//fmt.Println(m.GetAllByValue(5))
	//fmt.Println(m.GetByValue(5))
	fmt.Println(m.GetAll())
	m.Print()
	m.Clear()
	m.Print()
}
