package main

import (
	"MyWallet/api/server"
	"log"
)

func main() {
	log.Fatal(server.RunAPI("127.0.0.1:8080"))
}
