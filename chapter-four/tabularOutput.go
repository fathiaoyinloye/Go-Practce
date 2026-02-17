package main;
import "fmt";
	
func getPow(base, power int) int {
	 total := 1;
	for count:= 1; count <= power; count++{
		total *=  base;
	
	}
	
	return total;
}



func printTable(){
	for count := 0; count <= 5; count++{
		if(count == 0){
		fmt.Println("N\tN^2\tN^3\tN^4")
		}else{
		fmt.Println(getPow(count,1),"\t", getPow(count , 2),"\t", getPow(count,3),"\t", getPow(count,4),"\t");
		
		}
	
	}

}



func main(){
	printTable();


}

