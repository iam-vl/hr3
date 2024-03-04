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

func panicr(err error) {
	if err != nil {
		panic(err)
	}
}

func line() {
	fmt.Println("=====================")
}
