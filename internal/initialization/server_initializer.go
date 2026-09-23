package initialization

import (
	"fmt"
	"log"
	"net"
	"spider/internal/database/compute"
)

type Server struct {
	Compute  compute.Compute
	Listener net.Listener
}

func CreateServer(compute compute.Compute) (Server, error) {

	port := ":3333"

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
		return Server{}, err
	}

	fmt.Print("Server up at port " + port + "\n")

	//defer listener.Close()

	server := Server{
		Compute:  compute,
		Listener: listener,
	}
	return server, nil

}

/*
	for {

		conn, err := listener.Accept()
		if err != nil {
			log.Println("Acception error: " + err.Error())

		}
		go HandleQueries(conn)
	}*/
