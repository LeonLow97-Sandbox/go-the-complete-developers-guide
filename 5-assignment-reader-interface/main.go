package main

import (
	"fmt"
	"os"
)

func main() {
	// Access to the file name from command line argument list
	fmt.Println(os.Args[1])

}
