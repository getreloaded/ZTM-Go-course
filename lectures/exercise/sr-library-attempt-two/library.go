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

//library structure will have book[bookcard] and members[membercard] fields

type Title string
type Name string

type Library struct {
	booklist   map[Title]Bookcard
	memberlist map[Name]Membercard
}

type Bookcard struct {
	QtyTotal uint
	QtyOut   uint
}

type Membercard struct {
	booksIssued   map[Title]time.Time
	booksReturned map[Title]time.Time
}

func (library *Library) checkout(book Title, member Name) bool {
	timeIssued, issued := library.memberlist[member].booksIssued[book]
	timeReturned, returned := library.memberlist[member].booksReturned[book]
	bookStatus := timeIssued.Compare(timeReturned)

	if issued && !returned {
		fmt.Println(member, "has not returned the book yet")
		return false
	}
	if bookStatus == -1 {
		fmt.Println(member, "has already reissued the book")
		return false
	}

	item, found := library.booklist[book]
	//	fmt.Println(item, found, membercard.booksIssued, membercard.checkInTime, membercard.checkOutTime)
	if !found {
		fmt.Println("Item is not available in the library at all.")
		return false
	} else {
		if item.QtyTotal-item.QtyOut == 0 {
			fmt.Println("All the copies of", book, "has already been issued")
			return false
		} else {
			library.memberlist[member] = Membercard{
				booksIssued:   map[Title]time.Time{book: time.Now()},
				booksReturned: map[Title]time.Time{},
			}
			item.QtyOut += 1
			library.booklist[book] = Bookcard{QtyTotal: item.QtyTotal, QtyOut: item.QtyOut}
			//fmt.Println(item, found, membercard.booksIssued, membercard.checkInTime, membercard.checkOutTime)
			return true
		}
	}
}

func (library *Library) checkin(book Title, member Name) bool {

	timeIssued, issued := library.memberlist[member].booksIssued[book]
	timeReturned := library.memberlist[member].booksReturned[book]
	bookStatus := timeIssued.Compare(timeReturned)

	if bookStatus == 1 {
		fmt.Println(member, "has already returned the book")
		return false
	}
	if !issued {
		fmt.Println(member, "has not issued the book yet")
		return false
	}

	item, found := library.booklist[book]
	if !found {
		fmt.Println("The library does not own", book)
		return false
	} else {
		if item.QtyOut == 0 {
			fmt.Println("All our copies of", book, "are already in the library.")
			return false
		} else {
			library.memberlist[member] = Membercard{
				booksIssued:   map[Title]time.Time{book: library.memberlist[member].booksIssued[book]},
				booksReturned: map[Title]time.Time{book: time.Now()},
			}
			item.QtyOut -= 1
			library.booklist[book] = Bookcard{QtyTotal: item.QtyTotal, QtyOut: item.QtyOut}
			//fmt.Println("book returned after", library.memberlist[member].booksReturned[book].Compare(library.memberlist[member].booksIssued[book]))
			return true
		}
	}
}

func (library *Library) libraryPrinter() {
	fmt.Println("-------START--------")
	for books, bookcards := range library.booklist {
		fmt.Println(books, ": Total", bookcards.QtyTotal, "of which", bookcards.QtyOut, "are out")
	}
	for name, membercard := range library.memberlist {
		for book, time := range membercard.booksIssued {
			fmt.Println(name, "has issued", book, "at the time", time.String())
		}
		for book, time := range membercard.booksReturned {
			fmt.Println(name, "has returned", book, "at the time", time.String())
		}
	}
	fmt.Println("--------END---------")

}

func main() {

	lb := Library{
		booklist:   make(map[Title]Bookcard),
		memberlist: make(map[Name]Membercard),
	}

	// lb.booklist["Eragon"] = Bookcard{4, 0}
	// lb.booklist["Eldest"] = Bookcard{2, 0}
	// lb.booklist["Brisingr"] = Bookcard{3, 0}
	// lb.booklist["Inheritance"] = Bookcard{5, 0}

	lb.booklist = map[Title]Bookcard{
		"Eragon":      {4, 0},
		"Eldest":      {2, 0},
		"Brisingr":    {3, 0},
		"Inheritance": {5, 0},
	}

	lb.memberlist = map[Name]Membercard{
		"Scooby":    {},
		"Daffy":     {},
		"BugsBunny": {},
	}

	lb.libraryPrinter()

	lb.checkout("Eragon", "Scooby")
	lb.checkin("Eragon", "Daffy")
	lb.checkout("Eragon", "BugsBunny")
	lb.checkout("Eldest", "BugsBunny")
	lb.checkout("Eldest", "Daffy")
	lb.checkout("Eldest", "Scooby")

	lb.libraryPrinter()
}
