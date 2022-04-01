package DBPackage

import (
	"Homework4/Models"
	"gorm.io/gorm"
)

var (
	DB         *gorm.DB
	BookSlice1 Models.BookSlice
	Books      []Models.Book
	Authors    []Models.Author
	BookRepo   *Models.BookRepository
	AuthorRepo *Models.AuthorRepository
)

func InitializeDB() {
	DB = InitialMigration()

	//JSON to struct methods
	BookSlice1.ReadJSON("BookList.json")
	Books = BookSlice1.ConvertBook()
	Authors = BookSlice1.ExtractAuthor()

	AuthorRepo = Models.NewAuthorRepository(DB)
	AuthorRepo.Migration()

	BookRepo = Models.NewBookRepository(DB)
	BookRepo.Migrations()
	BookRepo.InsertSampleData(Books)
}
