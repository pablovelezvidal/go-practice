package data

import (
	"fmt"
	"sync"
)

type Book struct {
	Id       int
	Name     string
	Finished bool
}

var books = []*Book{
	{1, "Dune", false},
	{2, "El perfume", false},
	{3, "The world of ice and fire", false},
	{4, "Teoria de la noche", false},
	{5, "100 años de soledad", false},
	{6, "Metamorfosis", false},
	{7, "Entrevista con el vampiro", false},
	{8, "El alquimista", false},
	{9, "Maze runner", false},
	{10, "El principito", false},
}

func FindBook(id int, m *sync.RWMutex) (int, *Book) {
	index := -1
	var book *Book

	m.RLock()
	for i, b := range books {
		if b.Id == id {
			index = i
			book = b
		}
	}
	m.RUnlock()
	return index, book
}

func FinishBook(id int, m *sync.RWMutex) {
	i, book := FindBook(id, m)
	if i < 0 {
		return
	}
	m.Lock()
	book.Finished = true
	books[i] = book
	m.Unlock()
	fmt.Printf("Book %s finished \n", book.Name)
}
