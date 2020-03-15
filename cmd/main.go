package main

import (
	"fmt"
	"log"

	"github.com/bitfield/yijing"
)

func main() {
	h, err := yijing.RandomHexagram()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d\n", h)
}
