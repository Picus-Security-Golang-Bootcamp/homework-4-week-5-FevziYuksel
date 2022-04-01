package Models

import (
	"errors"
	"strings"

	"gorm.io/gorm"
)

type BookRepository struct {
	//Router *mux.Router
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{db: db}
}
func (b *BookRepository) Migrations() {
	err := b.db.AutoMigrate(&Book{})
	if err != nil {
		return
	}

}
func (b *BookRepository) Create(book Book) error {
	result := b.db.Create(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BookRepository) InsertSampleData(books []Book) {

	for _, book := range books {
		newBook := Book{
			Name:      book.GetBookName(),
			Pages:     book.GetBookPages(),
			Stocks:    book.GetBookStocks(),
			Price:     book.GetBookPrice(),
			StockCode: book.GetBookStockCode(),
			ISBN:      book.GetBookISBN(),
			AuthorID:  book.GetBookAuthorID(),
			Author:    book.GetBookAuthor(),
		}
		b.db.Where(Book{Name: newBook.GetBookName()}).FirstOrCreate(&newBook)
	}
}
func (b *BookRepository) FindAllBooks() []Book {
	var books []Book
	b.db.Find(&books)
	return books
}

func (b *BookRepository) GetBookByID(id int) (*Book, error) {
	var book Book
	result := b.db.First(&book, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &book, nil
}

func (b *BookRepository) FindBookByName(name string) []Book {
	var book []Book
	//name = strings.ToLower(name)
	b.db.Where("name LIKE ? ", "%"+name+"%").Find(&book)

	return book
}

func (b *BookRepository) FindByNameWithRawSQL(name string) []Book {
	var books []Book
	name = strings.ToLower(name)
	b.db.Raw("SELECT * FROM books WHERE name LIKE ?", "%"+name+"%").Scan(&books)

	return books
}

func (b *BookRepository) DeleteBook(book Book) error {
	result := b.db.Delete(book)

	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (b *BookRepository) DeleteBookById(id uint64) error {
	result := b.db.Delete(&Book{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b *BookRepository) GetBooksWithAuthorInformation() ([]Book, error) {
	var books []Book
	result := b.db.Preload("Author").Find(&books)
	if result.Error != nil {
		return nil, result.Error
	}
	return books, nil
}

func (b *BookRepository) InsertSingleBook(newBook *Book) {

	b.Create(*newBook)
}

func (b *BookRepository) FindBooksWithLimitOffset(limit int, offset int) []Book {
	var books []Book
	b.db.Limit(limit).Offset(offset).Find(&books)
	return books
}
func (b *BookRepository) FindBookByItsStockCode(requestCode string) Book {
	var book Book
	b.db.Where("stock_code = ?", requestCode).First(&book)
	return book
}
func (b *BookRepository) BuyBookByItsId(id uint, quantity uint64) {
	var book Book
	b.db.Model(&book).Where("id = ?", id).
		Where("stocks > ?", quantity).
		UpdateColumn("stocks", gorm.Expr("stocks - ?", quantity))
}
