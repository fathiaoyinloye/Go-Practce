package main
func sortNumbers(numbers []int)[] int{
	
	for count:= 0; count < len(numbers); count++{

		smallest := numbers[count]

		for counter:=count + 1; counter < len(numbers); counter++{
			if numbers[counter] < smallest{
				temp := smallest
				smallest = numbers[counter]
				numbers[counter] = temp

			}
		}
		numbers[count] = smallest

	}
	return numbers;


}