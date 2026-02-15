package main
import "fmt"

func main() {
	fmt.Println("1 + 1 = ", 1 + 1)
	//var x string =  "fathia";

	//var number = 2;

	//numb := 9.2

	//var name = "fathia";

	//name = "jjj";

	
	
	
	//fmt.Print("Enter a number: ");
	//var input float64
	//fmt.Scanf("%f", &input)
	//output := input * 2


	//var num = 2
	//for num < 5{
	//	fmt.Println(num);
	//	num += 1;
	//}
	
	
	for nums := 10; nums < 15; nums ++{
		fmt.Println(nums);

	}
	
	var array[5]int
	for index := 0; index < 5; index++{
		array[index] = index + 5;
		
	}
	fmt.Println(array);

}
