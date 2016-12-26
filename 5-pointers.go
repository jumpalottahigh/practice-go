package main

import "fmt"

func main()  {
  x := 0
  yPtr := new(int)
  changeY(yPtr)
  fmt.Println("y =", *yPtr)

  changeX(&x)
  fmt.Println("x =", x)
  fmt.Println("Memory address for x =", &x)
}

// Basic pointer value changes
func changeX(x *int) {
  *x = 2
}

func changeY(yPtr *int) {
  *yPtr = 100
}
