package main

import (
	"fmt"
	"testing"
)

func TestGetFileMd5(t *testing.T) {
	var fileMd5 = getFileMd5(`C:\projects\mydump\src\mydump\mydump.bat`)
	fmt.Println(fileMd5)
}
