package main
import "fmt"	

func main() {

	var expenseCategory = []string{"Food", "Clothing", "Housing", "Transportation", "Education", "Healthcare", "Entertainment", "Other"	}	
	
	var expense float64
	var totalExpense float64
	fmt.Print("Enter expense for category: ")
	fmt.Scan(&expense)

	for i := 0; i < len(expenseCategory); i++ {	
		fmt.Printf("Enter expense for %s: ", expenseCategory[i])
		fmt.Scan(&expense)
		totalExpense += expense
	}	
	var fairTaxRate float64 = 0.15
	var tax float64 = totalExpense * fairTaxRate
	fmt.Printf("\nTotal expenses: %.2f\n", totalExpense)
	fmt.Printf("Tax: %.2f\n", tax)	




	

}
