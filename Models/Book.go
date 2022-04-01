package Models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Name      string `json:"Name"`
	Pages     uint64 `json:"Pages"`
	Stocks    uint64 `json:"Stocks"`
	Price     uint64 `json:"Price"`
	StockCode string `json:"StockCode"`
	ISBN      string `json:"ISBN"`
	AuthorID  uint64 `json:"AuthorID"`
	Author    Author `gorm:"foreignKey:AuthorID"  json:"Author"`
}

type BookSlice struct {
	Books []Book `json:"Books"`
}

//Getters
func (b *Book) GetBookID() uint {
	return b.ID
}
func (b *Book) GetBookName() string {
	return b.Name
}
func (b *Book) GetBookPages() uint64 {
	return b.Pages
}
func (b *Book) GetBookStocks() uint64 {
	return b.Stocks
}
func (b *Book) GetBookPrice() uint64 {
	return b.Price
}
func (b *Book) GetBookStockCode() string {
	return b.StockCode
}
func (b *Book) GetBookISBN() string {
	return b.ISBN
}
func (b *Book) GetBookAuthorID() uint64 {
	return b.AuthorID
}
func (b *Book) GetBookAuthor() Author {
	return b.Author
}

//Setters
func (b *Book) SetBookID(id uint) {
	b.ID = id
}
func (b *Book) SetBookName(name string) {
	b.Name = name
}
func (b *Book) SetBookAuthor(author Author) {
	b.Author = author
}
func (b *Book) SetBookStock(stocks uint64) {
	b.Stocks = stocks
}
func (b *Book) SetBookPrice(price uint64) {
	b.Price = price
}
func (b *Book) SetBookAuthorID(authorID uint64) {
	b.AuthorID = authorID
}
