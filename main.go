package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"math/rand"
)

var categories = [2]string {
	"SUPERHEROES", "WAIFUS",
}
var superHeroes = [4][3]string {
	{"Spider-Man", "Great Responsibility|Dead Uncle|'He's a Menace!'|Red Blue and Black|Friendly Neighbourhood...|Queens|Marvel", "Spider-Man|Spiderman|Peter Parker"},
	{"The Flash", "Red|Fast|Speedforce|DC|Forensic Scientist", "The Flash|Flash|Barry Allen|BarryAllen|Barry-Allen"},
	{"Batman", "World's Greatest Detective|Bats|League of Shadows|Utility Belt|DC|Billionaire", "Batman|Bruce Wayne|Bruce-Wayne|BruceWayne"},
}
//make sure to trim their answer so that 'Bat man' and 'Batman' will equal.
//First index determines which superhero, second index determines what property.
//For 2nd array, "Name Of Hero", "Hint1|Hint2|Hint3|Hint4|Hint5, etc.", "Possible Answers", 

func main() {
	categories := [2]string {
		"SUPERHEROES", "WAIFUS" }
	fmt.Println("Hi, welcome to Duc's 'Go Guessing Game'! You have 1 minute to guess the character from 3 short descriptions.\nThe available categories to play are:")
	printCategories(categories)
	fmt.Println("Type in a category below and press enter to play.")
	scanUserCategory()
	fmt.Println("End of game.")
}

func printCategories(categories [2]string) {
	for i := 0; i < len(categories); i++	{
		fmt.Println(" - " + categories[i])
	}
}

func scanUserCategory()	{
	reader := bufio.NewReader(os.Stdin)
	userCategory, error := reader.ReadString('\n')
	userCategoryUp := strings.TrimRight(strings.ToUpper(userCategory), "\r\n") 	//all readers have \r\n at end (newline). Need to trim the ToUpper string. \r is needed for windows machines.
	if (error == nil) {
		switch
		{
		case userCategoryUp == "SUPERHEROES":
			playSuperhero()
			break

		case userCategoryUp == "WAIFUS":
			playWaifu()
			break
		}
	} else {
		fmt.Println("Oops! Try typing a valid category.")
		scanUserCategory()
	}
}

func countdown() {
	countdownTime := 3
	ticker := time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <- ticker.C:
				if (countdownTime == 0)	{
					//do nothing, allow a second of no countdown to make console output look smoother
				} else {
				fmt.Println(countdownTime)
				countdownTime--
				}
			}
		}
	}()

	time.Sleep(4 * time.Second)	//TODO: sometimes prints 1, sometimes doesn't.
	ticker.Stop()
	return
}

func playSuperhero() {
	var remainingIndexValues = []int {}
	var correctAnswers = 0
	rand.Seed(time.Now().UnixNano())	//without this, the environment won't actually generate random numbers each game, and it will generate the same ones instead.

	for i := 0; i < len(superHeroes) - 1; i++	{
		remainingIndexValues = append(remainingIndexValues, i)
	}

	fmt.Println("Playing Category: Superheroes")
	countdown()
	
	ticker := time.NewTicker(time.Second)

	go func() {
		for i := 0; i < len(superHeroes); i++	{
			randomIndexValue := remainingIndexValues[rand.Intn(len(remainingIndexValues))]	//randomIndex can equal 0, 1, 2, 3. Purpose is to remove this from remainingIndexValues array after it is used.
			descriptiveWords := strings.Split(superHeroes[randomIndexValue][1], "|")
			descriptiveWordsPrint := [3]string{} //array that holds 3 strings that are each different from descriptiveWords array. Then, it prints out to console.
			descriptiveWordsCounter := 0
			for {
				line := descriptiveWords[rand.Intn(len(descriptiveWords))]
				for x := 0; x < len(descriptiveWordsPrint); x++	{
					if (strings.Compare(descriptiveWordsPrint[x], line) == 0)	{
						break	//breaks out of this loop, and the infinite for loop will continue.
					}
					if (x == 2)	{	//the line is unique, and can be added to descriptiveWordsPrint
						descriptiveWordsPrint[descriptiveWordsCounter] = line
						descriptiveWordsCounter++
						break	//break out of x loop back to infinite loop, so that line can reset to a new line
					}
				}

				if (descriptiveWordsCounter == 3) {	//break out of infinite loop, since I have 3 descriptive words.
					break
				}
			}

			for x := range descriptiveWordsPrint {
				if (x != 2)	{
					fmt.Print(descriptiveWordsPrint[x] + ", ")
				} else {
					fmt.Println(descriptiveWordsPrint[x])
				}
			}
			reader := bufio.NewReader(os.Stdin)
			userAnswer, error := reader.ReadString('\n')
			if (error == nil)	{
				if (userAnswer == "correct\r\n")	{	//TODO: obviously replace correct with if the userAnswer matches answers from superheroes array.
					fmt.Println(userAnswer)
					for x := 0; x < len(remainingIndexValues); x++ {	//deletes randomIndexValue number from the remainingIndexValues.
						if (remainingIndexValues[x] == randomIndexValue)	{
							remainingIndexValues[x] = remainingIndexValues[len(remainingIndexValues) - 1]
							remainingIndexValues[len(remainingIndexValues) - 1] = -1	//-1 will ensure randomIndexValue cannot choose this.
							remainingIndexValues = remainingIndexValues[:len(remainingIndexValues) - 1]
						}
					}
					fmt.Println(remainingIndexValues)
				}
			}
		}
	}()

	time.Sleep(3 * time.Second)	//leave at 5-10 seconds initially for debug purposes.
	ticker.Stop()

	if (correctAnswers == len(superHeroes))	{
		fmt.Println("You got all of them correct! Good job!")
	} else {
		fmt.Println("You got ", correctAnswers, " correct. Better luck next time!")
	}
}

func playWaifu() {
	fmt.Println("Playing as Waifu")
}