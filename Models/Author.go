package Models

import (
	"gorm.io/gorm"
)

type Author struct {
	gorm.Model
	AuthorName string `json:"AuthorName"`
	Book       []Book
}
type AuthorSlice []Author

//Getters
func (a *Author) GetAuthorName() string {
	return a.AuthorName
}
func (a *Author) GetAuthorID() uint {
	return a.ID
}
func (a *Author) GetAuthorBook() []Book {
	return a.Book
}

//Setters
func (a *Author) SetAuthorName(name string) {
	a.AuthorName = name
}
func (a *Author) SetAuthorID(id uint) {
	a.ID = id
}
func (a *Author) SetAuthorBook(book []Book) {
	a.Book = book
}
