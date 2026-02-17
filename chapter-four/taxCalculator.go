package main
import "fmt"
func calculateTaxRate(earning float64)float64{
	if earning > 30000{
		return earning * 0.5
	}else{
	return earning * 0.5
	}
}


func main(){
	var name string;
	var earning float64;
	var names = make([]string, 0);
	var taxes = make([]float64, 0);

	for count := 1; count != 0; count++{
		fmt.Print("Enter citizen name: ")
		fmt.Scan(&name);
		names = append(names, name);

		fmt.Print("Enter citizen earning: ")
		fmt.Scan(&earning);
		taxes = append(taxes, calculateTaxRate(earning));
	
		fmt.Print("Enter -1 if you have entered all citizen tax or any number to continue: ")
		fmt.Scan(&count);
	
	}
	
	for count:= 0; count < len(names); count++{
		fmt.Println(names[count], "tax rate is:", taxes[count])
	
	}


}
