package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Print("Starting app!")
	output := flag.Bool("output", false, "Should there be output?")
	flag.Parse()
	fmt.Print(*output)
}
