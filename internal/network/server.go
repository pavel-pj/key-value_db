package network

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"sync"
)

type TCPHandler = func([]byte) []byte

type Server struct {
	BufferSize int
	Listener   net.Listener
}

func CreateServer() (Server, error) {

	port := ":3333"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
		return Server{}, err
	}

	fmt.Print("Server up at port " + port + "\n")

	server := Server{
		BufferSize: 4096,
		Listener:   listener,
	}
	return server, nil

}

func (s *Server) HandleQueries(handler TCPHandler) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for {
			//var buf = make([]byte, 4056)

			conn, err := s.Listener.Accept()
			if err != nil {
				log.Println("Error Connection")
			}
			for {
				request := make([]byte, s.BufferSize)
				count, err := conn.Read(request)

				if err != nil && err != io.EOF {
					fmt.Println("address", conn.RemoteAddr().String())
					break
				} else if count == s.BufferSize {
					fmt.Println("small buffer size" + strconv.Itoa(s.BufferSize))
					break
				}

				response := handler(request[:count])
				if _, err = conn.Write(append(response, '\n')); err != nil {
					fmt.Println("Error: could not send response")
					break
				}

			}

		}
	}()

	//s.Listener.Close()

	//wg.Wait()

	select {}

}
