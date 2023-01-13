package main

import (
	"bufio"
	"log"
	"os"

	"github.com/atotto/clipboard"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	buffer := ""
	for scanner.Scan() {
		buffer += scanner.Text()
	}
	if scanner.Err() != nil {
		log.Fatalf("Error occurred: %s", scanner.Err().Error())
	}

	if err := clipboard.WriteAll(buffer); err != nil {
		log.Fatalf("Error occurred: %s\n\n[Previous Output]:\n%s", err.Error(), buffer)
	}
}
