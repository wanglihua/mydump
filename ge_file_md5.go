package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
)

func getFileMd5(fileFullName string) string {
	f, err := os.Open(fileFullName)
	if err != nil {
		log.Println("Open File Error: ", err)
		return ""
	}

	defer f.Close()

	md5hash := md5.New()
	if _, err := io.Copy(md5hash, f); err != nil {
		log.Println("Copy File Content Error: ", err)
		return ""
	}

	return fmt.Sprintf("%x", md5hash.Sum(nil))
}
