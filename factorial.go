package main
import "fmt"
func getFactorial(number int) int {
	var total = 1;
	for count:= number; count > 1; count--{
		total *= count;
	
	}
	return total

}

func main(){
	fmt.Println(getFactorial(5))

}
