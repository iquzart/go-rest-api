package main

import (
	"github.com/iquzart/go-rest-api/api/router"
)

func main() {

	routersInit := router.InitRouter()

	routersInit.Run()
}
