package main

import (
	"fmt"
	"mypackage/library"
	"os"
	"path"
	"strings"
)

func main() {
	switch args := os.Args; {
	case len(args) == 1:
		projectName := path.Base(args[0])
		fmt.Printf("Available commands for %s: \n search => search books \n list => list all books\n", projectName)
	case len(args) == 2 && args[1] == "list":
		booksInTheBookShelf := library.List()
		if len(booksInTheBookShelf) > 0 {
			fmt.Printf("Books in the bookshelf: %s.\n", strings.Join(booksInTheBookShelf, ", "))
		} else {
			fmt.Println("No books in the bookshelf!")
		}
	case len(args) == 2 && args[1] == "search":
		fmt.Println("Please enter name of the books you would like to search.")
	case len(args) > 2 && args[1] == "search":
		booksToSearch := args[2:]
		foundBooks := library.Search(booksToSearch...)
		if len(foundBooks) > 0 {
			fmt.Printf("Found books in the bookshelf: %s.\n", strings.Join(foundBooks, ", "))
		} else {
			fmt.Println("No books found!")
		}
	default:
		fmt.Println("Invalid arguments.")
	}
}
