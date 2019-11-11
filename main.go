package main

import (
	"fmt"
	"sort"
	"strings"
)

func SortString(w string) string {
	var vowel, conso, word string
	s := strings.Split(w, "")
	sort.Strings(s)
	newS := strings.Join(s, "")

	for _, value := range newS {
		switch value {
		case 'a', 'e', 'i', 'u', 'o', 'A', 'E', 'I', 'U', 'O':
			vowel += string(value)
		default:
			conso += string(value)
		}
	}
	word = (vowel + conso)
	fmt.Println("Hasil Sort: ", vowel+conso)
	return word
}

func main() {
	fmt.Println("Go String Vowel-Consonant Sorter")
	inputString := readLine("Input Word to Sort: ")

	SortString(inputString)

}

func readLine(word string) string {
	fmt.Print(word)
	var input string
	fmt.Scanln(&input)
	return input
}
