package paintingWalls

import (
	"fmt"
)

var metersPerliter float64

func PaintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f liters needed\n", area/10.0)
}
