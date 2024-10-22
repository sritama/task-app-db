package api

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Start() {

	s := NewService()
	s.Routes()

	port := "8080"

	server := &http.Server{
		Addr:    ":8080",
		Handler: s.router,
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("error in starting http server at port %s", port)
		}
	}()

	log.Printf("server listening in port %s", port)

	<-stop

	log.Printf("Starting graceful shutdown of server")

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(10)*time.Second)

	defer func() {
		err := s.TaskDB.Close()
		if err != nil {
			log.Fatalf("error in shutting down DB %s\n", err)
		}
		fmt.Println("Released connections")
		cancel()
	}()

	err := server.Shutdown(ctx)
	if err != nil {
		log.Fatalf("error in shutting down server %s\n", err)
	}

	fmt.Println("graceful shutdown of http server complete")
}
