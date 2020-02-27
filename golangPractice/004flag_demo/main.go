package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"runtime"
	"sort"
)

func main() {
	var (
		esURL = flag.CommandLine.String("url", "http://192.168.204.222:9200/", "ElasticSearch conn string")
	)
	flag.Parse()
	if *esURL == "" {
		log.Fatal("missing -url flag, 使用默认 http://192.168.204.222:9200")
		// esURL = "http://localhost:9200"
	}
	url, err := url.Parse(*esURL)
	if err != nil {
		log.Fatalf("无效的 -url flag, Failed code: %#v", err)
	}
	log.Println(url)
	log.Println(len(os.Args))
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "showenv":
			log.Println("Environment ...")
			env := os.Environ()
			sort.Strings(env)
			for _, v := range env {
				log.Println("- ", v)
			}
			os.Exit(0)
		}
	}

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running %s\n", runtime.Version())
	// log.Printf("Version of github.com/olivere/elastic: %s\n", elastic.Version())

	log.Printf("Lookup for hostname %q returns the following IPs:", url.Hostname())
	ips, err := net.LookupIP(url.Hostname())
	if err != nil {
		log.Printf("url.Hostname(), Failed code: %#v", err)
	}
	for k, v := range ips {
		log.Println(k, v)
	}

	// Check ES version and status
	{
		log.Printf("Retrieving %s", *esURL)
		res, err := http.DefaultClient.Get(*esURL)
		if err != nil {
			log.Printf("From url get err, Failed code: %#v", err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("ioutil body err, Failed code: %#v", err)
		}
		log.Printf("%v", string(body))

	}
}
