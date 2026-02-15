package main
import "fmt";
func calculatenewBalance(beginingBalance, charges, credit int) int{
	
	return beginingBalance + charges - credit;
}


func checkCreditLimit(newBalance, creditLimit int) string{
	if newBalance > creditLimit{
		return "Credit Limit Exceeded";
	}else{
		return "Credit Limit not Exceeded";
	}
}

func main(){
	var accountNumber int;
	var beginingBalance int;
	var charges int;
	var credit int;
	var creditLimit int;
	var accounts = make([]int, 0)
	var creditLimits = make([]string, 0)
	
	for count:= 1; count != 0; count++{
		fmt.Print("Enter customer ", count, "account number: ");
		fmt.Scan(&accountNumber);
		accounts = append(accounts,accountNumber);
	
		fmt.Print("Enter customer beginning balance: ");
		fmt.Scan(&beginingBalance);
	
		fmt.Print("Enter customer charges: ");
		fmt.Scan(&charges);
	
		fmt.Print("Enter customer credit: ");
		fmt.Scan(&credit);
		
		var newBalance = calculatenewBalance(beginingBalance, charges, credit);
	
		fmt.Print("Enter customer Credit Limit: ");
		fmt.Scan(&creditLimit);
		creditLimits = append(creditLimits, checkCreditLimit(newBalance, creditLimit))

		
		fmt.Print("Enter -1 if you have enter all customers details: ");
		fmt.Scan(&count);

	}
	
	for count:= 0; count < len(creditLimits); count++{
		fmt.Println(creditLimits[count] + " by ", accounts[count]);
	}

}
