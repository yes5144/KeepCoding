package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handler("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
