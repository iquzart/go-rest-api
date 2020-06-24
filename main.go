package main

import (
	"github.com/iquzart/go-rest-api/api/router"
)

func main() {

	routersInit := router.InitRouter()

	routersInit.Run()

	//s := &http.Server{
	//	Addr:           ":3000",
	//	Handler:        routersInit,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//s.ListenAndServe()

}
