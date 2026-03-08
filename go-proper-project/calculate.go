package main
import( 
	"fmt"
	"errors"
)

func calculate(firstNumber, secondNumber int) (string, error){
	if secondNumber == 0{
		return "", errors.New("Cannot divide by two");
	}
	message := fmt.Sprintf("%v divided by %v is %v", firstNumber, secondNumber, firstNumber / secondNumber)
	return message,nil;

}
