// main
package main

import (
	"fmt"

	"../controller"
	"../db"
)

func main() {
	data, error := db.InitData()
	if error != nil {
		fmt.Println("error!!!", error)
		return
	}
	fmt.Println(data)
	for true {
		choice := 0
		controller.Menu(&choice)
		fmt.Println("choice:", choice)
		controller.Function(choice)
	}

}
