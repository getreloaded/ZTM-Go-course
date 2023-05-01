//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Book struct {
	name        string
	status      string
	status_time time.Time
}

type Member struct {
	name string
}

type Library struct {
	books   []Book
	members []Member
}

func checkin(book *Book) Book {
	book.status = "In Library"
	book.status_time = time.Now()
	return *book
}

func checkout(book *Book, member *Member) Book {
	book.status = "With" + member.name
	book.status_time = time.Now()
	return *book
}

func libraryPrinter(library *Library) {
	fmt.Println("Books:")
	for _, book := range library.books {
		fmt.Println(book.name, book.status, book.status_time.String())
	}
	fmt.Println("Members:")
	for i, member := range library.members {
		fmt.Println(i+1, member.name)
	}
}

func main() {

	booklist := make([]Book, 4)
	booklist[0] = Book{name: "Alice in Wonderland", status: "In Library", status_time: time.Now()}
	booklist[1] = Book{name: "Bob in Wonderland", status: "In Library", status_time: time.Now()}
	booklist[2] = Book{name: "Charlie in Wonderland", status: "In Library", status_time: time.Now()}
	booklist[3] = Book{name: "David in Wonderland", status: "In Library", status_time: time.Now()}

	memberlist := make([]Member, 3)
	memberlist[0].name = "Scooby"
	memberlist[1].name = "Daffy"
	memberlist[2].name = "Bugs Bunny"

	library := Library{books: booklist, members: memberlist}

	libraryPrinter(&library)

	checkout(&booklist[0], &memberlist[0])

	libraryPrinter(&library)

	checkin(&booklist[0])

	libraryPrinter(&library)
}
