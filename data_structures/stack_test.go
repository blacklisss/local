package data_structures

import "testing"

func TestStack(t *testing.T) {
	s := &Stack{}

	// Проверяем, что стек пуст
	if !s.IsEmpty() {
		t.Errorf("IsEmpty: expected true, got false")
	}

	// Проверяем, что Peek возвращает ошибку, когда стек пуст
	_, err := s.Peek()
	if err == nil {
		t.Errorf("Peek: expected error, got nil")
	}

	// Проверяем, что Pop возвращает ошибку, когда стек пуст
	_, err = s.Pop()
	if err == nil {
		t.Errorf("Pop: expected error, got nil")
	}

	// Добавляем элементы в стек
	s.Push(1)
	s.Push(2)
	s.Push(3)

	// Проверяем, что стек не пуст и верхний элемент равен 3
	if s.IsEmpty() {
		t.Errorf("IsEmpty: expected false, got true")
	}
	top, err := s.Peek()
	if err != nil {
		t.Errorf("Peek: expected nil, got %v", err)
	}
	if top != 3 {
		t.Errorf("Peek: expected 3, got %v", top)
	}

	// Удаляем элементы из стека
	popped, err := s.Pop()
	if err != nil {
		t.Errorf("Pop: expected nil, got %v", err)
	}
	if popped != 3 {
		t.Errorf("Pop: expected 3, got %v", popped)
	}

	popped, err = s.Pop()
	if err != nil {
		t.Errorf("Pop: expected nil, got %v", err)
	}
	if popped != 2 {
		t.Errorf("Pop: expected 2, got %v", popped)
	}

	popped, err = s.Pop()
	if err != nil {
		t.Errorf("Pop: expected nil, got %v", err)
	}
	if popped != 1 {
		t.Errorf("Pop: expected 1, got %v", popped)
	}

	// Проверяем, что стек пуст
	if !s.IsEmpty() {
		t.Errorf("IsEmpty: expected true, got false")
	}
}
