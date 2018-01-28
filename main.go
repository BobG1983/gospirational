package main

import (
	"fmt"
	"log"

	"github.com/ShriekBob/goismatic"
)

func main() {
	quote, err := goismatic.Get(goismatic.English)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(quote)
}
