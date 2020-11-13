package main

import (
	"fmt"
	"net/http"

	"github.com/yes5144/KeepCoding/ginBlogApi/pkg/setting"
	"github.com/yes5144/KeepCoding/ginBlogApi/routers"
)

func main() {
	router := routers.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    setting.ReadTimeout,
		WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
