package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func startInterpretation(fileName string) {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	GetTokensFromSource(string(data))

}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("usage: go-lox <script>")
		return
	}

	startInterpretation(args[1])
}
