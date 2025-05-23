package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

var (
	AlphabetsSmallCased string = "abcdefghijklmnopqrstuvwxyz"
	AlphabetsUpperCased string = strings.ToUpper(AlphabetsSmallCased)
	Numbers             string = "1234567890"
	SpecialCharacters   string = "!!!!@#$%^*()_+=-"
)

func GeneratePassword(passwordLength int) string {
	var password []byte

	possibleTypes := []string{
		AlphabetsSmallCased,
		AlphabetsUpperCased,
		Numbers,
		SpecialCharacters,
	}

	for i := 0; i < passwordLength; i++ {
		var characterTypeIndex int = rand.Intn(len(possibleTypes))
		var selectedType string = possibleTypes[characterTypeIndex]
		var element byte = selectedType[rand.Intn(len(selectedType))]
		password = append(password, element)
	}

	return string(password)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("User-Error: Pass a Integer Length for auto-generated password")
		os.Exit(20)
	}

	passwordLength, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println("User-Error: Incorrect type passed Required value is string")
		os.Exit(30)
	}

	password := GeneratePassword(passwordLength)
	fmt.Println(password)

}
