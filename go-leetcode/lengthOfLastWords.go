package main

import "strings"
func lengthOfLastWord(s string) int {
    check := 0;
    trimmed := strings.TrimSpace(s) 
    for index := len(trimmed)-1; index >= 0; index--{
        if string(trimmed[index]) == " "{
            break;
        }
        check += 1

    }
    return check;

    
}
