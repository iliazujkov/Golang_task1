package list

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type List struct {
	len       int64
	firstNode *node
}

var mutexAdd sync.Mutex

// NewList создает новый список
func (l *List) NewStorage() interface{} {
	return &List{len: 0, firstNode: nil}
}

// Len возвращает длину списка
func (l *List) Len() (len int64) {
	return l.len
}

// Add добавляет элемент в список и возвращает его индекс
func (l *List) Add(value interface{}) (id int64) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	newNode := &node{value: value}
	if l.firstNode == nil {
		l.firstNode = newNode
		l.len++
		l.firstNode.index = l.len - 1
		return l.firstNode.index
	}
	if reflect.TypeOf(l.firstNode.value) != reflect.TypeOf(value) {
		var ErrMismatchType = errors.New("mismatched type: the type of the provided value does not match the type of items already in the storage")
		fmt.Println(ErrMismatchType)
		return 0
	}
	next_node := l.firstNode
	for ; next_node.next != nil; next_node = next_node.next {
	}
	next_node.next = newNode
	l.len++
	next_node.next.index = l.len - 1
	return next_node.next.index
}

// RemoveByIndex удаляет элемент из списка по индексу
func (l *List) RemoveByIndex(id int64) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	if l.firstNode == nil {
		return
	}
	if l.firstNode.index == id {
		l.firstNode = l.firstNode.next
		for n := l.firstNode; n != nil; n = n.next {
			n.index--
		}
		l.len--
		return
	}
	n := l.firstNode
	for n.next != nil && n.next.index != id {
		n = n.next
	}
	if n.next != nil {
		n.next = n.next.next
		l.len--
	}
	for it := l.firstNode; it != nil; it = it.next {
		it.index--
	}
}

// RemoveByValue удаляет элемент из списка по значению
func (l *List) RemoveByValue(value interface{}) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	if l.firstNode == nil {
		return
	}
	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		for n := l.firstNode; n != nil; n = n.next {
			n.index--
		}
		l.len--
		return
	}
	n := l.firstNode
	for n.next != nil && n.next.value != value {
		n = n.next
	}
	if n.next != nil {
		n.next = n.next.next
	}
	for it := l.firstNode; it != nil; it = it.next {
		it.index--
	}
}

// RemoveAllByValue удаляет все элементы из списка по значению
func (l *List) RemoveAllByValue(value interface{}) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	if l.firstNode == nil {
		return
	}
	if l.firstNode.value == value {
		l.firstNode = l.firstNode.next
		for n := l.firstNode; n != nil; n = n.next {
			n.index--
		}
		l.len--
	}
	n := l.firstNode
	for n.next != nil {
		if n.next.value == value {
			n.next = n.next.next
		}
		n = n.next
	}
	for it := l.firstNode; it != nil; it = it.next {
		it.index--
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (l *List) GetByIndex(id int64) (value interface{}, ok bool) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	for n := l.firstNode; n != nil; n = n.next {
		if n.index == id {

			return n.value, true
		}
	}
	return 0, false
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (l *List) GetByValue(value interface{}) (id int64, ok bool) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	for n := l.firstNode; n != nil; n = n.next {
		if n.value == value {

			return n.index, true
		}
	}
	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (l *List) GetAllByValue(value interface{}) (ids []int64, ok bool) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	var nums []int64
	for n := l.firstNode; n != nil; n = n.next {
		if n.value == value {
			nums = append(nums, n.index)
		}
	}
	if nums != nil {

		return nums, true
	}
	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если список пуст, то возвращается nil и false.
func (l *List) GetAll() (values []interface{}, ok bool) {
	mutexAdd.Lock()
	defer mutexAdd.Unlock()
	var nums []interface{}
	for n := l.firstNode; n != nil; n = n.next {
		nums = append(nums, n.value)
	}
	if nums != nil {
		return nums, true
	}
	return nil, false
}

// Clear очищает список
func (l *List) Clear() {
	l.firstNode = nil
	l.len = 0
}

// Print выводит список в консоль
func (l *List) Print() {
	n := l.firstNode
	if l.firstNode == nil {
		fmt.Println("Пусто")
		return
	}
	for ; n != nil; n = n.next {
		fmt.Println(n.value)
	}
}
