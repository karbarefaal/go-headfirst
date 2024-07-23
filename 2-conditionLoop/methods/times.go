package methods

import (
	"fmt"
	"time"
)

func MyTime() {
	var now time.Time = time.Now()
	var year int = now.Year()
	fmt.Println(year)
}
