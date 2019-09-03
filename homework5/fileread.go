package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("fileread.go")
	if err != nil {
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
