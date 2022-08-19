package main

import (
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

type Book struct {
	bookId int
	title  string
	isbn   int
	author string
	genre  string
}

func printBookList(bookList []Book) {
	writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, '\t', tabwriter.AlignRight)
	_, err := fmt.Fprintln(writer, "BookID\tTitle\tISBN\tAuthor\tGenre")
	if err != nil {
		return
	}
	for i, _ := range bookList {
		_, err := fmt.Fprintln(writer, bookList[i].bookId, "\t", bookList[i].title, "\t", bookList[i].isbn, "\t", bookList[i].author, "\t", bookList[i].genre)
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}

func addBookToList(bookList []Book, book Book) []Book {
	return append(bookList, book)
}

func removeBookFromList(bookList []Book, bookId int) []Book {
	for i, _ := range bookList {
		if bookList[i].bookId == bookId {
			return append(bookList[:i], bookList[i+1:]...)
		}
	}
	return bookList
}

func updateBookFromList(bookList []Book, updateBookId int, book Book) []Book {
	bookList = removeBookFromList(bookList, updateBookId)
	return addBookToList(bookList, book)
}

func saveBooks(bookList []Book) {
	file, err := os.Create("books.csv")
	if err != nil {
		panic(err)
	}
	for i, _ := range bookList {
		_, err := file.WriteString(strconv.Itoa(bookList[i].bookId) + "," + bookList[i].title + "," + strconv.Itoa(bookList[i].isbn) + "," + bookList[i].author + "," + bookList[i].genre + "\n")
		if err != nil {
			return
		}
	}
}
