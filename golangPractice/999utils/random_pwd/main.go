package main

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"math/rand"
	"strings"

	"github.com/astaxie/beego/logs"
	"github.com/hashicorp/go-uuid"
)

func main() {
	// uuid
	for i := 0; i < 4; i++ {
		logs.Debug(uuid.GenerateUUID())
	}
	uuidStr, _ := uuid.GenerateUUID()
	log.Println(strings.ReplaceAll(uuidStr, "-", ""))
	// logs.Warn(base64.RawURLEncoding.EncodeToString([]byte(uuidStr)))
	uuidSlice := strings.Split(uuidStr, "-")
	// rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	rand.Seed(99)
	randNum0 := rand.Intn(len(uuidSlice))
	uuidStrNew := uuidSlice[randNum0] + uuidSlice[rand.Intn(len(uuidSlice))] + uuidSlice[rand.Intn(len(uuidSlice))] + uuidSlice[rand.Intn(len(uuidSlice))]
	log.Println(uuidStrNew)
	base64Ret := base64.StdEncoding.EncodeToString([]byte(uuidStr))
	log.Println(len(base64Ret))

	logs.Warn(base64Ret)
	log.Println(rand.Intn(33))
	log.Println("-----------------------")
	// randNum1 := rand.Intn(len(base64Ret) / 4)
	// randNum2 := rand.Intn(len(base64Ret)/4) + len(base64Ret)/2
	rand.Seed(12)
	randNum1 := rand.Intn(12)
	randNum2 := rand.Intn(12) + len(base64Ret)/2
	log.Println(randNum0, randNum1, randNum2)
	randPass := base64Ret[randNum1:randNum1+11] + base64Ret[randNum2:randNum2+11]
	logs.Warn(randPass)

	h := md5.New()
	io.WriteString(h, "The fog is getting thicker!")
	io.WriteString(h, "And Leon's getting laaarger!")
	fmt.Printf("%x", h.Sum(nil))
}
