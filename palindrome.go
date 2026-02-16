package main
import "fmt"

func palindrome(number int) bool{
	
	var check = number;
	var unit = 10;
	var newNumber = 1;
	for count:= 1; check > 0; count++{
		var result = check % 10

		check = check/10

		if(count == 1){
			newNumber = result;

		}else{
			newNumber = newNumber * unit + result

		}
	
	}
	return number == newNumber;

}


func main(){
	fmt.Println(palindrome(12321));


}
