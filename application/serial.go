package main

import (
	"fmt"
	"log"

	"github.com/tarm/serial"
)

func main() {
	config := &serial.Config{
		Name:        "COM20",
		Baud:        115200,
		ReadTimeout: 1,
		Size:        8,
	}

	stream, err := serial.OpenPort(config)
	if err != nil {
		log.Fatal(err)
	}

	//stream.Write([]byte("L ff0000 ff0000\n"))
	stream.Write([]byte("I\n"))
	
	buf := make([]byte, 1024)
	for {
		n, err := stream.Read(buf)
		if err != nil {
			log.Fatal(err)
		}
		s := string(buf[:n])
		fmt.Println(s)
	}
}
