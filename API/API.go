package API

import (
	"github.com/gorilla/mux"
	"net/http"
)

func InitializeRouter() {

	// Initializing Router and Handle functions
	router := mux.NewRouter()

	router.HandleFunc("/books", getAllBooks).Methods("GET")

	router.HandleFunc("/authors", getAllAuthor).Methods("GET")

	router.HandleFunc("/bookID/{id}", getBookByID).Methods("GET")

	router.HandleFunc("/authorID/{id}", getAuthorByID).Methods("GET")

	router.HandleFunc("/bookName/{name}", getBookByName).Methods("GET")

	router.HandleFunc("/authorName/{name}", getAuthorByName).Methods("GET")

	router.HandleFunc("/bookID/{id}", deleteBookByID).Methods("DELETE")

	router.HandleFunc("/bookName/{name}", deleteBookByName).Methods("DELETE")

	router.HandleFunc("/book", addBook).Methods("POST")

	router.HandleFunc("/buy/{id}/{quantity}", buyBook).Methods("PUT")

	//router.HandleFunc("/posts/{id}", patchBook).Methods("PATCH")

	http.ListenAndServe(":5000", router)

}
