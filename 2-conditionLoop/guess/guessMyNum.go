package guess

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// guess challenges players to guess a random number.
func GeussMyNum() {
	//	getting a number from the system
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	// comparing two values
	yourNum := bufio.NewReader(os.Stdin)
	success := false
	for guesses := 0; guesses < 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")
		//	getting a number from our user
		input, err := yourNum.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess > target {
			fmt.Println("Ops. Your guess was HIGH")
		} else if guess < target {
			fmt.Println("Ops. Your guess was LOW")
		} else {
			success = true
			fmt.Println("Good job! You guessed it.")
			break
		}
	}
	if !success {
		fmt.Println("Sorry, you didn't get my number. It was:", target)
	}
}
