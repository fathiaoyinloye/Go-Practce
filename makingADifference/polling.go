package main
import "fmt"
func getTotal(numbers [10]int{
	total:= 0;
	for count:= 0; len(number);count++{
	total += numbers[count];
	}
	return total;

}

func getTotal(numbers [10]float{
	total:= getTotal(numbers);
	
	return total / 10;

}
func getTotal(numbers [5][10]int)[4][5]int{

	var numbers[4] int;
	for count:= 0; len(number);count++{
			for count:= 0; len(number);count++{
	
		digit:= number % 10;
		digit += 7;
		digit = digit % 10;
		numbers[count] = digit;
		number = number/10;
		
	}
	return numbers;

}
func getSumPlusSevenModuloTenDecrypt(number int)[4]int{

	var numbers[4] int;
	for count:= 3; number != 0 ;count--{
		digit:= number % 10;
		digit += 3;
		digit = digit % 10;
		numbers[count] = digit;
		number = number/10;
		
	}
	return numbers;

}

func swapNumbers(numbers [4]int)[4]int{
	number:= numbers[2]
	numbers[2] = numbers[0]
	numbers[0] = number
	number = numbers[1]
	numbers[1] = numbers[3]
	numbers[3] = number
	return numbers;


}

func main(){
	var numbers[4] int;
	numbers = getSumPlusSevenModuloTenDecrypt(2189);
	fmt.Print(swapNumbers(numbers));


}
