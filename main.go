package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := ":4444"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	fmt.Println("Сервер запущен на порту " + port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}

		go handler(conn)

	}

}

func handler(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)

	n, err := conn.Read(buf)
	if err != nil {
		return
	}
	line := string(buf[:n])
	fmt.Println("received:", line)
	conn.Write([]byte("You've send" + line + "\n"))

}

/*





func main() {
	listener, err := net.Listen("tcp", ":3223")
	if err != nil {
		fmt.Println("error:", err)
		os.Exit(1)
	}
	defer listener.Close()

	fmt.Println("server started on :3223")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			return
		}
		line := string(buf[:n])
		fmt.Println("received:", line)
		conn.Write([]byte("You've send" + line + "\n"))
	}

}
*/
