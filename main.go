// https://gist.github.com/paulsmith/775764

package main

import (
    "io"
    "net"
    "strconv"
    "fmt"
)

const PORT = 7778

func main() {
    ServerAddr, err := net.ResolveUDPAddr("udp", ":" + strconv.Itoa(PORT))
    server, err := net.ListenUDP("udp", ServerAddr)
	
    if server == nil {
        panic("couldn't start listening: " + err.String())
    }
    conns := clientConns(server)
    for {
        go handleConn(<-conns)
    }
}

func clientConns(listener net.Listener) chan net.Conn {
    ch := make(chan net.Conn)
    i := 0
    go func() {
        for {
            client, err := listener.Accept()
            if client == nil {
                fmt.Printf("couldn't accept: " + err.String())
                continue
            }
            i++
            fmt.Printf("%d: %v <-> %v\n", i, client.LocalAddr(), client.RemoteAddr())
            ch <- client
        }
    }()
    return ch
}

func handleConn(client net.Conn) {
    io.Copy(client, client);
}
