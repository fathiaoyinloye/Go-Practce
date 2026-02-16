package main
import "fmt"

func main(){
	var passes = 0;
	var failures = 0;
	for count:= 1; count <= 10; count++{
		var result int64;
		fmt.Print("Enter result (1 = pass, 2 = fail): ");
		fmt.Scan(&result)
		if (result == 1) {
			passes = passes + 1;
		}else if result == 2{
			failures = failures + 1;
	
		} else {
			fmt.Println("Invalid Input Result can either be 1 or 2");
			count -= 1;
		}
		
	}


	fmt.Println("Passed", passes, "failures", failures)
	if (passes > 8) {
	fmt.Println("Bonus to instructor!")
	}
}
