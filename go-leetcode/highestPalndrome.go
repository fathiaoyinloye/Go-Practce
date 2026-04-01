package main
import (
	"fmt"
)


func longestPalindrome(s string) string {

    if len(s) == 1{
        return s;
    }
    answer := ""
	for index := 0; index < len(s); index++ {
        check := string(s[index])
        word := string(s[index])
		for counter :=  index + 1; counter < len(s); counter++{
                // count := counter + 1
				if string(s[index]) != string(s[counter]) {
                    check = string(s[counter]) + "" + check
                    word = word + "" + string(s[counter])
                

                // }else if len(s) > count && string(s[index]) == string(s[counter]) && string(s[counter]) == string(s[count]) {
                //     fmt.Println( " Count ", count, " Index ", index, " Counter ", counter, "len(s) ", len(s))
                //     fmt.Println( " Count ", count)
                //     check = string(s[counter]) + "" + check
                //     word = word + "" + string(s[counter])
                }else {
					check = string(s[counter]) + "" + check
                    word = word + "" + string(s[counter])
                    if  check == word && len(word) > len(answer){
                        answer = word
                    }

                   
				}

			}
	}
    if len(s) == 2 && answer == ""{
        return string(s[0])
    }
	return answer;



}
func main() {
	fmt.Println(longestPalindrome("ccc "), "answer")
}