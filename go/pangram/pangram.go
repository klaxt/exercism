package pangram

import (
	"fmt"
	"strings"
)


/**
 * A pangram (Greek: παν γράμμα, pan gramma, "every letter") is a sentence using every letter of the alphabet at least once. The best known English pangram is: The quick brown fox jumps over the lazy dog
 */
func IsPangram(input string) bool {
	var chars = [26]rune{'a','b','c','d','e','f','g','h','i','j','k','l','m','n','o','p','q','r','s','t','u','v','w','x','y','z'};
	var lowerInput = strings.ToLower(input)
	for i, c := range chars {
		fmt.Printf("%d %c\n", i, c)
		if !strings.ContainsRune(lowerInput, c) {
			fmt.Printf("mising %s\n", input)
			return false
		}
	}

	return true
	//panic("Please implement the IsPangram function")
}
