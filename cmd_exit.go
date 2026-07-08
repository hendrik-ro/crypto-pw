package main

import (
	"fmt"
	"os"
)

func cmdExit(_ *config, _ []string) error {
	fmt.Println("exiting pw...")
	os.Exit(0)
	return nil
}
