package data_structures

import (
	"bytes"
	"io"
	"os"
	"sync"
	"testing"
)

func TestLinkedList(t *testing.T) {
	// Создаем список и добавляем несколько элементов
	list := &LinkedList{}
	list.Append(1)
	list.Append(2)
	list.Append(3)

	// Проверяем, что размер списка равен 3
	if list.size != 3 {
		t.Errorf("Expected size 3, but got %d", list.size)
	}

	// Удаляем элемент, который есть в списке
	removed := list.Remove(2)
	if !removed {
		t.Errorf("Expected Remove(2) to return true")
	}
	// Проверяем, что элемент действительно удален
	if list.size != 2 {
		t.Errorf("Expected size 2, but got %d", list.size)
	}
	expected := "1\n3\n"
	result := captureOutput(func() {
		list.Print()
	})
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}

	// Удаляем элемент, которого нет в списке
	removed = list.Remove(5)
	if removed {
		t.Errorf("Expected Remove(5) to return false")
	}
	// Проверяем, что размер списка не изменился
	if list.size != 2 {
		t.Errorf("Expected size 2, but got %d", list.size)
	}
	expected = "1\n3\n"
	result = captureOutput(func() {
		list.Print()
	})
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}

	// Удаляем головной элемент списка
	removed = list.Remove(1)
	if !removed {
		t.Errorf("Expected Remove(1) to return true")
	}
	// Проверяем, что головной элемент действительно удален
	if list.size != 1 {
		t.Errorf("Expected size 1, but got %d", list.size)
	}
	expected = "3\n"
	result = captureOutput(func() {
		list.Print()
	})
	if result != expected {
		t.Errorf("Expected %q, but got %q", expected, result)
	}

	// Удаляем элемент, который есть в списке
	removed = list.Remove(3)
	if !removed {
		t.Errorf("Expected Remove(3) to return true")
	}
	// Проверяем, что элемент действительно удален
	if list.size != 0 {
		t.Errorf("Expected size 0, but got %d", list.size)
	}

	// Удаляем элемент в пустом списке
	removed = list.Remove(2)
	if removed {
		t.Errorf("Expected false, but got true")
	}
}

func captureOutput(f func()) string {
	reader, writer, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	stdout := os.Stdout
	stderr := os.Stderr
	defer func() {
		os.Stdout = stdout
		os.Stderr = stderr
	}()
	os.Stdout = writer
	os.Stderr = writer

	out := make(chan string)
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		var buf bytes.Buffer
		wg.Done()
		io.Copy(&buf, reader)
		out <- buf.String()
	}()
	wg.Wait()
	f()
	_ = writer.Close()
	return <-out
}
