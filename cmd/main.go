package main

import (
	"github.com/dollarkillerx/gin-template/internal/server"
	"github.com/dollarkillerx/gin-template/internal/utils"

	"log"
)

func main() {
	utils.InitJWT()

	log.SetFlags(log.LstdFlags | log.Llongfile)

	server := server.NewServer()
	if err := server.Run(); err != nil {
		log.Fatalln(err)
	}
}
