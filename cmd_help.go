package main

import (
	"fmt"
	"sort"
)

func cmdHelp(_ *config, _ []string) error {
	commands := getCommands()
	keys := make([]string, 0, len(commands))
	for k := range commands {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Printf("    '%s' - %s\n", commands[key].name, commands[key].description)
	}
	return nil
}
