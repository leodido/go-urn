// SPDX-License-Identifier: MIT
//
// Copyright © 2019- Leonardo Di Donato <leodidonato@gmail.com>
package main

import (
	"log"
	"os"
	"strings"

	gosed "github.com/rwtodd/Go.Sed/sed"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("first (and only) argument (file path) is mandatory")
	}
	path := os.Args[1]

	sed, err := gosed.New(strings.NewReader(`/^\/\/line/d`))
	if err != nil {
		log.Fatal(err)
	}

	buf, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	res, err := sed.RunString(string(buf))
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile(path, []byte(res), 0); err != nil {
		log.Fatal(err)
	}
}
