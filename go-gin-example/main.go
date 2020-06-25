package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/yes5144/KeepCoding/go-gin-example/pkg/setting"
	"github.com/yes5144/KeepCoding/go-gin-example/routers"
)

func main() {
	r := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        r,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	log.Println("http_port: ", setting.HTTPPort)
	s.ListenAndServe()
}
