package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	xj "github.com/basgys/goxml2json"
)

// convert YAML from stdin to JSON
func main() {
	stdin, err := io.ReadAll(os.Stdin)

	if err != nil {
		panic(err)
	}
	if len(stdin) < 1 {
		fmt.Printf("Usage:\n\txml_from_stdin | %s\n", os.Args[0])
		return
	}
	xml := strings.NewReader(string(stdin))
	json, err := xj.Convert(xml)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}

	fmt.Println(json.String())

}
