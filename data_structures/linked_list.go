package data_structures

import "fmt"

type Node struct {
	data interface{} // данные узла
	next *Node       // ссылка на следующий узел
}

type LinkedList struct {
	head *Node // указатель на начальный узел списка
	size int   // количество элементов в списке
}

// Метод, который добавляет элемент в конец связного списка
func (l *LinkedList) Append(data interface{}) {
	newNode := &Node{data: data, next: nil}

	if l.head == nil {
		l.head = newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = newNode
	}
	l.size++
}

// Метод, который удаляет элемент из связного списка
func (l *LinkedList) Remove(value int) bool {
	if l.head == nil {
		return false
	}
	if l.head.data == value {
		l.head = l.head.next
		l.size--
		return true
	}
	prev := l.head
	curr := l.head.next
	for curr != nil {
		if curr.data == value {
			prev.next = curr.next
			l.size--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

// Метод, который выводит все элементы связного списка
func (l *LinkedList) Print() {
	current := l.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}
