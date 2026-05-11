package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

//Global Variables

var lettersArray []string
var askForLetter bool = true
var emptySpace int32 = 500
var tried []string

var reRolls int32 = 0

//**

func main() {
	for i := 64; i < 123; i++ {
		fmt.Println(" \n")
	}
	getLetter()
	askForTries()
	guessWord()
	color.Blue("Finished!")
	fmt.Println(strconv.Itoa(int(reRolls)))

}

func getLetter() {
	for askForLetter {

		if len(lettersArray) > 0 {
			var compund string = fmt.Sprintf("Letters: %s \n================================", strings.Join(lettersArray, ""))
			color.Red(compund)
		}

		color.Blue("Enter a letter:")
		var userLetter string
		fmt.Scanln(&userLetter)

		if userLetter != "end" {
			lettersArray = append(lettersArray, userLetter)

		} else {
			askForLetter = false
		}

	}

}

func askForTries() {
	color.Blue("Number of tries (default 500):\n================")
	fmt.Scanln(&emptySpace)

}

func guessWord() {

	var uniqueIndex int32 = 0

	//Repeat the process of guessing until we run out of empty spaces.
	for i := 0; i < int(emptySpace); i++ {

		//No fucking clue how this works, but i made it work... somehow
		rand.Shuffle(len(lettersArray), func(i, j int) {
			lettersArray[i], lettersArray[j] = lettersArray[j], lettersArray[i]

			var curentGuess string = strings.Join(lettersArray, "")

			if !slices.Contains(tried, curentGuess) {
				tried = append(tried, curentGuess)

				uniqueIndex++

				guess := englishDictionary(curentGuess)
				if guess != "" {
					var compund string = fmt.Sprintf("# %d | GUESSED: %s \n================================", uniqueIndex, strings.Join(lettersArray, ""))
					color.Green(compund)
					return
				}

			} else {
				//If the random word was already guessed, roll it again.
				i--
				reRolls++
			}

		})

	}

}

func englishDictionary(generatedWord string) string {
	dictionary := make(map[string]bool)

	var returnString string

	file, err := os.Open("words_alpha.txt")
	if err != nil {
		color.Red("Error opening dictionary file:", err)
		return ""
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dictionary[scanner.Text()] = true
	}

	if err := scanner.Err(); err != nil {
		color.Red("Error reading dictionary file:", err)
	}

	if dictionary[generatedWord] {
		returnString = generatedWord
	}

	return returnString
}
