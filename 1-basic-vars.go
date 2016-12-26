package main

import "fmt"

func main() {
  fmt.Println("Hello World")

  var age int = 30
  var favNum float64 = 7.311324
  randNum := 11

  const pi float64 = 3.14159265
  var myName string = "Georgi"
  var isOld bool = true

  fmt.Println("You are:", age, "years old.")
  fmt.Printf("%.3f \n", favNum)
  fmt.Println(randNum, pi)
  fmt.Println(len(myName))
  fmt.Printf("Type of age var is: %T \n", age)
  fmt.Printf("%t \n", isOld)
  fmt.Printf("%d \n", 420)
  fmt.Printf("2 in binary is: %b \n", 2)
}
