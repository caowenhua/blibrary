package db

//map{
//	slice[]//book, person
//}
import (
	"../bean"

	"encoding/json"
	"fmt"
)

type data struct {
	books []*bean.Book
}

func SaveToFile() {

}

func LoadFromFile() data {
	var d data
	//TODO getJsonInFile
	var b []byte
	err := json.Unmarshal(b, &d)
	if err == nil {
		fmt.Println("load file failed!!!!!!!!")
	}
	return d
}
