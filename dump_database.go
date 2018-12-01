package main

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
)

func dumpDatabase() { // cmd := exec.Command("c:\\mysql\\bin\\mysqldump", "--user=root", "--password=root", "--databases", "eems", ">", "d:\\eems.sql")
	// cmd := exec.Command("c:\\mysql\\bin\\mysqldump", "--user=root --password=root --databases eems > d:\\eems.sql")
	cmd := exec.Command("c:\\mysql\\bin\\mysqldump", "--user=root", "--password=root", "--databases", "eems")
	stdout, _ := cmd.StdoutPipe()

	err := cmd.Start()

	zipFile, err := os.Create("d:\\eems.zip")
	defer zipFile.Close()

	var zipWriter = zip.NewWriter(zipFile)
	defer zipWriter.Close()

	fileWriter, err := zipWriter.Create("eems.sql")

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
