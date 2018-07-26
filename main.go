// Shamefully stolen from stackoverflow, did a bit of tuning though
// https://stackoverflow.com/questions/28400340/how-support-concurrent-connections-with-a-udp-server-using-go

package main

import (
	"fmt"
	"net"
)

func listen(connection *net.UDPConn, quit chan struct{}) {
	buffer := make([]byte, 1024)
	_, remoteAddr, err := 0, new(net.UDPAddr), error(nil)

	for err == nil {
		_, remoteAddr, err = connection.ReadFromUDP(buffer)
		buffer_tmp := make([]byte, 1024)

		copy(buffer_tmp, buffer)
		buffer = make([]byte, 1024)
		go connection.WriteTo(buffer_tmp, remoteAddr)
	}

	fmt.Println("listener failed - ", err)
	quit <- struct{}{}
}

func main() {
	addr := net.UDPAddr{
		Port: 7778,
		IP:   net.IP{0, 0, 0, 0},
	}

	connection, err := net.ListenUDP("udp", &addr)

	if err != nil {
		panic(err)
	}

	quit := make(chan struct{})

	for i := 0; i < 1; i++ {
		go listen(connection, quit)
	}

	<-quit // hang until an error
}
