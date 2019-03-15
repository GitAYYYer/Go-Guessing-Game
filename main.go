import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

//First index determines which superhero, second index determines what property.
//For 2nd index, "Name Of Hero", "Hint1|Hint2|Hint3|Hint4|Hint5, etc.", "Possible Answers", 
categories := [2]string {
	"SUPERHEROES", "WAIFUS"
}

superHeroes := [4][3]string	{
	{"SuperHeroes", "Descriptions", "Possible Answers"}
	{"Spider-Man", "Great Responsibility|Dead Uncle|'He's a Menace!'|Red, Blue and Black|Friendly Neighbourhood|Queens|Marvel", "Spider-Man|Spiderman|Peter Parker"},
	{"The Flash", "Red|Fast|Speedforce|DC|Forensic Scientist", "The Flash|Flash|Barry Allen|BarryAllen|Barry-Allen"},
	{"Batman", "World's Greatest Detective|Bats|League of Shadows|Utility Belt|DC|Billionaire", "Batman|Bruce Wayne|Bruce-Wayne|BruceWayne"},
}
//make sure to trim their answer so that 'Bat man' and 'Batman' will equal.

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi, welcome to Duc's 'Go Guessing Game'! The available categories to play are:")
	printCategories()
	fmt.Println("Type in a category below and press enter to play.")
	userCategory, error := reader.ReadString('\n')
	scanUserCategory(userCategory)
	fmt.Println("End of game.")
}

func printCategories() {
	for i := 0; i < len(categories); i++	{
		fmt.Println(" - " + categories[i])
	}
}

func scanUserCategory(userCategory)	{
	switch(string.toUpper(userCategory)	{
	case ("SUPERHEROES"):
		playSuperhero()
		break

	case ("WAIFUS"):
		playWaifu()
		break
	}
}

func playSuperhero() {
	fmt.Println("Playing as Superhero")
}

func playWaifu {
	fmt.Println("Playing as Waifu")
}