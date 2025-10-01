package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"

	"gitlab.vk-golang.ru/vk-golang/lectures/08_microservices/7_grpc_stream/translit"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalln("can't listen port", err)
	}

	server := grpc.NewServer()

	tr := NewTr()
	translit.RegisterTransliterationServer(server, tr)

	fmt.Println("starting server at :8081")
	log.Fatal(server.Serve(lis))
}
