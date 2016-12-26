package main

import "fmt"

func main()  {
  // Closure
  num := 3
  doubleNum := func() int {
    num *= 2
    return num
  }

  fmt.Println(doubleNum())
  fmt.Println(doubleNum())
  fmt.Println(doubleNum())
  fmt.Println("------------------")

  fmt.Println(factorial(5))
  fmt.Println("------------------")

  // Defer = execute after main has finished
  defer printTwo()
  printOne()
  fmt.Println("------------------")

  // Safe division
  fmt.Println(safeDiv(3,0))
  fmt.Println(safeDiv(3,2))
  fmt.Println("------------------")

  // Panic
  demoPanic()
}

// Recursion
func factorial(num int) int {
  if num == 0 {
    return 1
  }

  return num * factorial(num-1)
}

// Defer
func printOne()  {
  fmt.Println(1)
}
func printTwo()  {
  fmt.Println(2)
}

func safeDiv(num1, num2 int) int {
  defer func() {
    fmt.Println(recover())
  }()

  solution := num1 / num2
  return solution
}

func demoPanic() {
  defer func ()  {
    fmt.Println(recover())
  }()
  panic("Panic")
}
