package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	conn, err := net.Dial("tcp4", "localhost:12345")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	_, err = conn.Write([]byte("time\n"))
	if err != nil {
		log.Fatal(err)
	}

	for {
		reader := bufio.NewReader(conn)
		b, err := reader.ReadBytes('\n')
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Ответ от сервера: ", string(b))
	}
}
