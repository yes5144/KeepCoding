package main

//  https://github.com/olivere/elastic-with-docker/blob/v6/app/main.go
import (
	"context"
	"flag"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"

	"github.com/olivere/elastic"
)

func main() {
	var (
		esURL = flag.String("url", "http://192.168.204.222:9200/", "ElasticSearch url string")
	)

	flag.Parse()
	if *esURL == "" {
		log.Fatal("missing -url flag")
	}

	url, err := url.Parse(*esURL)
	if err != nil {
		log.Printf("esURL 格式不正确, Failed code: %#v", err)
	}
	log.Printf("Running %s\n", runtime.Version())

	ips, err := net.LookupIP(url.Hostname())
	if err != nil {
		log.Printf("url ip解析错误, Failed code: %#v", err)
	}
	for _, v := range ips {
		log.Println("url中的ip", v)
	}

	// Check Es version and status
	{
		log.Println("Retrieving: ", *esURL)
		res, err := http.DefaultClient.Get(*esURL)
		if err != nil {
			log.Printf("http get error, Failed code: %#v", err)
		}
		defer res.Body.Close()
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Printf("body error, Failed code: %#v", err)
		}

		log.Println(string(body))
	}

	// connect to es
	client, err := elastic.NewClient(elastic.SetURL(*esURL), elastic.SetSniff(false))
	if err != nil {
		log.Printf("连接es失败, Failed code: %#v", err)
		panic(err)
	}
	log.Println("终于成功连接到es")
	log.Printf("Type: %T, value: %#v", client, client)

	// client
	ctx := context.Background()
	info, err := client.NodesInfo().Do(ctx)
	if err != nil {
		log.Printf("info, Failed code: %#v", err)
	}
	log.Println(info.ClusterName, len(info.Nodes))
	for k, v := range info.Nodes {
		log.Println(k, v)
	}

}
