package Models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func (b *BookSlice) ReadJSON(fileAdress string) {
	file, err := os.Open(fileAdress)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()
	jsonFile, _ := ioutil.ReadAll(file)
	err = json.Unmarshal([]byte(jsonFile), b)
	if err != nil {
		fmt.Println(err)
	}
}

func (b *BookSlice) ConvertBook() []Book {
	var books []Book
	books = append(books, b.Books...)
	return books
}

func (b *BookSlice) ExtractAuthor() []Author {
	var authors []Author
	for _, eachAuthor := range b.Books {
		authors = append(authors, eachAuthor.Author)
	}
	return authors
}
