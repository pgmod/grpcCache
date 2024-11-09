package main

import "github.com/pgmod/grpcCache/server"

func main() {
	server.StartServer(3, "50051")
}
