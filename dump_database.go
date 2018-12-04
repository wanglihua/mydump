package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"
)

func dumpDatabase() {

	var batFileName = "mydump.bat"
	var batFileFullName = workDir + string(os.PathSeparator) + batFileName

	if !pathExists(batFileFullName) {
		var goPath = os.Getenv("GOPATH")
		batFileFullName = goPath + string(os.PathSeparator) + "src" + string(os.PathSeparator) + "mydump" + string(os.PathSeparator) + batFileName
	}

	if !pathExists(batFileFullName) {
		log.Fatal("mydump.bat file not exist!")
	}

	cmd := exec.Command(batFileFullName, time.Now().Format("20060102150405"))

	err := cmd.Run()

	if err != nil {
		fmt.Println("Execute 'mysqldump' Command failed: " + err.Error())
		return
	}
}
