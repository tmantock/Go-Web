package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("cat file1 file2 ...\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	for _, fn := range flag.Args() {
		f, err := os.Open(fn)
		if err != nil {
			panic(err)
		}
		_, err = io.Copy(os.Stdout, f)
		if err != nil {
			panic(err)
		}
	}
}
