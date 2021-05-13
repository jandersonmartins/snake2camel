package main

import (
	"fmt"
	"github.com/jandersonmartins/snake2camel"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := os.Args[1]
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalf("file %q doesn't exist", filePath)
	}
	out := snake2camel.ConvertToCamel(string(file))
	fmt.Fprint(os.Stdout, out)
}
