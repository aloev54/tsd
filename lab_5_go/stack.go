package stack

import (
	"log"
)

type Stack struct {
	size     int
	capacity int
	elements []string
}

// Проверка стека на пустоту
func (s Stack) isEmpty() bool {
	return s.size == 0
}

// Проверка стека на полноту
func (s Stack) isFull() bool {
	return s.size == s.capacity
}

// Добавить эллемент в стек
func (s Stack) push(element string) {
	if s.isFull() {
		log.Fatal("Stack is full")
	}
	s.elements = append(s.elements, element)
	s.size++
}

// Удалить эллемент из стека
func (s Stack) pop() (item []string) {
	if s.isEmpty() {
		log.Fatal("Stack is empty")
	}
	s.elements = s.elements[:len(s.elements)-1]
	s.size--
	return s.elements
}

// Достать последний эллемент из стека
func (s Stack) peek() string {
	if s.isEmpty() {
		log.Fatal("Stack is empty")
	}
	return s.elements[s.size-1]
}
