package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/bamchoh/json_parser/lexer"
	"github.com/bamchoh/json_parser/parser"
)

func load(filename string) (string, error) {
	fullpath, err := filepath.Abs(filename)
	if err != nil {
		return "", err
	}
	fp, err := os.Open(fullpath)
	if err != nil {
		return "", err
	}
	defer fp.Close()

	script, err := ioutil.ReadAll(fp)
	if err != nil {
		return "", err
	}

	return string(script), nil
}

func dump(data interface{}) string {
	var out bytes.Buffer
	switch data.(type) {
	case map[string]interface{}:
		fmt.Fprintf(&out, "{\n")
		i := 0
		for k, v := range data.(map[string]interface{}) {
			if i != 0 {
				fmt.Fprint(&out, ",\n")
			}
			fmt.Fprintf(&out, "%2s%q: ", "", k)
			lines := strings.Split(dump(v), "\n")
			for j := 0; j < len(lines); j++ {
				if j != 0 {
					fmt.Fprint(&out, "\n")
				}
				if j == 0 {
					fmt.Fprint(&out, lines[j])
				} else {
					fmt.Fprintf(&out, "%2s%v", "", lines[j])
				}
			}
			i++
		}
		fmt.Fprintf(&out, "\n}")
	case []interface{}:
		fmt.Fprintf(&out, "[\n")
		for i, v := range data.([]interface{}) {
			if i != 0 {
				fmt.Fprint(&out, ",\n")
			}
			lines := strings.Split(dump(v), "\n")
			for j := 0; j < len(lines); j++ {
				if j != 0 {
					fmt.Fprint(&out, "\n")
				}
				if j == 0 {
					fmt.Fprintf(&out, "%2s%v", "", lines[j])
				} else {
					fmt.Fprintf(&out, "%2s%v", "", lines[j])
				}
			}
		}
		fmt.Fprintf(&out, "\n]")
	case string:
		fmt.Fprintf(&out, "%q", data)
	default:
		fmt.Fprintf(&out, "%v", data)
	}
	return out.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please specify JSON text file")
		os.Exit(-1)
	}

	text, err := load(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

	l := lexer.New(text)
	p := parser.New(l)
	data := p.Parse()
	fmt.Println(dump(data))
}
