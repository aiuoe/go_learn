package poo

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	title  string
	author string
	pages  int
}

func NewBook(title, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

func (book *Book) GetTitle() string {
	return book.title
}

func (book *Book) GetAuthor() string {
	return book.author
}

func (book *Book) GetPages() int {
	return book.pages
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d\n", book.title, book.author, book.pages)
}

type TextBook struct {
	Book
	editorial string
	level     int
}

func NewTextBook(title, author string, pages int, editorial string, level int) *TextBook {
	return &TextBook{
		Book:      *NewBook(title, author, pages),
		editorial: editorial,
		level:     level,
	}
}

func (txtBook *TextBook) PrintInfo() {
	fmt.Printf("Title: %s, Author: %s, Pages: %d, Editorial: %s, Level: %d\n", txtBook.title, txtBook.author, txtBook.pages, txtBook.editorial, txtBook.level)
}

func ExecStruct() {
	myBook := NewBook("Pattern of Design", "Alexander Shvet", 437)
	myBook.SetTitle("Clean code")

	myTextBook := NewTextBook(myBook.GetTitle(), myBook.GetAuthor(), myBook.GetPages(), "O'Reilly", 1)
	myTextBook.Book.SetTitle("Solid Principles")

	Print(myBook)
	Print(myTextBook)
}
