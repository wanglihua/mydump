package main

import (
	"log"
	"os"
	"os/exec"
	"time"
)

func dumpDatabase() {

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

	cmd := exec.Command(batFileFullName, time.Now().Format("20060102150405"))

	err := cmd.Run()

	if err != nil {
		log.Println("Execute 'mysqldump' Command failed: " + err.Error())
		return
	}
}
