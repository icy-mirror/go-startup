package setting

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

var reader *bufio.Reader
var iniMap map[string]string

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
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if !strings.Contains(str, "=") || strings.HasPrefix(str, "#") {
			continue
		}
		temp := strings.Split(str, "=")
		key := temp[0]
		val := temp[1]
		if iniMap == nil {
			iniMap = make(map[string]string, 1)
		}
		iniMap[key] = val
	}

	fmt.Println("ini ok")
}

func GetIniVal(key string) string {
	return iniMap[key]
}
