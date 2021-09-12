package main

import (
	"flag"
	"fmt"

	"go-server-template/api"
)

func main() {
	var port int
	//var mapFile string

	flag.IntVar(&port, "port", 8380, "server listening port")
	//flag.StringVar(&mapFile, "mapfile", "", "path to synonym map file")
	flag.Parse()

	api.SetupServer().Run(fmt.Sprintf(":%d", port))
}
