// Create basic TCP server, 13_hands-on
package main

import (
    "net"
    "log"
    "io"
    "bufio"
    "fmt"
    "strings"
)

func serve(conn net.Conn) {
    defer conn.Close()
    var i int
    var rMethod, rURI string
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        if i == 0 {
			// we're in REQUEST LINE
			xs := strings.Fields(ln)
			rMethod = xs[0]
			rURI = xs[1]
			fmt.Println("METHOD:", rMethod)
			fmt.Println("URI:", rURI)
		}
        if ln == "" {
            break
        }
        fmt.Println(ln)
        i++
    }
    body := "I see you connected.\n"
    io.WriteString(conn, "HTTP/1.1 200 OK\r\n")
    fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
    fmt.Fprint(conn, "Content-Type: text/plain\r\n")
    io.WriteString(conn, "\r\n")
    io.WriteString(conn, body)
}



func main() {
    // TAKEN FROM DOCS
    // Listen on TCP port 2000 on all available unicast and
    // anycast IP addresses of the local system.
    l, err := net.Listen("tcp", ":8080")
    if err != nil {
        log.Fatal(err)
    }
    defer l.Close()
    for {
        // Wait for a connection.
        conn, err := l.Accept()
        if err != nil {
            log.Fatal(err)
        }

        go serve(conn)
    }
}
