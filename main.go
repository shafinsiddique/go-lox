package main

import (
	"bufio"
	"fmt"
	"os"
)

func startInterpretation(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func main() {
	args := os.Args

	if len(args) != 2 {
		fmt.Printf("usage: go-lox <script>")
		return
	}

	startInterpretation(args[1])

}
