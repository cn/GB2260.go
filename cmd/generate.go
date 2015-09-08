package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func getNameSuffix(currentSource string) string {
	ext := path.Ext(currentSource)
	pathData := currentSource[0 : len(currentSource)-len(ext)]
	splits := strings.SplitN(pathData, "-", 2)
	if len(splits) == 2 {
		return splits[1]
	}

	return ""
}

func main() {
	length := len(os.Args)
	if length < 3 {
		fmt.Println(fmt.Sprintf("Usage %s [A] [B] ... [DESTINATION]", os.Args[0]))
		return
	}

	source := os.Args[1 : length-1]
	destination := os.Args[length-1]

	lineData := make(map[string]string)

	for _, currentSource := range source {
		suffix := getNameSuffix(currentSource)
		content, err := ioutil.ReadFile(currentSource)
		if err != nil {
			log.Fatal(err)
			return
		}

		buff := bytes.NewBuffer(content)

		for {
			line, err := buff.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					break
				}
			}

			lineScanner := bufio.NewScanner(strings.NewReader(line))
			lineScanner.Split(bufio.ScanWords)

			lineScanner.Scan()
			code := lineScanner.Text()

			lineScanner.Scan()
			name := lineScanner.Text()

			lineData[code] = fmt.Sprintf(`%s: Division{Code: "%s", Name: "%s", Year: "%s"},`, code, code, name, suffix)
		}

	}

	gocode := `package gb2260

var (
	divisions map[int]Division
)

func init(){
	divisions = map[int]Division{
%s
	}
}

`

	allLine := ""

	for _, value := range lineData {
		allLine += "\t\t" + value + "\n"
	}

	gocode = fmt.Sprintf(gocode, allLine)

	err := ioutil.WriteFile(destination, []byte(gocode), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}
	return
}
