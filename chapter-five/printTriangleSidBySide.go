package main
import "fmt"

func printRighAngleTriangleSideBySide(){
	var firstTriangle = 1;
	var firstSpaceBegining = 2;
	var firstSpaceEnding = 12;
	
	for count := 1; count <=10;count++{
		for counter := 1; counter <= 46; counter++{
			if count <= firstTriangle && count <=12{
			fmt.Print("*")
			}
			if counter >= firstSpaceBegining && counter >= firstSpaceEnding{
				fmt.Print(" ")
			}

		}
		fmt.Println()
		firstSpaceBegining = firstSpaceBegining + 1
		
	
	}


}
