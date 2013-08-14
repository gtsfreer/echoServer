// PingPong project server main.go
package main

import (
	"flag"
	"log"
	"net"
	"os"
	"runtime"
	"strconv"
	"time"
)

func SocketServer(port int) {

	listen, err := net.Listen("tcp4", ":"+strconv.Itoa(port))
	defer listen.Close()
	if err != nil {
		log.Fatalf("Socket listen port %d failed,%s", port, err)
		os.Exit(1)
	}
	log.Printf("Begin listen port: %d\n", port)

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Println("Socket accept failed,%s", err)
			continue
		}
		go handler(conn)
	}

}

func handler(conn net.Conn) {

	var (
		buf = make([]byte, 1024)
	)
	conn.SetReadDeadline(time.Now().Add(time.Second * 1800))
ILOOP:
	for {
		n, err := conn.Read(buf)
		if err != nil {
			break ILOOP
		}
		conn.Write(buf[:n])
	}
	conn.Close()
}

func main() {
	servport := flag.Int("port", 3333, "listen port of the server!")
	flag.Parse()
	runtime.GOMAXPROCS(runtime.NumCPU())
	SocketServer(*servport)
}
