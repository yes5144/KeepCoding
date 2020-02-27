package main

import (
	"log"
	"net/http"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	log.Printf("path: %s, type: %T", r.URL.Path, r.URL.Path)
	log.Println(r.Form["url_long"])
	for k, v := range r.Form {
		log.Println("r.Form:", k, v)
	}
	w.Write([]byte("你的访问路径是：" +

		r.URL.Path))

}

func main() {
	http.HandleFunc("/", sayHelloName)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Printf("http web运行错误, Failed code: %#v", err)
	}
}
