// book
package bean

import "fmt"

type book struct {
	price  float64
	name   string
	isdn   string
	author string
	time   int64
}

func createBook2(price float64, name string, isdn string, author string, time int64) *book {
	if price < 0 {
		return nil
	}

	return &book{price, name, isdn, author, time}
}

func CreateBook(price float64, name string, isdn string, author string, time int64) *book {
	if price < 0 {
		fmt.Println("the price must be larger than 0.0")
		return nil
	}

	book := new(book)
	(*book).price = price
	(*book).name = name
	(*book).isdn = isdn
	(*book).author = author
	(*book).time = time
	return book
}

func CreateNilBook() *book {
	book := new(book)
	return book
}

func (b *book) ChangePrice(price float64) {
	b.price = price
}
