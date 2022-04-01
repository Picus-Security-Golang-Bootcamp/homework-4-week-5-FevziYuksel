package API

import (
	DBPackage "Homework4/DB"
	"Homework4/Models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func getAllBooks(w http.ResponseWriter, r *http.Request) {
	books := DBPackage.BookRepo.FindAllBooks()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBookByID(w http.ResponseWriter, r *http.Request) {

	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	// error checking
	var find *Models.Book
	find, err = DBPackage.BookRepo.GetBookByID(id)

	if err != nil || id <= 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*find)
}
func getBookByName(w http.ResponseWriter, r *http.Request) {

	nameFind := mux.Vars(r)["name"]
	find := DBPackage.BookRepo.FindBookByName(nameFind)

	if len(find) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified name"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(find)
}
func deleteBookByID(w http.ResponseWriter, r *http.Request) {

	// get the ID of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		// there was an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	// error checking
	if err != nil || id <= 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}
	// Delete the post from the slice an DB

	_ = DBPackage.BookRepo.DeleteBookById(uint64(id))

	w.WriteHeader(200)

}
func deleteBookByName(w http.ResponseWriter, r *http.Request) {

	nameFind := mux.Vars(r)["name"]
	find := DBPackage.BookRepo.FindBookByName(nameFind)
	if len(find) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified name"))
		return
	}
	// Delete the post from the slice an DB
	for _, book := range find {
		_ = DBPackage.BookRepo.DeleteBookById(uint64(book.ID))
	}

	w.WriteHeader(200)
}
func addBook(w http.ResponseWriter, r *http.Request) {
	// get Item value from the JSON body
	var newBook Models.Book

	json.NewDecoder(r.Body).Decode(&newBook)

	DBPackage.BookRepo.InsertSingleBook(&newBook)

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(newBook)
}
func buyBook(w http.ResponseWriter, r *http.Request) {

	var amount string = mux.Vars(r)["quantity"]
	quantity, err := strconv.Atoi(amount)
	if err != nil || quantity <= 0 {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	fmt.Println(quantity)
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}
	fmt.Println(id)
	//Change quantity of the book
	DBPackage.BookRepo.BuyBookByItsId(uint(id), uint64(quantity))

	updatedBook, _ := DBPackage.BookRepo.GetBookByID(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedBook)
}
