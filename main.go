package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		name := os.Args[0]
		fmt.Printf("Usage: %v \"expression\"\n", name)
		os.Exit(1)
	}
	expression := os.Args[1]

	inst := &URN{Buffer: expression}
	inst.Init()
	if err := inst.Parse(); err != nil {
		log.Fatal(err)
	}

	inst.PrettyPrintSyntaxTree(inst.Buffer)
}
