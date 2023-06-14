package main

import "fmt"

type Books struct {
	title   string
	author  string
	subject string
	book_id int
}

func main() {
	var book1 Books
	var book *Books
	//var book2 Books
	book = &book1
	book1.title = "Ape gedara"
	book1.subject = "Sinhala"
	book1.author = "Amaneris deyogu"
	book1.book_id = 9189

	printPointerbook(book)

}

func printPointerbook(book *Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)
}

func printBook(book Books) {
	fmt.Printf("Book title : %s\n", book.title)
	fmt.Printf("Book author : %s\n", book.author)
	fmt.Printf("Book subject : %s\n", book.subject)
	fmt.Printf("Book book_id : %d\n", book.book_id)

}
