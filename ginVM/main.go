package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	// "github.com/yes5144/KeepCoding/ginVM/routers"
	"ginVM/routers"
)

func main() {
	log.Println("hello")

	router := routers.InitRouters()

	s := &http.Server{
		Addr:    fmt.Sprintf(":%d", 8080),
		Handler: router,
		// ReadTimeout:    200,
		// WriteTimeout:   200,
		// MaxHeaderBytes: 1 << 20,
	}
	go func() {
		if err := s.ListenAndServe(); err != nil {
			log.Printf("error, Failed code: %#v", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdonw Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 25*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown, Failed code: %#v", err)
	}
	log.Println("Server exiting")
}
