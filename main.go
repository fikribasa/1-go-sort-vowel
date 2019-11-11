package main 

import (
	"fmt"
	"sort"
	"strings"
)

func SortString (w string) string{
	s:=strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func main(){
	var conso, vowel string
	inputString:="osama"
	fmt.Println(inputString)
	
	inputString=SortString(inputString)
	for _, value := range inputString {
		switch value {
			case 'a','e','i','u','o','A','E','I','U','O' :
				vowel+=string(value)
			default :
				conso+=string(value)
		}
	}
	fmt.Println("Hasil : ",vowel+conso)

}
