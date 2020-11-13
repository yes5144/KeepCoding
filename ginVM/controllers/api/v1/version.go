package v1

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/yes5144/KeepCoding/ginVM/models"
	"ginVM/models"
)

func NewOneVersionApi(c *gin.Context) {
	c.Request.ParseForm()

	c.JSON(http.StatusOK, gin.H{
		"msg": "new one ",
	})
}

func GetOneVersionApi(c *gin.Context) {
	id := c.Param("id")
	log.Println(id)
	nid, err := strconv.Atoi(id)
	if err != nil {
		log.Fatalf("Atoi error, Failed code: %#v", err)
	}
	data := models.GetOneVersion(nid)
	// data.version
	// version := "1.3.22"
	newVersionList := strings.Split(data.Version, ".")
	lastVer, _ := strconv.Atoi(newVersionList[len(newVersionList)-1])
	log.Printf("Type: %T, value: %#v", lastVer, lastVer)
	nLVer := fmt.Sprintf("%d", lastVer+1)
	log.Printf("Type: %T, value: %#v", nLVer, nLVer)
	newVersionList[len(newVersionList)-1] = nLVer
	newVersion := strings.Join(newVersionList[:], ".")
	log.Println(newVersion)

	data.Version = newVersion

	cmd := exec.Command("python3 ../scripts/zipByapp.py")
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err = cmd.Start()
	log.Println(stderr.String())
	log.Println(stdout.Bytes())
	// out, err := exec.Command("python ../scripts/zipByApp.py").Output()
	if err != nil {
		log.Printf("命令执行失败：, Failed code: %#v", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"msg":    "get one version api",
		"status": http.StatusOK,
		"data":   data,
	})
}

func GetAllVersionApi(c *gin.Context) {
	data := models.GetAllVersion()
	log.Println("from api v1 version", data)
	c.JSON(http.StatusOK, gin.H{
		"msg":    "all versions",
		"data":   data,
		"status": http.StatusOK,
	})

}
