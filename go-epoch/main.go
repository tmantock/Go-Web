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

	if len(os.Args) > 0 {

		cmd = strings.ToUpper(os.Args[1])

		if cmd == "TIME" {
			var tm time.Time

			if len(os.Args) > 2 {
				stamp, err := strconv.ParseInt(os.Args[2], 10, 64)
				if err != nil {
					panic(err)
				}
				tm = time.Unix(stamp, 0)
			} else {
				tm = time.Now()
			}
			fmt.Println(tm)
		}
	}
}
