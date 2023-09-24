package main

import (
	"MachineLearning_1/storage/list"
	"fmt"
)

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
}
