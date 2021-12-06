package main

import (
	"bufio"
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

func generateFileData(fileName string) (string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return "", err
	}

	divisions := "map[string]string{"
	buff := bufio.NewReader(file)
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
		lineScanner.Scan()

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
	dir := "../data/mca/"
	names := []int{201904}

	gocode := `package gb2260
var (
	Divisions map[string]map[string]string
)

func init(){
	Divisions = map[string]map[string]string{
		%s
	}
}
`

	code := ""

	for _, n := range names {
		fileName := dir + fmt.Sprintf("%d.tsv", n)

		data, err := generateFileData(fileName)
		if err != nil {
			log.Fatal(err)
			return
		}

		code += fmt.Sprintf("\"%d\": %s,\n", n, data)
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
