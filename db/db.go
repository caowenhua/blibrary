package db

//map{
//	slice[]//book, person
//}
import (
	"../bean"

	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

var Data data

type data struct {
	books []*bean.Book
}

func SaveToFile() {
	js, _ := json.Marshal(Data)
	fmt.Printf("JSON format: %s", js)

	file, _ := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(Data)
	if err != nil {
		fmt.Println("Error in encoding json")
	}
}

func LoadFromFile() error {
	//	inputFile, inputError := os.Open("data.txt")
	//	if inputError == nil {
	//		fmt.Println("Open file failed!!!!!!!!")
	//		return d, errors.New("open failed")
	//	}
	//	defer inputFile.Close()
	b, ioError := ioutil.ReadFile("data.txt")
	if ioError == nil {
		fmt.Println("Open file failed!!!!!!!!")
		return errors.New("open failed")
	}
	err := json.Unmarshal(b, &Data)
	if err == nil {
		fmt.Println("Load file failed!!!!!!!!")
	}
	return nil
}

func initFile() {
	//	m := new(map[string]data)
	//	m["book"] = new([]Book, 0)
	Data := new(data)
	fmt.Println(Data)
	SaveToFile()
}

func InitData() error {
	_, ioError := ioutil.ReadFile("data.txt")
	if ioError == nil {
		fmt.Println("read file error!")
		initFile()
	}
	return LoadFromFile()
}

func AddBook(b *bean.Book) {
	Data.books = append(Data.books, b)
	SaveToFile()
}

func FindBookById(isdn string) (result []bean.Book) {
	result = make([]bean.Book, 0)
	for _, b := range Data.books {
		if b.Isdn == isdn {
			result = append(result, *b)
		}
	}
	return
}

func FindBookByName(name string) (result []bean.Book) {
	result = make([]bean.Book, 0)
	for _, b := range Data.books {
		if b.Name == name {
			result = append(result, *b)
		}
	}
	return
}
func FindBookByAuthor(author string) (result []bean.Book) {
	result = make([]bean.Book, 0)
	for _, b := range Data.books {
		if b.Author == author {
			result = append(result, *b)
		}
	}
	return
}
func FindBookByPrice(price float64) (result []bean.Book) {
	result = make([]bean.Book, 0)
	for _, b := range Data.books {
		if b.Price == price {
			result = append(result, *b)
		}
	}
	return
}
func FindBookByTime(time int64) (result []bean.Book) {
	result = make([]bean.Book, 0)
	for _, b := range Data.books {
		if b.Time == time {
			result = append(result, *b)
		}
	}
	return
}
func DeleteBookById(isdn string) {
	for i, b := range Data.books {
		if b.Isdn == isdn {
			kk := i + 1
			Data.books = append(Data.books[:i], Data.books[kk:]...)
		}
	}
}
