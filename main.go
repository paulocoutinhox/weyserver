package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/prsolucoes/weyserver/server"
)

func main() {
	host := flag.String("h", "", "server host to bind")
	port := flag.Int("p", 80, "server port to bind")
	directory := flag.String("d", os.Getenv("WEY_SERVER_DIRECTORY"), "directory to be served")

	flag.Parse()

	gin.SetMode(gin.ReleaseMode)

	router := gin.New()
	router.Use(server.Logger())
	router.StaticFS("/", http.Dir(*directory))

	log.Println("> WeyServer started")

	err := router.Run(fmt.Sprintf("%v:%d", *host, *port))

	if err != nil {
		log.Printf("Error on start WeyServer: %v\n", err)
	}
}
