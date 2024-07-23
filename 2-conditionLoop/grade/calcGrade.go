package grade

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func passedOrFailed() {

	fmt.Println("Enter a grade: ")
	buff := bufio.NewReader(os.Stdin)
	input, err := buff.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(input)
	input = strings.TrimSpace(input)
	grade, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if grade >= 60 {
		status = "passed"
	} else {
		status = "failed"
	}
	fmt.Printf("You are %s", status)

}
