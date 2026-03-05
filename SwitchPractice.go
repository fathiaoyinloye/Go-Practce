package main
import (
	"fmt"
	"time"
	)


func main(){
	i := 9
    	fmt.Print("Write ", i, " as ")
   	 switch i {
   	 case 1:
        	fmt.Println("one")
    	case 2:
        	fmt.Println("two")
    	case 3:
        	fmt.Println("three")
        	
        default:
        	fmt.Println("Just a number")
        	
        }
        
        //case with two value seperated by comma
        switch time.Now().Weekday() {
    	case time.Saturday, time.Sunday:
        	fmt.Println("It's the weekend")
    	default:
        	fmt.Println("It's a weekday")
    }
    
    
    	//switch without an expression
	t := time.Now()
    	switch {
    	case t.Hour() < 12:
        	fmt.Println("It's before noon")
    	default:
        	fmt.Println("It's after noon")
    }
	fmt.Println(time.Now().Weekday())
	
	
	
	whatAmI := func(i any) {
        switch t := i.(type) {
        case bool:
            fmt.Println("I'm a bool")
        case int:
            fmt.Println("I'm an int")
        default:
            fmt.Printf("Don't know type %T\n", t)
        }
    }
    whatAmI(true)
    whatAmI(1)
    whatAmI("hey")


}
