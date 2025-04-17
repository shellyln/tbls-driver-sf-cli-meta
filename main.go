package main

import (
	"os"

	"github.com/shellyln/tbls-driver-sf-cli-meta/pkg/driver"
)

func main() {
	err := driver.Run()
	if err != nil {
		panic(err) // TODO:
	}

	os.Exit(0)
}
