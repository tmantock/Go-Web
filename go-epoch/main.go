package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var cmd string

	cmd = strings.ToUpper(os.Args[1])

	if cmd == "TIME" {
		stamp, err := strconv.ParseInt(os.Args[2], 10, 64)
		if err != nil {
			panic(err)
		}
		tm := time.Unix(stamp, 0)
		fmt.Println(tm)
	}

	// if cmd == "EPOCH" {
	// 	var str string

	// 	strtime := os.Args[2:]

	// 	for _, v := range strtime {
	// 		str = str + " " + v
	// 	}

	// 	t, err := time.Parse(nil, str)
	// 	if err != nil {
	// 		panic(err)
	// 	}

	// 	fmt.Println(t.Unix())

	// }

}
