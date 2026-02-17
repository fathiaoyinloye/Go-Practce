package main
import "fmt"

func printRightAngleTriangle(length int){
	for count := 1; count <= length; count++{
		for counter := 1; counter <= count; counter++{
			fmt.Print("* ");
		
		}
		fmt.Println();
	}

}

func main(){
	printRightAngleTriangle(10)

}
