// main
package main

import (
	"fmt"

	"../controller"
)

func main() {
	choice := 0
	controller.Menu(&choice)
	fmt.Println(choice)
	controller.Function(choice)
}
