package main
import (
	"slices"
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	slice := []string{}
    check :=[]string{}
	for index := 0; index < len(s); index++ {
		exists := slices.Contains(slice, string(s[index]))
		if !exists {
			slice = append(slice, string(s[index]))
		}else{
            if len(slice) > len(check){
                check = slices.Clone(slice);

            }
            slice = nil
            slice = append(slice,string(s[index]))
        }
	}
    //fmt.Print("check ",check)
    fmt.Print(slice)

    if len(check) > len(slice){
        return len(check)

    }
   
	return len(slice)

}