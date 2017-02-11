package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	//argCount := len(os.Args[1:])
	//cmd := os.Args[0]

	// for i, a := range os.Args[1:] {
	// 	fmt.Printf("Arguments %d is %s\n", i+1, a)
	// }

	//fmt.Printf("Total Arguments (excluding program name): %d\n", argCount)

	// var port int

	// flag.IntVar(&port, "p", 8000, "Specify port to use. Defaults to 8000")
	// flag.Parse()

	// fmt.Printf("Port = %d\n", port)
	// fmt.Printf("Other args: %+v\n", flag.Args())

	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("    basic file1 file2 ... \n")
		flag.PrintDefaults()
	}
	flag.Parse()
}
