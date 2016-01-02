// book
package bean

import "fmt"

type Book struct {
	price  float64
	name   string
	isdn   string
	author string
	time   int64
}

func createBook2(price float64, name string, isdn string, author string, time int64) *Book {
	if price < 0 {
		return nil
	}

	return &Book{price, name, isdn, author, time}
}

func CreateBook(price float64, name string, isdn string, author string, time int64) *Book {
	if price < 0 {
		fmt.Println("the price must be larger than 0.0")
		return nil
	}

	book := new(Book)
	(*book).price = price
	(*book).name = name
	(*book).isdn = isdn
	(*book).author = author
	(*book).time = time
	return book
}

func CreateNilBook() *Book {
	book := new(Book)
	return book
}

func (b *Book) ChangePrice(price float64) {
	b.price = price
}
