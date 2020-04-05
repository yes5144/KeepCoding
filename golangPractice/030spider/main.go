package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

var (
	// tarUrl = "http://www.haomagujia.com/"

	// var rePhone = `1[3456789]\d{9}`
	// 正则分组
	rePhone = `(1[3456789]\d)(\d{4})(\d{4})`
	// 10000332@qq.com
	// reEmail = `[1-9]\d{4,}@qq.com`
	// 2kd3@er2k.com.cn
	// \w 所有字符， () 分组  ?一次或没有(非贪婪)
	reEmail = `[\w\.]+@\w+\.[a-z]{2,5}(\.[a-z]{2.5})?`
	// <a ... href="https://www.baidu.com?jflj">jdj</a>
	tarUrl = "http://www.hao123.com/"
	// reLink = `<a[\s\S]+?href="(http[\s\S]+?)"`
	// [\s\S] 包含\t \n \r等
	reLink = `<a.*?href="(http.*?)"`
	// 身份证号 41020220020202100X
	reID    = `[1-6]\d{5}-(year)-(month)-(day)-3232X`
	reImg   = `<img.*?src="(http.*?)"`
	chSigle = make(chan int, 2)
	dloadWG sync.WaitGroup
	// 锁
	randomMT sync.Mutex
)

func main() {
	html := getHtml(tarUrl)

	re := regexp.MustCompile(reImg)
	allImg := re.FindAllStringSubmatch(html, -1)
	imgUrls := make([]string, 0)
	for _, v := range allImg {
		// log.Println(v[1])
		imgUrls = append(imgUrls, v[1])
	}
	log.Println(imgUrls)

	// strings
	testStr := "wo_shi_wk.tech"
	testStr = strings.ReplaceAll(testStr, "_", "-")
	log.Println("")
	log.Println(strings.LastIndex(testStr, "."))
	log.Println(testStr[strings.LastIndex(testStr, ".")+1:])
	// strings.Trim(testStr,"_")

}

func GetRandomInt(start, end int) int {
	randomMT.Lock()
	<-time.After(1 * time.Nanosecond)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	ret := start + r.Intn(end-start)
	randomMT.Unlock()
	return ret
}

func DownLoadimg(filename, url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("open url err, Failed code %#v", err)
	}
	defer resp.Body.Close()
	imgbytes, _ := ioutil.ReadAll(resp.Body)
	err = ioutil.WriteFile(filename, imgbytes, 0644)
	if err != nil {
		log.Fatalf("下载失败, Failed code %#v", err)
	}
	log.Println("下载成功！")

}

func DownLoadimgAsync(filename, url string) {
	dloadWG.Add(1)
	go func() {
		chSigle <- 123
		DownLoadimg(filename, url)
		<-chSigle
		dloadWG.Done()
	}()
	dloadWG.Wait()
}

func spiderPhone() {
	html := getHtml(tarUrl)
	html += "11324242299"
	// log.Println(html)
	re := regexp.MustCompile(rePhone)
	// -1 匹配全部，否则只是匹配次数
	allstring := re.FindAllStringSubmatch(html, 6)
	log.Println(allstring)
	for _, v := range allstring {
		log.Println(v[1])
	}
}

func spiderLink() {
	html := getHtml(tarUrl)
	re := regexp.MustCompile(reLink)
	allstring := re.FindAllStringSubmatch(html, -1)
	// log.Println(allstring)
	for _, v := range allstring {
		log.Println(v[1])
	}
}

func getHtml(tarUrl string) string {
	// tarUrl := "http://www.haoma.com/xh/?lanmu=0"
	resp, err := http.Get(tarUrl)
	if err != nil {
		log.Fatalf("err, Failed code %#v", err)
	}
	bytes, _ := ioutil.ReadAll(resp.Body)
	html := string(bytes)
	return html
}
