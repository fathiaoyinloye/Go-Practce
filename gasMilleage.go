package main
import "fmt"


func getMilesPerGallon(milesDriven, gallonsUsed int) float64{
	return float64(milesDriven/gallonsUsed )

}
func getTotalMilesPerGallon(milesPerGallons []float64) float64{
	var total = 0.0;
	for count:= 0; count < len(milesPerGallons); count++{
	total += float64(milesPerGallons[count]);
	}
	return total;

}

func main(){
	milesPerGallons := make([]float64, 0);
	var milesDriven = 0
	var gallonUsed = 0

	for count:= 0; count != -1; count++{
	fmt.Print("Enter miles Driven for trip: ");
	fmt.Scan(&milesDriven);
	
	fmt.Print("Enter gallon used for trip: ");
	fmt.Scan(&gallonUsed);
	milesPerGallon := getMilesPerGallon(milesDriven,gallonUsed);

	milesPerGallons = append(milesPerGallons, milesPerGallon);
	
	
	fmt.Print("Enter -2 to stop or any number to continue: ");
	fmt.Scan(&count);
	
	}
	
	for index:= 1; index <= len(milesPerGallons);index++{

		fmt.Println("Miles Per Gallon of trip", index,"is: ", milesPerGallons[index -1] );
	
	}
		fmt.Println("Total Miles Per Gallon of all trips is: ",getTotalMilesPerGallon(milesPerGallons) );

}

