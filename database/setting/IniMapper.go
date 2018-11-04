package setting

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

var reader *bufio.Reader

/**
  Load Setting and parse
  @author D.Tamura
*/

func InitializeIni() {
	iniFile, err := os.Open("Setting.ini")
	if err != nil {
		fmt.Errorf("Error! ini file could not load.")
		os.Exit(50)
	}
	reader = bufio.NewReader(iniFile)

	fmt.Println("ok")
}

func GetIniVal(key string) string {
	fmt.Println("key: " + key)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Println(str)
	}
	return "hoge"
}
