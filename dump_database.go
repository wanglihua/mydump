package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func dumpDatabase() {
	var mysqldumpCmd = iniFile.Section("").Key("mysqldump").String()
	var user = iniFile.Section("").Key("user").String()
	var password = iniFile.Section("").Key("password").String()
	var backupDir = iniFile.Section("").Key("backupDir").String()
	var databaseNameList = strings.Split(iniFile.Section("").Key("databases").String(), ",")

	for _, databaseName := range databaseNameList {
		cmd := exec.Command(mysqldumpCmd, "--user="+user, "--password="+password, "--databases", databaseName)
		stdout, _ := cmd.StdoutPipe()

		err := cmd.Start()

		zipFile, err := os.Create(backupDir + string(os.PathSeparator) + databaseName + "_" + time.Now().Format("20060102150405") + ".zip")
		defer zipFile.Close()

		var zipWriter = zip.NewWriter(zipFile)
		defer zipWriter.Close()

		fileWriter, err := zipWriter.Create(databaseName + ".sql")

		bytes, err := ioutil.ReadAll(stdout)

		_, err = fileWriter.Write(bytes)

		stdout.Close()

		if err := cmd.Wait(); err != nil {
			fmt.Println("Execute failed when Wait:" + err.Error())
			return
		}

		if err != nil {
			fmt.Println("Execute 'mysqldump' Command failed: " + err.Error())
			return
		}
	}
}
