package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	// file read
	f, err := os.ReadFile("/Users/setohiroyuki/work/docker/gopher/image-modify-format/image/mountain-g5a8babad2_1920.jpg")
	if err != nil {
		log.Fatal(err)

	}
	fmt.Print(f)

	// modify
}
