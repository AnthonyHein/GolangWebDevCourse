
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
        log.Fatalln(err)
    }
    defer li.Close()

    for {
        conn, err := li.Accept()
        if err != nil {
            log.Fatalln(err)
            continue
        }

        go handle(conn)
    }
}

func handle(conn net.Conn) {
    defer conn.Close()

    request(conn)
}

func request(conn net.Conn) {
    i := 0

    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        fmt.Println(ln)
        if i == 0 {
            reqln := strings.Fields(ln)
            m := reqln[0]
            fmt.Println("***METHOD", m)
            url := reqln[1]
            fmt.Println("***URL", url)
            if url == "/home" {
                defer respond1(conn)
            } else if url == "/about" {
                defer respond2(conn)
            } else {
                defer respond3(conn)
            }

        }
        if ln == "" {
            break
        }
        i++
    }
}

func respond1(conn net.Conn) {
    body := `<!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>
    </title>
    </head>
    <body>
    <strong>
    Home
    </strong>
    </body>
    </html>`

    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    fmt.Fprint(conn, body)
}

func respond2(conn net.Conn) {
    body := `<!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>
    </title>
    </head>
    <body>
    <strong>
    About Us
    </strong>
    </body>
    </html>`

    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    fmt.Fprint(conn, body)
}


func respond3(conn net.Conn) {
    body := `<!DOCTYPE html>
    <html lang="en">
    <head>
    <meta charset="UTF-8">
    <title>
    </title>
    </head>
    <body>
    <strong>
    Anthony Hein
    </strong>
    </body>
    </html>`

    fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/html\r\n")
    fmt.Fprint(conn, "\r\n")
    fmt.Fprint(conn, body)
}
