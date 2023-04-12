package data_structures

import (
	"errors"
	"sync"
)

// Stack представляет стек, реализованный в виде среза целых чисел.
type Stack struct {
	items []int
	lock  sync.RWMutex
}

// Push добавляет элемент в стек.
func (s *Stack) Push(item int) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.items = append(s.items, item)
}

// Pop удаляет и возвращает последний элемент из стека.
func (s *Stack) Pop() (int, error) {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	lastIndex := len(s.items) - 1
	lastItem := s.items[lastIndex]
	s.items = s.items[:lastIndex]
	return lastItem, nil
}

// IsEmpty возвращает true, если стек пуст, и false в противном случае.
func (s *Stack) IsEmpty() bool {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return len(s.items) == 0
}

// Peek возвращает последний элемент из стека, но не удаляет его.
func (s *Stack) Peek() (int, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}

	return s.items[len(s.items)-1], nil
}
