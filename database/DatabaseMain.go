package main

import (
	"fmt"

	"./setting"
)

func main() {
	setting.InitializeIni()
	fmt.Println(setting.GetIniVal("enter mission!"))
	fmt.Println(setting.GetIniVal("db-file-dir"))
}
