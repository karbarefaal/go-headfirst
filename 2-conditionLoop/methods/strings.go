package methods

import (
	"fmt"
	"strings"
)

func MyReplace() {
	var text string = "Be H#mal#a khosh amad#d"
	replacer := strings.NewReplacer("#", "i")
	fixed := replacer.Replace(text)
	fmt.Println(fixed)
}
