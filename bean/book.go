// book
package bean

import "fmt"

type Book struct {
	Price  float64
	Name   string
	Isdn   string
	Author string
	Time   int64
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
	(*book).Price = price
	(*book).Name = name
	(*book).Isdn = isdn
	(*book).Author = author
	(*book).Time = time
	return book
}

func CreateNilBook() *Book {
	book := new(Book)
	return book
}

func (b *Book) ChangePrice(price float64) {
	b.Price = price
}

func (b *Book) ToString() {
	fmt.Println("\n书籍名称", b.Name)
	fmt.Println("书籍ISDN", b.Isdn)
	fmt.Println("书籍作者", b.Author)
	fmt.Println("书籍价格", b.Price)
	fmt.Println("书籍出版日期", b.Time, "\n")
}

func ToArrayString(b []Book) {
	if len(b) == 0 {
		fmt.Println("无结果")
	} else {
		for _, book := range b {
			book.ToString()
		}
	}
}
