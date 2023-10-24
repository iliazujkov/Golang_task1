package mp

import (
	"fmt"
	"sort"
	"sync"
)


type Map struct {
	mp        map[int64]interface{}
	nextIndex int64
	mutexAdd  sync.RWMutex
}

func (m *Map) NewStorage() interface{} {
	return &Map{
		mp:        make(map[int64]interface{}),
		nextIndex: 0,
	}
}

func (m *Map) Len() (l int64) {
	return int64(len(m.mp))
}

func (m *Map) Add(value interface{}) (id int64) {
	defer mutexAdd.Unlock()
	mutexAdd.Lock()
	if m.mp == nil {
		m.mp = make(map[int64]interface{})
	}
	m.mp[m.nextIndex] = value
	m.nextIndex++
	return m.nextIndex - 1
}

// RemoveByIndex удаляет элемент из map по индексу
func (m *Map) RemoveByIndex(id int64) {
	defer mutexAdd.Unlock()
	mutexAdd.Lock()
	delete(m.mp, id)
}

// RemoveByValue удаляет элемент из map по значению
func (m *Map) RemoveByValue(value interface{}) {
	defer mutexAdd.Unlock()
	mutexAdd.Lock()
	for k, v := range m.mp {
		if v == value {
			delete(m.mp, k)
			return
		}
	}
}

// RemoveAllByValue удаляет все элементы из map по значению
func (m *Map) RemoveAllByValue(value interface{}) {
	defer mutexAdd.Unlock()
	mutexAdd.Lock()
	for k, v := range m.mp {
		if v == value {
			delete(m.mp, k)
		}
	}
}

// GetByIndex возвращает значение элемента по индексу.
//
// Если элемента с таким индексом нет, то возвращается 0 и false.
func (m *Map) GetByIndex(id int64) (interface{}, bool) {
	defer mutexAdd.RUnlock()
	mutexAdd.RLock()
	value, ok := m.mp[id]
	return value, ok
}

// GetByValue возвращает индекс первого найденного элемента по значению.
//
// Если элемента с таким значением нет, то возвращается 0 и false.
func (m *Map) GetByValue(value interface{}) (int64, bool) {
	defer mutexAdd.RUnlock()
	mutexAdd.RLock()
	for k, v := range m.mp {
		if v == value {
			return k, true
		}
	}
	return 0, false
}

// GetAllByValue возвращает индексы всех найденных элементов по значению
//
// Если элементов с таким значением нет, то возвращается nil и false.
func (m *Map) GetAllByValue(value interface{}) (ids []int64, ok bool) {
	defer mutexAdd.RUnlock()
	mutexAdd.RLock()
	for k, v := range m.mp {
		if v == value {
			ids = append(ids, k)
		}
	}
	if ids != nil {
		return ids, true
	}
	return nil, false
}

// GetAll возвращает все элементы списка
//
// Если map пуст, то возвращается nil и false.
func (m *Map) GetAll() (values []interface{}, ok bool) {
	defer mutexAdd.RUnlock()
	mutexAdd.RLock()
	keys := make([]int64, 0, len(m.mp))
	for k := range m.mp {
		keys = append(keys, k)
	}
	for _, k := range keys {
		values = append(values, m.mp[k])
	}
	if values != nil {
		return values, true
	}
	return nil, false
}

// Clear очищает map
func (m *Map) Clear() {
	m.mp = make(map[int64]interface{})
	m.nextIndex = 0
}

// Print выводит map в консоль
func (m *Map) Print() {
	defer mutexAdd.Unlock()
	mutexAdd.Lock()
	for k, v := range m.mp {
		fmt.Println("id: ", k, "value: ", v)
	}
}
