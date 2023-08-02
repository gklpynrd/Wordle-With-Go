package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"github.com/fatih/color"
	"math/rand"
	"time"
)

var (
	//Amount of words in data
	listSize int = 1000

	//amount of guesses
	guesses int = 6

	//wordsize (there is only 5 word data and the file selection is hardcoded)
	wordsize int = 5

	//Win condition initialized as false
	won bool = false

	//For storing each guess words score to compare with secret word
	status [5]int

	//background assigns for coloring letters
	Greenbackground = color.New(color.BgGreen)
	Yellowbackground = color.New(color.BgYellow)
	Redbackground = color.New(color.BgRed)

	//If set true, selected random word will appear at the start
	test = true
)

func main() {
	
	//Greeter
	Greenbackground.Println("THIS IS WORDLE")
	fmt.Println("You have 6 tries to guess the 5-letter word I'm thinking of")

	//Picking a random number referencing time (couldn't find a better way)
	rand.NewSource(time.Now().Unix())
	n := rand.Int() % listSize

	//File opener
	file, err := os.Open("5.txt")

	//Error handler
	if err != nil {
		log.Fatal(err)
	}

	//Assigning a scanner using bufio for scanning and iterating through txt file
	Scanner := bufio.NewScanner(file)
	Scanner.Split(bufio.ScanWords)

	//Will break the function at selected random number
	var breaker int = 0

	//will be the main word that user has to find
	var word string

	//Iterating through txt file, returning and breaking at random number we picked earlier
	for Scanner.Scan() {
		breaker++
		if breaker == n{
			word = Scanner.Text()
			break
		}
	}

	//Hinting word for test
	if test{
		fmt.Println(word)
	}
	
	//Error handler for scanner
	if err := Scanner.Err(); err != nil {
		log.Fatal(err)
	}
	
	//Main game loop
	for j := 0; j < guesses; j++ {

		//Asking for guess and making sure its correct
		var guess string = getGuess()

		//Status resetter (has to reset for each guess)
		for i := 0; i < wordsize; i++ {
			status[i] = 0;
		}

		//Score counter
		var score int
		scoreCounter(guess, word, &score);
		
		//Printing the guess
		printGuess(guess ,j)
		
		//Max score for ech letter is 2, max score for word is 2*5
		if score == 10{
			won = true
			break
		}
	}
	
	//Results
	if won {
		fmt.Println("You Won")
	}else{
		fmt.Println("Word was :", word)
	}

	//For stopping before exit immediately
	fmt.Print("Press 'Enter' to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Scanln()
}


func getGuess() string {
	a := ""
	fmt.Print("your guess: ")
	fmt.Scan(&a)
	if len(a) != 5 {
		a = getGuess()
	}
	return a
}


//Scoring each letter for accuracy and storing score in an array 
func scoreCounter(guess string, word string, score *int) {
	
	for b := 0; b < wordsize; b++ {
		if guess[b] == word[b] {
			//Local var has to point value
			*score += 2
			//Global var so didnt have to use pointer
			status[b] = 2
		}else {
			for a := 0; a < wordsize; a++ {
				if guess[b] == word[a]{
					*score += 1
					status[b] = 1
				}
			}
		}
	}
}


//Printing every letter and coloring them based on scores
func printGuess(guess string, j int){
	fmt.Printf("Guess %v: ", j + 1)
		for i := 0; i < wordsize; i++ {
			if status[i] == 2{
				Greenbackground.Printf("%c",guess[i])
			}else if status[i] == 1 {
				Yellowbackground.Printf("%c",guess[i])
			}else{
				Redbackground.Printf("%c", guess[i])
			}
		}
		fmt.Println("")
}