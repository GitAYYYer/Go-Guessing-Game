package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var categories = [2]string {
	"SUPERHEROES", "WAIFUS",
}
var superHeroes = [4][3]string {
	{"SuperHeroes", "Descriptions", "Possible Answers"},
	{"Spider-Man", "Great Responsibility|Dead Uncle|'He's a Menace!'|Red, Blue and Black|Friendly Neighbourhood|Queens|Marvel", "Spider-Man|Spiderman|Peter Parker"},
	{"The Flash", "Red|Fast|Speedforce|DC|Forensic Scientist", "The Flash|Flash|Barry Allen|BarryAllen|Barry-Allen"},
	{"Batman", "World's Greatest Detective|Bats|League of Shadows|Utility Belt|DC|Billionaire", "Batman|Bruce Wayne|Bruce-Wayne|BruceWayne"},
}
//make sure to trim their answer so that 'Bat man' and 'Batman' will equal.
//First index determines which superhero, second index determines what property.
//For 2nd index, "Name Of Hero", "Hint1|Hint2|Hint3|Hint4|Hint5, etc.", "Possible Answers", 

func main() {
	categories := [2]string {
		"SUPERHEROES", "WAIFUS" }
	fmt.Println("Hi, welcome to Duc's 'Go Guessing Game'! The available categories to play are:")
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

func playSuperhero() {
	fmt.Println("Playing as Superhero")
}

func playWaifu() {
	fmt.Println("Playing as Waifu")
}