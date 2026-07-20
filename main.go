package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type config struct {
	Numbers    []string
	Alphabet   []string
	Specials   []string
	Characters []string
	Length     int
	CommonPWs  []string
}

var commonPWs = []string{
	"123456789",
	"qwerty",
	"password",
	"12345",
	"qwerty123",
	"1q2w3e",
	"12345678",
	"111111",
	"1234567890",
	"admin",
	"111111",
	"000000",
	"iloveyou",
	"letmein",
	"welcome",
	"monkey",
	"dragon",
	"123456789",
	"123123",
	"12345",
	"123456",
	"1234567",
	"12345678",
	"1234567890",
}

func loadConfig() config {
	numbers := "0123456789"
	alphabet := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialChars := "!@$#"
	numbersArray := strings.Split(numbers, "")
	alphabetArray := strings.Split(alphabet, "")
	specialsArray := strings.Split(specialChars, "")
	characters := slices.Concat(numbersArray, alphabetArray, specialsArray)
	cfg := config{
		Numbers:    numbersArray,
		Alphabet:   alphabetArray,
		Specials:   specialsArray,
		Characters: characters,
		Length:     15,
		CommonPWs:  commonPWs,
	}
	return cfg
}

func main() {
	fmt.Println("launching pw...")
	cfg := loadConfig()
	fmt.Println("type 'help' for available commands")
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("PW > ")
		reader.Scan()
		err := reader.Err()
		if err != nil {
			fmt.Println(err)
		}

		input := cleanInput(reader.Text())
		if len(input) == 0 {
			continue
		}

		cmd := input[0]
		args := input[1:]

		command, exists := getCommands()[cmd]
		if exists {
			err := command.callback(&cfg, args)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	input := strings.ToLower(text)
	fields := strings.Fields(input)
	return fields
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"check": {
			name:        "check",
			description: "Checks password for its strength",
			callback:    cmdStrength,
		},
		"exit": {
			name:        "exit",
			description: "Exits program",
			callback:    cmdExit,
		},
		"generate": {
			name:        "generate [n]",
			description: "Generates random password with (optional) length n",
			callback:    cmdGenerate,
		},
		"help": {
			name:        "help",
			description: "Shows available commands",
			callback:    cmdHelp,
		},
	}
}
