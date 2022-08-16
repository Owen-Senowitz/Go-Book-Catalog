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
	for i, _ := range bookList {
		fmt.Println(bookList[i])
	}
	var input int
	for true {
		fmt.Println("1: View Book List")
		fmt.Println("2: Add Book")
		fmt.Println("3: Remove Book")
		fmt.Println("4: Update Book")
		fmt.Println("5: Save Books to File")
		fmt.Println("6: Exit")
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
			input = 0
			break
		//Remove Book
		case 3:
			input = 0
			break
		//Update Book
		case 4:
			input = 0
			break
		//Save Books to File
		case 5:
			input = 0
			break
		//Exit
		case 6:
			os.Exit(0)
		}

	}

}
