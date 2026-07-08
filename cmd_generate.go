package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

func cmdGenerate(cfg *config, args []string) error {
	length := cfg.Length
	if len(args) > 0 {
		l, err := strconv.Atoi(args[0])
		if err != nil {
			return err
		}
		if l > 9 && l < 26 {
			length = l
		} else {
			fmt.Println("recommended password length: 10-25 characters")
			fmt.Printf("%d outside of recommended range - generating %d characters\n", l, cfg.Length)
		}
	}
	pw := generate(cfg, length)
	fmt.Printf("Generated password: %s\n", pw)
	return nil
}

func generate(cfg *config, n int) string {
	chars := cfg.Characters
	builder := strings.Builder{}
	i := 0
	for range n {
		char := chars[rand.Intn(len(chars))]
		if i == 0 {
			char = chars[rand.Intn(len(chars)-(len(cfg.Numbers)+len(cfg.Specials)))+len(cfg.Numbers)]
		}
		builder.WriteString(char)
		i++
	}
	pw := builder.String()
	specialFound := checkSpecials(cfg, pw)
	if !specialFound {
		// fmt.Println("weak password - re-generating...")
		pw = generate(cfg, n)
	}
	return pw
}
