package main

import (
	"Homework4/API"
	DBPackage "Homework4/DB"
)

func main() {
	DBPackage.InitializeDB()
	API.InitializeRouter()

}
