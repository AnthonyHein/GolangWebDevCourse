// Create basic TCP server.
package main

import (
    "net"
    "log"
    "io"
)


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
        // Handle the connection in a new goroutine.
        // The loop then returns to accepting, so that
        // multiple connections may be served concurrently.
        go func(c net.Conn) {
            // Echo all incoming data.
            io.WriteString(c, "I see you connected.")
            // Shut down the connection.
            c.Close()
        }(conn)
    }
}
