package main

import (
	"fmt"
	"os"
	"slices"
	"strings"

	"golang.org/x/term"
)

func cmdStrength(cfg *config, _ []string) error {
	// Prompt masked password input
	fmt.Print("Enter password > ")
	password, err := term.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
	}
	args := strings.Fields(string(password))
	for i := 0; i < len(args[0]); i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
	if len(args) == 0 {
		return fmt.Errorf("no password provided")
	}
	pw := args[0]
	valErr := validate(cfg, pw)
	if valErr != nil {
		return valErr
	}
	return nil
}

func validate(cfg *config, pw string) error {
	if slices.Contains(cfg.CommonPWs, pw) {
		return fmt.Errorf("extremely weak - password is too common")
	}
	if len(pw) > 15 {
		if checkNumbers(cfg, pw) || checkSpecials(cfg, pw) {
			fmt.Println("very strong")
		} else {
			fmt.Println("strong")
		}
		return nil
	}
	if len(pw) >= 10 {
		if checkNumbers(cfg, pw) && checkSpecials(cfg, pw) {
			fmt.Println("strong")
		} else if checkNumbers(cfg, pw) {
			fmt.Println("moderate - concider adding special characters")
		} else if checkSpecials(cfg, pw) {
			fmt.Println("moderate - concider adding numbers")
		} else {
			fmt.Println("weak - concider adding special characters and/or numbers")
		}
		return nil
	}
	if len(pw) < 10 {
		if checkNumbers(cfg, pw) && checkSpecials(cfg, pw) {
			fmt.Println("weak - concider a longer password")
		} else {
			fmt.Println("extremely weak - concider adding special characters and/or numbers")
		}
		return nil
	}
	return fmt.Errorf("could not validate password")
}

func checkSpecials(cfg *config, pw string) bool {
	found := false
	i := 0
	for _, s := range cfg.Specials {
		for _, c := range pw {
			if string(c) == s {
				found = true
			}
		}
		i++
	}
	return found
}

func checkNumbers(cfg *config, pw string) bool {
	found := false
	i := 0
	for _, n := range cfg.Numbers {
		for _, c := range pw {
			if string(c) == n {
				found = true
			}
		}
		i++
	}
	return found
}
