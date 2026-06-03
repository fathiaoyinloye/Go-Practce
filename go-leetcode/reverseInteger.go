package main
import (
    "strconv"
    "math"
)

func reverse(x int) int {
    // highestVal := 2147483647
    // lowestVal := -2147483648
    stringValue := strconv.Itoa(x) 
    lastChar := string(stringValue[len(stringValue)-1])
    firstChar:= string(stringValue[0])
    lastDigit, _ := strconv.Atoi(lastChar)
    firstDigit, _ := strconv.Atoi(firstChar)


    if x > 0 && len(stringValue) > 10{
        return 0;

    }else if x > 0 && len(stringValue) > 10 {
        if firstDigit > 7 {
            return 0
        }else if lastDigit > 2{
            return 0
        }
        return 0;


    }else if x < 0 && len(stringValue) > 11{
        return 0;

    }else if x < 0 && len(stringValue) == 11 && lastDigit >= 8 {
        return 0;
    }
    
    absoluteValue := int(math.Abs(float64(x)))
    answer := ""
    for count := 1; absoluteValue > 0; count++ {
        digit := absoluteValue % 10
        answer = answer + "" + strconv.Itoa(digit)
        if count == 1 && digit == 0 {
            answer = ""
        }
        absoluteValue = absoluteValue / 10

    }
        y , _ := strconv.Atoi(answer)

    if x < 0{

        return y * -1
    }

    return y
}


