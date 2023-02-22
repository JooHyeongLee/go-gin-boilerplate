package main

import (
	"flag"
	"fmt"
	"go-gin-boilerplate/config"
	"go-gin-boilerplate/db"
	"go-gin-boilerplate/server"
	"os"
)

func main() {
	environment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: server -e {mode}")
		os.Exit(1)
	}
	flag.Parse()
	config.Init(*environment)
	db.Init()
	server.Init()
}
