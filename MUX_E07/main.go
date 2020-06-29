// Create basic TCP server.
package main

import (
    "net"
    "log"
    "bufio"
    "fmt"
)

func serve(conn net.Conn) {
    defer conn.Close()
    scanner := bufio.NewScanner(conn)
    for scanner.Scan() {
        ln := scanner.Text()
        if ln == "" {
            break
        }
        fmt.Println(ln)
    }
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
