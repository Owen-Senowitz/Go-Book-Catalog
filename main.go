package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var bookList []Book

	file, ferr := os.Open("books.csv")
	if ferr != nil {
		panic(ferr)
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, ",")
		bookId, err := strconv.ParseInt(split[0], 10, 64)
		title := split[1]
		isbn, err := strconv.ParseInt(split[2], 10, 64)
		author := split[3]
		genre := split[4]
		bookList = append(bookList, Book{int(bookId), title, int(isbn), author, genre})
		if err != nil {
			panic(err)
		}
	}
	printBookList(bookList)
	var input int
	for true {
		fmt.Println("")
		fmt.Println("1: View Book List")
		fmt.Println("2: Add Book")
		fmt.Println("3: Remove Book")
		fmt.Println("4: Update Book")
		fmt.Println("5: Save Books to File")
		fmt.Println("6: Exit")
		fmt.Println("")
		fmt.Print("Input: ")
		fmt.Scan(&input)
		switch input {
		//View Book List
		case 1:
			printBookList(bookList)
			input = 0
			break
		//Add Book
		case 2:
			var bookId int
			var title string
			var isbn int
			var author string
			var genre string
			fmt.Println("Enter Book ID:")
			fmt.Scan(&bookId)
			fmt.Println("Enter Title:")
			fmt.Scan(&title)
			fmt.Println("Enter ISBN:")
			fmt.Scan(&isbn)
			fmt.Println("Enter Author:")
			fmt.Scan(&author)
			fmt.Println("Enter Genre:")
			fmt.Scan(&genre)
			bookList = addBookToList(bookList, Book{bookId, title, isbn, author, genre})
			input = 0
			break
		//Remove Book
		case 3:
			var bookId int
			fmt.Println("Enter BookID to Delete:")
			fmt.Scan(&bookId)
			bookList = removeBookFromList(bookList, bookId)
			input = 0
			break
		//Update Book
		case 4:
			var bookId int
			var newBookId int
			var title string
			var isbn int
			var author string
			var genre string
			fmt.Println("Enter The Book ID of the Book Updating:")
			fmt.Scan(&bookId)
			fmt.Println("Enter New Book ID:")
			fmt.Scan(&newBookId)
			fmt.Println("Enter New Title:")
			fmt.Scan(&title)
			fmt.Println("Enter New ISBN:")
			fmt.Scan(&isbn)
			fmt.Println("Enter New Author:")
			fmt.Scan(&author)
			fmt.Println("Enter New Genre:")
			fmt.Scan(&genre)
			bookList = updateBookFromList(bookList, bookId, Book{newBookId, title, isbn, author, genre})
			input = 0
			break
		//Save Books to File
		case 5:
			saveBooks(bookList)
			input = 0
			break
		//Exit
		case 6:
			os.Exit(0)
		}

	}

}
