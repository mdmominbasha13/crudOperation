package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/crud/pkg/models"
	"github.com/crud/pkg/utils"
	"github.com/gorilla/mux"
)

var newBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetBook()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, _ := models.GetBookById(id)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	createBook := &models.Book{}
	utils.ParseBody(r, createBook)
	b := createBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBookById(id)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookid := vars["bookid"]
	id, err := strconv.ParseInt(bookid, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDetails, db := models.GetBookById(id)
	if updateBook.Name != "" {
		bookDetails.Name = updateBook.Name

	}
	if updateBook.Author != "" {
		bookDetails.Author = updateBook.Author

	}
	if updateBook.Publication != "" {
		bookDetails.Publication = updateBook.Publication

	}
	db.Save(bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
