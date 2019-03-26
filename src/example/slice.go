package main

import (
    "fmt"
)

func main() {
  a := []int{5,4,3,2,1} //creat a slice which is more flexible
  a = append(a, 13)
  fmt.Println(a)
}
