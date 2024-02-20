package hamming

import (
	"errors"
	"fmt"
)

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Input size doesn't match")
	}
	return calcDistance(a, b)
}

func calcDistance(a, b string) (int, error) {
	// for i, ac := range a { // runes?
	var distance = 0
	for i := 0; i < len(a); i++ {
		var ac = a[i]
		var bc = b[i]
		fmt.Printf("%d %c=%c\n", i, ac, bc)
		if ac != bc {
			fmt.Printf("match\n")
			distance++
		}
	}
	return distance, nil
	//panic("Implement the Distance function")
}
