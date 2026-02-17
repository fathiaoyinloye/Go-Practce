package main
import "fmt"

func printRighAngleTriangleSideBySide(){
	var firstTriangle = 1;
	var firstSpaceBegining = 2;
	var secondTriangleBegining = 13
	var secondTriangleEnding = 22
	var secondSpaceBegining = 23
	var secondSpaceEnding = 24
	var thirdTriangleBegining = 25
	var thirdTriangleSpaceEnding = 45
	var fourthTrianglBegining = 46

	
	
	for count := 1; count <=10;count++{
		for counter := 1; counter <= 46; counter++{
			if counter <= firstTriangle && counter <=12{
			fmt.Print("*")
			}
			if counter >= firstSpaceBegining && counter <= 12{
				fmt.Print(" ")
			}
			if counter >= secondTriangleBegining && counter <= secondTriangleEnding{
				fmt.Print("*")
			}
			if counter >= secondSpaceBegining && counter <= secondSpaceEnding{
				fmt.Print(" ")
			}
			
			if counter >= thirdTriangleBegining && counter <= 34{
				fmt.Print("*")
			}
			if counter >= 35 && counter <= thirdTriangleSpaceEnding{
				fmt.Print(" ")
			}
			if counter >= fourthTrianglBegining && counter <= 46{
				fmt.Print("*")
			}
			

		}
		fmt.Println()
		firstTriangle = firstTriangle + 1
		firstSpaceBegining = firstSpaceBegining + 1;
		secondTriangleEnding = secondTriangleEnding - 1;
		secondSpaceBegining = secondSpaceBegining - 1;
		secondSpaceEnding = secondSpaceEnding + 1;
		thirdTriangleBegining = thirdTriangleBegining +  1
		thirdTriangleSpaceEnding = thirdTriangleSpaceEnding -  1
		fourthTrianglBegining = fourthTrianglBegining - 1

	}


}

func main(){
printRighAngleTriangleSideBySide()
}
