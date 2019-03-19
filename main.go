package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"time"
	"math/rand"
)

var categories = [1]string {
	"SUPERHEROES",
}
var superHeroes = [19][3]string {
	{"Spider-Man", "Great Responsibility|Dead Uncle|'He's a Menace!'|Red Blue and Black|Friendly Neighbourhood...|Queens", "SPIDERMAN|SPIDER-MAN|PETERPARKER"},
	{"The Flash", "Red and Yellow|Fast|Speedforce|Justice-League|Fastest Man Alive", "THE FLASH|FLASH|BARRYALLEN|BARRY-ALLEN"},
	{"Batman", "World's Greatest Detective|Bats|Utility Belt|Billionaire|Justice-League", "BATMAN|BRUCEWAYNE"},
	{"Captain America", "Shield|America|Blue Red White|Super Soldier|Man out of Time|The First Avenger", "CAPTAINAMERICA|STEVEROGERS"},
	{"Hulk", "SMASH!|Green|The Incredible...|Gamma Ray Exposure|'I'm always angry'", "HULK|THEHULK|BRUCEBANNER"},
	{"Superman", "Krypton|Blue and Red|Justice-League|Man of Steel|Kryptonite", "SUPERMAN|CLARKKENT|KAL-EL"},
	{"Green Lantern", "Green|Ring|Willpower|Hard-Light Construct|Lantern|Justice-League", "GREENLANTERN|GREEN-LANTERN|HALJORDAN"},
	{"Wolverine", "Healing Factor|Adamantium|Mutant|Claws|X-Men", "WOLVERINE|LOGAN|HUGHJACKMAN"},
	{"Thor", "God of Thunder|Mjolnir|Asgardian|Worthy|Avenger|Norse Mythology", "THOR"},
	{"Iron Man", "Avenger|Billionaire Playboy|Suit of Armor|Red and Yellow", "IRONMAN|IRON-MAN|TONYSTARK"},
	{"Wonder Woman", "The Amazonian Princess|Justice-League|Lasso of Truth", "WONDERWOMAN|DIANAPRINCE"},
	{"Ant Man", "Ants|Pym Particles|Shrink", "ANTMAN|ANT-MAN|SCOTTLANG"},
	{"Doctor Strange", "Doctor|Car Accident|Master of the Mystic Arts|Time Stone", "DR.STRANGE|DRSTRANGE|DOCTORSTRANGE|STEPHENSTRANGE"},
	{"Hawkeye", "Bow and Arrow|Useless|Avenger", "HAWKEYE|CLINT|CLINTBARTON"},
	{"Groot", "Guardians of the Galaxy|Plant|'I am...'", "GROOT"},
	{"Black Panther", "Wakanda|Vibranium|King|Panther", "BLACKPANTHER|T'CHALLA"},
	{"Deadpool", "More like Antihero|Healing Factor|Funny|Ryan Reynolds <3", "DEADPOOL|WADEWILSON|RYANREYNOLDS"},
	{"Professor X", "Wheelchair|X-Men|Mutant|Telepath", "PROFESSORX|PROFESSORXAVIER|CHARLESXAVIER|XAVIER"},
	{"Black Widow", "Red Hair (Usually)|Russian Spy|Avenger", "BLACKWIDOW|NATASHAROMANOFF"},
}

var waifus = [1][3]string {
	{"Rui Tachibana", ""},
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

	for i := 0; i < len(superHeroes); i++	{
		remainingIndexValues = append(remainingIndexValues, i)
	}

	fmt.Println("Playing Category: Superheroes")
	countdown()
	
	done := make(chan bool)

	go func() {
		for {
			randomIndexValue := remainingIndexValues[rand.Intn(len(remainingIndexValues))]	//randomIndex can equal 0, 1, 2, 3. Purpose is to remove this from remainingIndexValues array after it is used.
			descriptiveWords := strings.Split(superHeroes[randomIndexValue][1], "|")
			descriptiveWordsPrint := [3]string{} //array that holds 3 strings that are each different from descriptiveWords array. Then, it prints out to console.
			descriptiveWordsCounter := 0
			possibleAnswers := strings.Split(superHeroes[randomIndexValue][2], "|")

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
			var correct = false
			if (error == nil)	{
				for y := 0; y < len(possibleAnswers); y++	{
					if (strings.Replace(strings.TrimSpace(strings.ToUpper(userAnswer)), " ", "", -1) == possibleAnswers[y])	{	//strips literally all white space from userAnswer to match possible answer
						correct = true						
						for x := 0; x < len(remainingIndexValues); x++ {	//deletes randomIndexValue number from the remainingIndexValues.
							if (remainingIndexValues[x] == randomIndexValue)	{
								remainingIndexValues[x] = remainingIndexValues[len(remainingIndexValues) - 1]
								remainingIndexValues[len(remainingIndexValues) - 1] = -1	//-1 will ensure randomIndexValue cannot choose this.
								remainingIndexValues = remainingIndexValues[:len(remainingIndexValues) - 1]
								correctAnswers++
							}
						}
					}

					if (strings.TrimSpace(strings.ToUpper(userAnswer)) == "EXIT")	{
						done <- true
					}
				}

				if (correct) {
					fmt.Println("Correct!")
				} else {
					fmt.Println("Incorrect!")
				}
			}

			if (correctAnswers == len(superHeroes))	{
				fmt.Println("You got all", correctAnswers, "SuperHeroes correct! Good job!")
				done <- true
				break
			}
		}
	}()

	select {
	case <- done:
		break
	case <- time.After(60 * time.Second):
		fmt.Println("")
		break
	}

	if (correctAnswers != len(superHeroes))	{
		fmt.Println("You got", correctAnswers, "correct. Better luck next time!")
	}
}

func playWaifu() {	//put 5 descriptions for waifus
	fmt.Println("Playing as Waifu")
}