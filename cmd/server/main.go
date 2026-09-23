package main

import (
	"io"
	"log"
	"net"
	"spider/internal/database/compute"
	"spider/internal/initialization"
)

func main() {

	i, err := initialization.CreateInitializer()
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := i.Server.Listener.Accept()
		if err != nil {
			log.Println("Acception error: " + err.Error())
		}
		go HandleQueries(conn, i.Server.Compute)
	}

}

func HandleQueries(conn net.Conn, compute compute.Compute) {

	defer conn.Close()

	buf := make([]byte, 1024)

	for {
		n, err := conn.Read(buf)
		if err != nil && err != io.EOF {
			log.Println("Error reading connection: ", err.Error())
			return
		}

		line := string(buf[:n])

		result, err := compute.Parse(line)
		if err != nil {
			conn.Write([]byte(err.Error() + "\n"))
			continue
		}
		conn.Write([]byte("Ваша Команда: " + result.Command + "; key: " + result.Arguments[0] + "\n"))
	}
}
