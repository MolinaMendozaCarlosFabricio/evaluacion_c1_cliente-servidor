package main

import (
	"evaluacionc1.com/e/client"
	"evaluacionc1.com/e/server"
)

func main() {
	server.Run()
	client.Run()
}