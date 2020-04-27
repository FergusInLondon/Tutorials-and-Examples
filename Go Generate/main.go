package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	statusCode, _ := strconv.Atoi(os.Args[1])
	fmt.Printf("Received Status '%d'.\n", statusCode)

	statusMessage, err := StatusCode(statusCode).Message()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Message: '%s'.\n", statusMessage)
}
