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

type data struct {
	books []*bean.Book
}

func SaveToFile(d *data) {
	js, _ := json.Marshal(d)
	fmt.Printf("JSON format: %s", js)

	file, _ := os.OpenFile("data.txt", os.O_CREATE|os.O_WRONLY, 0)
	defer file.Close()
	enc := json.NewEncoder(file)
	err := enc.Encode(d)
	if err != nil {
		fmt.Println("Error in encoding json")
	}
}

func LoadFromFile() (data, error) {
	var d data
	//	inputFile, inputError := os.Open("data.txt")
	//	if inputError == nil {
	//		fmt.Println("Open file failed!!!!!!!!")
	//		return d, errors.New("open failed")
	//	}
	//	defer inputFile.Close()
	b, ioError := ioutil.ReadFile("data.txt")
	if ioError == nil {
		fmt.Println("Open file failed!!!!!!!!")
		return d, errors.New("open failed")
	}
	err := json.Unmarshal(b, &d)
	if err == nil {
		fmt.Println("Load file failed!!!!!!!!")
	}
	return d, nil
}

func initFile() {
	//	m := new(map[string]data)
	//	m["book"] = new([]Book, 0)
	d := new(data)
	SaveToFile(d)
}

func InitData() (data, error) {
	_, ioError := ioutil.ReadFile("data.txt")
	if ioError == nil {
		initFile()
	}
	return LoadFromFile()
}
