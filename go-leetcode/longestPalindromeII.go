///package main
import (
	"fmt"
)


 func isPalindrone(s string) bool {
    runes := []rune(s)

    for i, j := 0, len(runes)-1; i<j; i, j=i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    //this is to reverse string in go, i searched this on google.
    return string(runes) == s
}


func longestPalindrome(s string) string {
    answer := ""

    for i := 0; i < len(s); i++{
        for j := i + 1; j < len(s); j++{
            substr := s[i:j]
            if isPalindrome(substr) {
                if len(substr) > len(answer){
                    answer = substr 
                }
            }
        }
    }

    return answer
