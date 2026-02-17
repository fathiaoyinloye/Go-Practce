package main
import "fmt"

func getLargestAndSecondLargest(numbers []int) (int, int){
	var largest = numbers[0];
	var lowest = numbers[0];
	for count:= 0; count < len(numbers); count++{
		if numbers[count]  < lowest{
			lowest = numbers[count];
		
		}
		if numbers[count] > largest{
			largest = numbers[count];

		}
		
	}
	
	return lowest,largest

}

func main(){

	var number int;
	var value int;
	fmt.Print("Enter how many value you want to enter: ")
	fmt.Scan(&value)
	var numbers = make([]int,0);
	for count := 1; count <= value;count++{
		fmt.Print("Enter number ", count, ": ")
		fmt.Scan(&number);
		numbers = append(numbers, number);
	
	}
	
	lowest,largest := getLargestAndSecondLargest(numbers);
	sum := largest + lowest;
	fmt.Println("Largest is : ", largest);
	fmt.Println("Lowest is : ", lowest);
	fmt.Println("Sum is : ", sum);


}





