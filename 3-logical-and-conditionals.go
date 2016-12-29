package main

import "fmt"

func main()  {
	// Boolean operators
	fmt.Println("true && false =", true && false)
	fmt.Println("true || false =", true || false)

	// For loops
	i := 0
	for i < 10 {
		fmt.Println("I = ", i)
		i++
	}

	// Another for loop
	for j := 1; j <= 5; j++ {
		fmt.Println("J = ", j)
	}

	// Conditionals
	myAge := 31

	if (myAge >= 30) {
		fmt.Println("You are old!")
	} else {
		fmt.Println("You are not that old.")
	}

	// Switch
	age := 40
	switch age {
		case 20: fmt.Println("You are 20 years old!")
		case 30: fmt.Println("You are 30 years old!")
		case 40: fmt.Println("You are 40 years old!")
		default:  fmt.Println("Hmm, you are a mystery!")
	}
}
