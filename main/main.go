// main
package main

import (
	"fmt"

	"../ui"
)

func main() {
	choice := 0
	ui.Menu(&choice)
	fmt.Println(choice)
	ui.Function(choice)
}
