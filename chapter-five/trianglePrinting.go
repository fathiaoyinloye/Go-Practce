package main
import "fmt"

func printTriangleDownward(){
	for count:= 1; count <=10;count++{
		for counter := 1; counter <= count; counter++{
			fmt.Print("*");
		}
			fmt.Println();
	}

}


func printTriangleToptoBottom(){
	for count:= 10; count >=1;count--{
		for counter := 1; counter <= count; counter++{
			fmt.Print("*");
		}
			fmt.Println();
	}

}


func printTriangleToptoBottomTiltToRight(){
	var check = 0;
	for count:= 10; count >=1;count--{
		for counter := 1; counter <= 10; counter++{
			if(counter <= check){
				print(" ")

			}else{
				fmt.Print("*");
			}

		}
			check = check + 1;
			fmt.Println();
	}

}

func printTriangleToptoBottomTiltToLeft(){
	var check = 9;
	for count:= 10; count >=1;count--{
		for counter := 1; counter <= 10; counter++{
			if(counter <= check){
				print(" ")

			}else{
				fmt.Print("*");
			}

		}
			check = check - 1;
			fmt.Println();
	}

}

func main(){
	printTriangleDownward()
	fmt.Println();
	printTriangleToptoBottom()
	fmt.Println();
	printTriangleToptoBottomTiltToRight();
	fmt.Println();
	printTriangleToptoBottomTiltToLeft();

}
