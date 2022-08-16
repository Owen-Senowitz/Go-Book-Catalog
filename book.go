package main

import (
	"fmt"
	"os"
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
	fmt.Fprintln(writer, "BookID\tTitle\tISBN\tAuthor\tGenre")
	for i, _ := range bookList {
		fmt.Fprintln(writer, bookList[i].bookId, "\t", bookList[i].title, "\t", bookList[i].isbn, "\t", bookList[i].author, "\t", bookList[i].genre)
	}
	writer.Flush()
}

func addBookList(bookList []Book, book Book) []Book {
	return append(bookList, book)
}

func removeBookFromList(bookList []Book, bookId int) []Book {
	for i, _ := range bookList {
		if bookList[i].bookId == bookId {
			fmt.Println(i)
			return append(bookList[:i], bookList[i+1:]...)
		}
	}
	return bookList
}
