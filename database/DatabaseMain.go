package main

import (
	"fmt"

	"./setting"
)

func main() {
	setting.InitializeIni()
	fmt.Println("okok")
	setting.GetIniVal("enter mission!")
}
