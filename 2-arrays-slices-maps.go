package main

import "fmt"

func main()  {
	var favNums[5] float64

	favNums[0] = 142
	favNums[1] = 1422
	favNums[2] = 14.2
	favNums[3] = 0.142
	favNums[4] = 1.42

	// With index
	for i, value := range favNums {
		fmt.Println("Value:", value, "at index:", i)
	}

	// Without index
	for _, value := range favNums {
		fmt.Println("Value:", value)
	}

	// Slice
	numSlice := []int {1, 2, 3, 4, 5}
	fmt.Println("numSlice[2] =", numSlice[:2])

	// Maps
	user := make(map[string] int)
	user["GeorgiYanev"] = 30
	user["Crap Bag"] = 42

	fmt.Println(user)
	fmt.Println(user["GeorgiYanev"])

	delete(user, "Crap Bag")
	fmt.Println(user)

	// List of numbers
	listNums := []float64{1, 2, 3, 4, 5}
	fmt.Println("Sum of nums:", sumNums(listNums))

	// Return 2 values from func
	num1, num2 := get2Values(5)
	fmt.Println(num1, num2)

	// Subtract values
	fmt.Println(subtractValues(1, 2, 3, 4, 5))
}

// Sum list values
func sumNums(numbers []float64) float64 {
	sum := 0.0

	for _, val := range numbers {
		sum += val
	}

	return sum
}

// Return 2 values
func get2Values(number int) (int, int) {
	return number+1, number+2
}

// Subtract values
func subtractValues(args ...int) int {
	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue
}
