package main
import "fmt"


func getPow(base, power int) int {
	 total := 1;
	for count:= 1; count <= power; count++{
		total *=  base;
	
	}
	
	return total;
}
func palindrome(number int) int{

	var decimal = 0;
	for count:= 0; number > 0; count++{
		if count == 0{
		decimal =  number % 10;
		}else{
			var result = number % 10
			decimal += result * getPow(2,count)
			
		}
		number = number/10

	
	}
	return decimal;

}


func main(){
	fmt.Println(palindrome(1101));


}
