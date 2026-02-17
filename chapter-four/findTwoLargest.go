package main
import "fmt"

func getLargestAndSecondLargest(numbers []int) (int, int){
	var largest = numbers[0];
	var secondLargest = 0;
	for count:= 1; count < len(numbers); count++{
		if numbers[count] > largest{
			secondLargest = largest;
			largest = numbers[count];

		}else if numbers[count] < largest && numbers[count] > secondLargest{
			secondLargest = numbers[count];
		
		
		}
	}
	
	return largest,secondLargest

}

func main(){
	var numbers = make([]int, 10);
	var number int;
	
	for count := 1; count < 11;count++{
		fmt.Print("Enter number ", count, ": ")
		fmt.Scan(&number);
		numbers = append(numbers, number);
	
	}
	
	largest,secondLargest := getLargestAndSecondLargest(numbers);
	
	fmt.Println("Largest is : ", largest);
	fmt.Println("Second Largest is : ", secondLargest);


}





