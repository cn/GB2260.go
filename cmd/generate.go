package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
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

func generateFileData(year int, fileName string) (string, error) {
	content, err := ioutil.ReadFile(fileName)
	if err != nil {
		return "", err
	}

	divisions := "map[string]string{"

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

		divisions += fmt.Sprintf(`"%s": "%s",`, code, name) + "\n"
	}
	divisions += "}"

	return divisions, nil
}

// before use it, you should use go build and then execute the generate (or generate.exe)
func main() {
	dir := "../data/"
	years := []int{2006, 2007, 2008, 2009, 2010, 2011, 2012, 2013, 2014}

	gocode := `package gb2260
var (
	divisions map[string]map[string]string
)

func init(){
	divisions = map[string]map[string]string{
		%s
	}
}
`

	code := ""

	for _, y := range years {
		fileName := dir + fmt.Sprintf("GB2260-%d.txt", y)

		data, err := generateFileData(y, fileName)
		if err != nil {
			log.Fatal(err)
			return
		}

		code += fmt.Sprintf("\"%d\": %s,\n", y, data)
	}

	gocode = fmt.Sprintf(gocode, code)
	err := ioutil.WriteFile("../data.go", []byte(gocode), os.ModePerm)
	if err != nil {
		log.Fatal(err)
		return
	}

	cmd := exec.Command("go", "fmt", "../data.go")
	cmd.Start()

	return
}
