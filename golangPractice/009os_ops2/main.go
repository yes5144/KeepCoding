package main

import (
	"log"

	"github.com/yes5144/KeepCoding/golangPractice/009os_ops2/utils"
)

func main() {
	// err := os.Mkdir("abc", os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("os.Mkdir, Failed code: %v", err)
	// }

	// err = os.MkdirAll("dir_a/dir_b/dir_c", os.ModePerm)
	// if err != nil {
	// 	log.Fatalf("os.MkdirAll, Failed code: %#v", err)
	// }

	log.Println("hehh")
	utils.CopyDir("H:\\temp", "H:\\xxx")
	utils.CopyDir("H:/jenkins.war", "H:/ttt")
	utils.ZipDir("H:\\temp", "H:\\xxx.zip")
	utils.ZipDir("H:/jenkins.war", "H:/jenk.zip")
}
