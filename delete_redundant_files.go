package main

import (
	"os"
	"path/filepath"
	"sort"
)

func deleteRedundantFiles() {
	files, _ := filepath.Glob("d:\\database_backup\\eems_*.zip")

	var fileNameList = FileNameList(files)
	sort.Sort(sort.Reverse(fileNameList))

	files = []string(fileNameList)

	if len(files) > 6 {
		var fileListToDelete = files[6:]
		for _, fileToDelete := range fileListToDelete {
			os.Remove(fileToDelete)
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
