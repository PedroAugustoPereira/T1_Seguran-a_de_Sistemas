package main

import (
	"fmt"
	"strings"
)

func main() {
	chave := "sol"
	message := "diadia"
	message = strings.ToLower(message)
	message = strings.Trim(message, " ")

	lenKey := len(chave)
	nIndexKey := 0
	result := ""

	// 97
	// 122

	for i := 0; i < len(message); i++ {
		nAsciimessage := int(message[i]) - 97   // Normalize 'a' to 0
		nAsciiKey := int(chave[nIndexKey]) - 97 // Normalize 'a' to 0

		fmt.Printf("Asciimessage: %c: %d\n", message[i], message[i])
		fmt.Printf("AsciiKey: %c: %d\n", chave[nIndexKey], chave[nIndexKey])

		// Calculate shifted character (0-25 range)
		nAscii := (nAsciimessage + nAsciiKey) % 26

		// Map it back to 'a' - 'z' (97-122)
		nAscii = nAscii + 97

		result = result + string(rune(nAscii))
		nIndexKey++

		if nIndexKey == lenKey {
			nIndexKey = 0
		}
	}

	fmt.Println(result)
}
