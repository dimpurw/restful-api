package main

import (
	"fmt"
	"net/http"
	"restful-api/helper"

	"github.com/julienschmidt/httprouter"
)

func main()  {
	fmt.Printf("Start Server")

	routes := httprouter.New()

	routes.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w,"Welcome!\n")
	})

	server := http.Server{Addr: "localhost:4000", Handler: routes}

	err := server.ListenAndServe()
	helper.PanicIfError((err))
}