package main

import (
	"fmt"
	"github.com/google/uuid"
	"os"
	"strconv"
)

func main() {
	var needGuidCount = 1

	if len(os.Args) > 1 {
		parsedInt, err := strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println("Error:", os.Args[1], "must be a number")
			os.Exit(-1)
		}
		needGuidCount = parsedInt
	}

	for i := 0; i < needGuidCount; i++ {
		outputUUID(generateUUID())
	}
}

func generateUUID() *string {
	result := uuid.New().String()
	return &result

}

func outputUUID(id *string) {
	fmt.Println(*id)
}
