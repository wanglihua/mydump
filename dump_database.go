package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func dumpDatabase() {
	var batCmdParam = time.Now().Format("20060102150405")

	var batFileName = "mydump.bat"
	var batFileFullName = workDir + string(os.PathSeparator) + batFileName

	if !pathExists(batFileFullName) {
		log.Println("mydump.bat file not exist!" + batFileFullName)

		var goPath = os.Getenv("GOPATH")
		batFileFullName = goPath + string(os.PathSeparator) + "src" + string(os.PathSeparator) + "mydump" + string(os.PathSeparator) + batFileName
	}

	if !pathExists(batFileFullName) {
		log.Println("mydump.bat file not exist!" + batFileFullName)
	}

	if getFileMd5(batFileFullName) != "949a7193ffa240116c534fbae7dac90a" {
		log.Println("批处理文件内容被篡改过！")
		return
	}

	cmd := exec.Command(batFileFullName, batCmdParam)

	err := cmd.Run()

	if err != nil {
		log.Println("Execute 'mysqldump' Command failed: " + err.Error())
		return
	}
}
