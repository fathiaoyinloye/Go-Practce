package main
import "fmt"

func getLargest(numbers []int)int{
	var largest = numbers[0];
	for count:= 1; count < len(numbers); count++{
		if numbers[count] > largest{
			largest = numbers[count];
		}
	}
	return largest

}

func main(){
	var numbers = make([]int, 10);
	var number int;
	
	for count := 1; count < 11;count++{
		fmt.Print("Enter number ", count, ": ")
		fmt.Scan(&number);
		numbers = append(numbers, number);
	
	}
	fmt.Print("Largest is : ", getLargest(numbers));


}



