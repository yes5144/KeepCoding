package utils

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/astaxie/beego/logs"
)

func FormatPath(s string) string {
	switch runtime.GOOS {
	case "windows":
		return strings.Replace(s, "/", "\\", -1)
	case "darwin", "linux":
		return strings.Replace(s, "\\", "/", -1)
	default:
		logs.Info("only support linux,windows,darwin, but os is " + runtime.GOOS)
		return s
	}
}

func CopyDir(src, dest string) {
	src = FormatPath(src)

	dest = FormatPath(dest)

	logs.Info(src)
	logs.Info(dest)

	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		stat, err := os.Stat(src)
		if err != nil {
			logs.Info("os.Stat, Failed code: %#v", err)
		}
		if stat.IsDir() {
			cmd = exec.Command("xcopy", src, dest, "/I", "/E")
		}
		cmd = exec.Command("xcopy", src, dest, "/I", "/k")

	case "darwin", "linux":
		cmd = exec.Command("cp", "-R", src, dest)
	}
	outPut, err := cmd.Output()
	if err != nil {
		logs.Error(", Failed code: %#v", err)
		return
	}
	fmt.Println(string(outPut))
}

func CopyFile(src, dest string) {

}

// https://blog.csdn.net/fwhezfwhez/article/details/97390199
