package main
import "fmt"
func getDivisibleByThree() int{
	var sum int
	var count int;
	for count = 1; count <= 30;count++{
		if(count % 3 == 0){
			sum += count
		}
	}
	return count;

}

func main(){
	fmt.Print(getDivisibleByThree());


}
