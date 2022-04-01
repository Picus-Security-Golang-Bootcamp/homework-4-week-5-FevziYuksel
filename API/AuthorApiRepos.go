package API

import (
	DBPackage "Homework4/DB"
	"Homework4/Models"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

//
func getAllAuthor(w http.ResponseWriter, r *http.Request) {
	authors, _ := DBPackage.AuthorRepo.GetAllAuthorsWithBookInformation()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(authors)
}
func getAuthorByID(w http.ResponseWriter, r *http.Request) {

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
	var find *Models.Author
	find, err = DBPackage.AuthorRepo.GetAuthorByIdWithBookInformation(id)

	if err != nil || id <= 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified ID"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(*find)
}
func getAuthorByName(w http.ResponseWriter, r *http.Request) {

	nameFind := mux.Vars(r)["name"]
	find := DBPackage.AuthorRepo.GetAuthorByNameWithBookInformation(nameFind)

	if len(find) == 0 {
		w.WriteHeader(404)
		w.Write([]byte("No book found with specified name"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(find)
}
