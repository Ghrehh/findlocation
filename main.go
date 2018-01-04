package main

import (
	"fmt"
)

func main() {
	locationFinder := NewLocationFinder()
	f := locationFinder.FindLocation("Tasmania Stuff")
	fmt.Println(f)
}
