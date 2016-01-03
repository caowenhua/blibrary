// main
package main

import (
	"fmt"

	"../controller"
	"../db"
)

func main() {
	error := db.InitData()
	if error != nil {
		fmt.Println("error!!!", error)
		return
	}
	for true {
		choice := -1
		controller.Menu(&choice)
		fmt.Println("choice:", choice)
		if choice == 0 {
			return
		}
		controller.Function(choice)
	}

}
