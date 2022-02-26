package library

import "strings"

var books = []string{
	"In Search of Lost Time",
	"Ulysses",
	"Don Quixote",
	"One Hundred Years of Solitude",
	"The Great Gatsby",
}

func List() []string {
	return books
}

func Search(booksToSearch ...string) []string {
	foundBooks := make([]string, 0, len(booksToSearch))

	for _, book1 := range books {
		for _, book2 := range booksToSearch {
			if strings.Contains(strings.ToLower(book1), strings.ToLower(book2)) {
				foundBooks = append(foundBooks, book1)
			}
		}
	}

	return foundBooks
}
