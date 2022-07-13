package main

import (
	"MyWallet/api/server"
	"MyWallet/app/cmd"
	"log"
)

func main() {
	log.Fatal(server.RunAPI("127.0.0.1:8080"))
	cmd.Execute()
}
