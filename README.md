## Homework | Week 5

* This application is created as a RESTFUL API server with PostgreSQL Database.
* Application excepts user inputs through client and execute these commands.

``` 
  router.HandleFunc("", getAllBooks).Methods("GET")

	router.HandleFunc("/authors", getAllAuthor).Methods("GET")

	router.HandleFunc("/bookID/{id}", getBookByID).Methods("GET")

	router.HandleFunc("/authorID/{id}", getAuthorByID).Methods("GET")

	router.HandleFunc("/bookName/{name}", getBookByName).Methods("GET")

	router.HandleFunc("/authorName/{name}", getAuthorByName).Methods("GET")

	router.HandleFunc("/bookID/{id}", deleteBookByID).Methods("DELETE")

	router.HandleFunc("/bookName/{name}", deleteBookByName).Methods("DELETE")

	router.HandleFunc("/book", addBook).Methods("POST")

	router.HandleFunc("/buy/{id}/{quantity}", buyBook).Methods("PUT")
``` 
