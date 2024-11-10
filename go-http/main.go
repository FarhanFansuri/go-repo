package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.GET("/nomatch/:id", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Barang dengan id : " + params.ByName("id")
		fmt.Fprint(writer, text)
	})
	router.GET("/path/:id/*path", func(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
		text := "Barang dengan id : "
		fmt.Fprint(writer, text)
	})

	router.PanicHandler = func()

	server := http.Server{
		Handler: router,
		Addr:    "localhost:3000",
	}

	server.ListenAndServe()
}
