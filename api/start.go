package api

import (
	"log"
	"net/http"
)

func Start() {

	s := NewService()
	s.Routes()

	port := "8080"

	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Panicf("error in starting http server at port %s", port)
	}

	log.Printf("server listening in port %s", port)

}
