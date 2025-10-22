package main

import ("fmt")
import ("math/rand/v2")

// main function
func main() {
	fmt.Println("Hello World! Welcome to Tamayo") // hello and welcome message print
	fmt.Println("An advanced interesting number guessing game!") // small description about this game prints

	var secret_number = rand.IntN(100) //a random secret number (between 100)
	var guess int //empty integer variable

	// a loop for asking the guess and check the guess
	for {
		fmt.Print("Guess a number : >") //print guess message
		fmt.Scan(&guess) //scan guess

		//if guess is greater than secret number
		if (guess > secret_number){
			fmt.Println("Your guess was too high!") //print "guess was too high" message
			continue //continue
		//else if guess is less than secret number
		}else if (guess < secret_number){
			fmt.Println("Your guess was too low!") //print "guess was too low" message
			continue //continue
		//else if guess is equals to the secret number
		}else if (guess == secret_number) {
			fmt.Println("Your guess and secret number matched!") //print success message
			break //break the loop
		}
	}
}
