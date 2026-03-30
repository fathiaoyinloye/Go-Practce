package main
import (
	"slices"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	slice := []string{}
    check :=[][]string{}
    if len(s) == 1{
        return 1
    }
	for index := 0; index < len(s); index++ {
		slice = append(slice, string(s[index]))
		for counter :=  index + 1; counter < len(s); counter++{
				exists := slices.Contains(slice, string(s[counter]))
				if !exists {
					slice = append(slice, string(s[counter]))
				}else{
					check = append(check, slice)
					slice = nil
					break
				}
                if counter == len(s)- 1{
                    check = append(check,slice)
                    slice = nil
                }

			}
	}
    fmt.Print(check)

	highest := 0;
	for _, row := range check {
		if len(row) > highest{
			highest = len(row)
		}
	}   
	return highest;

}