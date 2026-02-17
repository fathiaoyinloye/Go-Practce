package main
import "fmt"

func getLargestAndSecondLargest(numbers []int) (int, int){
	var largest = numbers[0];
	var lowest = 0;
	for count:= 1; count < len(numbers); count++{
		if numbers[count] > largest{
			largest = numbers[count];

		}else if numbers[count]  < lowest{
			lowest = numbers[count];
		
		}
	}
	
	return largest,lowest

}

func main(){
	var numbers = make([]int, 10);
	var number int;
	var value int;
	fmt.Print("Enter how many value you want to enter: ")
	fmt.Scan(&value)
	for count := 1; count <= value;count++{
		fmt.Print("Enter number ", count, ": ")
		fmt.Scan(&number);
		numbers = append(numbers, number);
	
	}
	
	largest,secondLargest := getLargestAndSecondLargest(numbers);
	sum := largset + lowest;
	fmt.Println("Largest is : ", largest);
	fmt.Println("Second Largest is : ", secondLargest);


}





