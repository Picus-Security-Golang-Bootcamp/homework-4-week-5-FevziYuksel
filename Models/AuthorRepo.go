package Models

import (
	"errors"
	"gorm.io/gorm"
)

type AuthorRepository struct {
	db *gorm.DB
}

func NewAuthorRepository(db *gorm.DB) *AuthorRepository {
	return &AuthorRepository{db: db}
}
func (a *AuthorRepository) Migration() {
	err := a.db.AutoMigrate(&Author{})
	if err != nil {
		return
	}
}

func (a *AuthorRepository) InsertSampleData(authors []Author) {
	var authorList []Author
	for _, author := range authors {
		newAuthors := Author{
			AuthorName: author.GetAuthorName(),
		}
		authorList = append(authorList, newAuthors)
	}
	for _, eachAuthor := range authorList {
		a.db.Create(&eachAuthor)

	}
}
func (a *AuthorRepository) GetAllAuthorsWithBookInformation() ([]Author, error) {
	var authors []Author
	result := a.db.Preload("Book").Find(&authors)
	if result.Error != nil {
		return []Author{}, result.Error
	}
	return authors, nil
}

func (a *AuthorRepository) GetAuthorByNameWithBookInformation(name string) []Author {
	var authors []Author
	a.db.Where("author_name LIKE ? ", "%"+name+"%").
		Find(&authors)
	return authors
}

func (a *AuthorRepository) GetAuthorByIdWithBookInformation(id int) (*Author, error) {
	var author Author
	result := a.db.First(&author, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, result.Error
	}
	return &author, nil
}

func (a *AuthorRepository) GetAuthorByIdWithBookInformationRawSQL(id int) []Author {
	var authors []Author
	a.db.Raw("select * from books b right join authors a on b.author_id = a.id where b.author_id = ?", id).Scan(&authors)
	return authors
}
