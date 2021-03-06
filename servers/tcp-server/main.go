package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err)
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	//read request
	request(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			//headers are done
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {
	//request line
	method := strings.Fields(ln)[0]
	uri := strings.Fields(ln)[1]
	fmt.Println("***METHOD", method)
	fmt.Println("***URI", uri)

	//multiplexer
	if method == "GET" && uri == "/" {
		index(conn)
	}

	if method == "GET" && uri == "/about" {
		about(conn)
	}

	if method == "POST" && uri == "/apply" {
		apply(conn)
	}
}

func index(conn net.Conn) {
	body := `
        <!DOCTYPE html>
        <html>
            <head>
                <title>GO TCP Server</title>
            </head>
            <body>
                <h1>Hello Welcome to Mars!</h1>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/about">About</a></li>
                </ul>
                <form method="POST" action="/apply">
                    <button type="submit">Apply</button>
                </form>
            </body>
        </html>
    `
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func about(conn net.Conn) {
	body := `
        <!DOCTYPE html>
        <html>
            <head>
                <title>GO TCP Server</title>
            </head>
            <body>
                <h1>Hello Learn About Mars!</h1>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/about">About</a></li>
                </ul>
                <form method="POST" action="/apply">
                    <button type="submit">Apply</button>
                </form>
            </body>
        </html>
    `
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}

func apply(conn net.Conn) {
	body := `
        <!DOCTYPE html>
        <html>
            <head>
                <title>GO TCP Server</title>
            </head>
            <body>
                <h1>Hello Apply to Fly to Mars!</h1>
                <ul>
                    <li><a href="/">Home</a></li>
                    <li><a href="/about">About</a></li>
                </ul>
                <form method="POST" action="/apply">
                    <button type="submit">Apply</button>
                </form>
            </body>
        </html>
    `
	fmt.Fprintf(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprintf(conn, "Content-Type: text/html\r\n")
	fmt.Fprintf(conn, "\r\n")
	fmt.Fprintf(conn, body)
}
