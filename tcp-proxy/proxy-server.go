package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func hadleConnection(localConn net.Conn, remoteAddr string) {
	defer localConn.Close()

	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Println("Cannot make remote addr dial with error: ", err)
	}
	defer remoteConn.Close()

	go func() {
		_, err := io.Copy(remoteConn, localConn)
		fmt.Println(remoteConn)
		fmt.Println(localConn)
		if err != nil && err != io.EOF {
			log.Println("Error copying from local request to remote request: ", err)
		}
	}()
	_, err = io.Copy(localConn, remoteConn)
	if err != nil && err != io.EOF {
		log.Println("Cannot copying response from remote to local: ", err)
	}
}

func main() {
	localPort := ":8080"
	// remotTarget := "185.204.168.236:3128"
	remotTarget := "188.121.106.132:3128"

	listener, err := net.Listen("tcp", localPort)
	if err != nil {
		log.Println("Cannot make localPort with err: ", err)
	}
	defer listener.Close()

	fmt.Printf("Listening on %s, forwarding to server %s\n", localPort, remotTarget)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accept connection with err: ", err)
			continue
		}
		go hadleConnection(conn, remotTarget)
	}

}
