package main

import (
	"fmt"
	"log"
)

func logf(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func line() {
	fmt.Println("=====================")
}
