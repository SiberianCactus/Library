package authors

import (
	"fmt"
	"sync"
)

// ID
// CamelCase
// lowcaseParametr
// ErrorFunc > 2
// swagger
// postgress
// dbeaver

type Authors_books struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Book    string `json:"book"`
}

type Authors_books_Store struct {
	sync.RWMutex

	authors map[int]Authors_books
	nextId  int
}

func New() *Authors_books_Store {
	auth := &Authors_books_Store{}
	auth.authors = make(map[int]Authors_books)
	auth.nextId = 0
	return auth
}

func (auth *Authors_books_Store) CreateAuthor(Name, Surname, Book string) int {
	auth.Lock()
	defer auth.Unlock()

	author := Authors_books{
		Id:      auth.nextId,
		Name:    Name,
		Surname: Surname,
		Book:    Book}

	auth.authors[auth.nextId] = author
	auth.nextId++
	return author.Id
}

func (auth *Authors_books_Store) GetAuthor(id int) (Authors_books, error) {
	auth.RLock()
	defer auth.RUnlock()

	t, ok := auth.authors[id]
	if ok {
		return t, nil
	}
	return Authors_books{}, fmt.Errorf("author with id=%d not found", id)
}

func (auth *Authors_books_Store) DeleteAuthor(id int) error {
	auth.Lock()
	defer auth.Unlock()

	if _, ok := auth.authors[id]; !ok {
		return fmt.Errorf("author with id=%d not found", id)
	}

	delete(auth.authors, id)
	return nil
}

func (auth *Authors_books_Store) GetAllAuthors() []Authors_books {
	auth.RLock()
	defer auth.RUnlock()

	allAuthors := make([]Authors_books, 0, len(auth.authors))
	for _, task := range auth.authors {
		allAuthors = append(allAuthors, task)
	}
	return allAuthors
}
