package main

import (
	"embed"
	"fmt"
	"log"
	"strings"
)

//go:embed dictionary.txt
var dictionary embed.FS
const target ="Crunch Berry"

func main() {
	input, err := dictionary.ReadFile("dictionary.txt")
	if err != nil {
		log.Println("unable to read dictionary file", err)
		return
	}

	lines := strings.Split(string(input), "\n")
	normalizedTarget := strings.ToLower(strings.ReplaceAll(target, " ", ""))
	var foundWords []string
	wordLoop:
	for _, line := range lines {
		// not allowing single letter "words" - we have standards around here
		if len(line) == 1 {
			continue wordLoop
		}

		movingTarget := normalizedTarget
		letters := strings.Split(strings.ToLower(line), "")
		for i:=0; i<len(letters); i++ {
			if strings.Contains(movingTarget, letters[i]) {
				movingTarget = strings.Replace(movingTarget, letters[i], "", 1)
			} else {
				continue wordLoop
			}
		}

		if strings.Join(letters, "") != "crunch" && strings.Join(letters, "") != "berry" {
			foundWords = append(foundWords, strings.Join(letters, ""))
		}
	}

	fmt.Printf("found %v words in \"%v\":\n", len(foundWords), target)
	for i:=0; i<len(foundWords); i++ {
		fmt.Print(fmt.Sprintf("%v ", foundWords[i]))
		if i > 1 && i % 8 == 0 {
			fmt.Print("\n")
		}
	}

	fmt.Print("\n")
	fmt.Scanf("exit")
}
