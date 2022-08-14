package main

import "github.com/vnmoliveira/webapi-go/server"

func main() {
	server := server.NewServer()

	server.Run()

}
