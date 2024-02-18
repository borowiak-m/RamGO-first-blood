package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}

type BookEntry struct {
	total  int // total books owned by library
	lended int // total books lended out
}

type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"
		} else {
			returnTime = audit.checkIn.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)
	}
}

func printLibraryAudit(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}

func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/total:", book.total, "/ lended:", book.lended)
	}
	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title] // check if book exists in the library
	if !found {
		fmt.Println("Book not part of library")
		return false
	}
	if book.lended == book.total { // check if books available to lend
		fmt.Println("No more books available to lend")
		return false
	}

	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}

func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title] // check if book exists in the library
	if !found {
		fmt.Println("Book not part of library")
		return false
	}

	audit, checkedOut := member.books[title]
	if !checkedOut {
		fmt.Println("Member did not check out this book")
		return false
	}

	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main_lib() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["Book1"] = BookEntry{total: 4, lended: 0}
	library.books["Book2"] = BookEntry{total: 1, lended: 0}
	library.books["Book3"] = BookEntry{total: 2, lended: 0}

	library.members["Member1"] = Member{"Member 1 Name", make(map[Title]LendAudit)}
	library.members["Member2"] = Member{"Member 2 Name", make(map[Title]LendAudit)}
	library.members["Member3"] = Member{"Member 3 Name", make(map[Title]LendAudit)}

	fmt.Println("\nInitial: ") // iniatial lib info - nothing lended, nothing returned
	printLibraryBooks(&library)
	printLibraryAudit(&library)

	member := library.members["Member2"]
	checkedOut := checkoutBook(&library, "Book2", &member) // lending a book out
	fmt.Println("\nCheck out a book:")
	if checkedOut {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}

	returned := returnBook(&library, "Book2", &member) // returning a book
	fmt.Println("\nCheck in a book:")

	if returned {
		printLibraryBooks(&library)
		printLibraryAudit(&library)
	}
}
