package main
import "fmt";
func main(){
	
	//y := make([] int, 5,7)
	var y = make([] int, 5,7)
	fmt.Println(y);
	
	for count :=0; count < 5; count++{
		y[count] = count + 1;
	
	}
		y = append(y, 10)
		
		for count :=6; count < 10; count++{
		y= append(y,count + 6);
	
	}
	slices.Insert(y, 0, 10)
		fmt.Println(y);
		fmt.Println(len(y));
}
