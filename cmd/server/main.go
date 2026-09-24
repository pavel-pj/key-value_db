package main

import (
	"fmt"
	"log"
	"spider/internal/initialization"
)

func main() {

	init, err := initialization.CreateInitializer()
	if err != nil {
		log.Fatal(err)
	}
	err = initialization.StartDatabase(init)
	if err != nil {
		fmt.Println(err)
	}

}
