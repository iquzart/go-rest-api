package main

import (
	"github.com/iquzart/go-rest-api/api/v1/router"
)

func main() {

	routersInit := router.InitRouter()

	routersInit.Run()
}
