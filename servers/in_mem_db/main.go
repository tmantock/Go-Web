package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			panic(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	//Instructions
	io.WriteString(conn, "\nIN-MEMORY-DATABASE\n\n"+
		"USE:\n"+
		"SET key value \n"+
		"GET key\n"+
		"DEL key \n\n"+
		"EXAMPLE:\n"+
		"SET fav chocolate \n"+
		"GET fav \n\n\n")
	//read & write
	data := make(map[string]string)
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fs := strings.Fields(ln)
		fs[0] = strings.ToUpper(fs[0])
		//logic
		switch fs[0] {
		case "GET":
			k := fs[1]
			v := data[k]
			if v != "" {
				fmt.Fprintf(conn, "%s\n", v)
			} else {
				fmt.Fprintln(conn, "NOT FOUND")
			}
		case "SET":
			if len(fs) != 3 {
				fmt.Fprintln(conn, "EXPECTED VALUE")
				continue
			}
			k := fs[1]
			v := fs[2]
			data[k] = v
		case "DEL":
			k := fs[1]
			delete(data, k)
		default:
			fmt.Fprintln(conn, "INVALID COMMAND")
		}
	}
}
