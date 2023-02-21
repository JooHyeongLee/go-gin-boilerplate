package server

import (
	"fmt"
	"go-gin-boilerplate/config"
)

func Init() {
	config := config.GetConfig()
	r := NewRouter()
	r.Run(fmt.Sprintf(":%s", config.GetString("server.port")))
}
