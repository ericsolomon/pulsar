package main

import (
	"fmt"
	"os"

	"github.com/ericsolomon/pulsar/repl"
)

func main() {
	fmt.Printf("Pulsar\n")
	repl.Start(os.Stdin, os.Stdout)
}
