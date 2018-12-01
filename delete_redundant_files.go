package main

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

func deleteRedundantFiles() {

	var backupDir = iniFile.Section("").Key("backupDir").String()
	var databaseNameList = strings.Split(iniFile.Section("").Key("databases").String(), ",")
	fileCountKeep, err := iniFile.Section("").Key("fileCountKeep").Int()
	if err != nil {
		log.Fatal("fileCountKeep config missing!")
	}

	for _, databaseName := range databaseNameList {
		files, _ := filepath.Glob(backupDir + string(os.PathSeparator) + databaseName + "_*.zip")

		var fileNameList = FileNameList(files)
		sort.Sort(sort.Reverse(fileNameList))

		files = []string(fileNameList)

		if len(files) > fileCountKeep {
			var fileListToDelete = files[6:]
			for _, fileToDelete := range fileListToDelete {
				os.Remove(fileToDelete)
			}
		}
	}
}

type FileNameList []string

func (fileNameList FileNameList) Len() int {
	return len(fileNameList)
}

func (fileNameList FileNameList) Less(i, j int) bool {
	return fileNameList[i] < fileNameList[j]
}

func (fileNameList FileNameList) Swap(i, j int) {
	fileNameList[i], fileNameList[j] = fileNameList[j], fileNameList[i]
}
