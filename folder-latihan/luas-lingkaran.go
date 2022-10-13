package main

import "fmt"

func main() {
    
    var spRadius, spSA float64

    fmt.Print("Masukan radius lingkaran: ")
    fmt.Scanln(&spRadius)

    spSA = 4 * 3.14 * spRadius * spRadius
  
    fmt.Println("Luas lingkaran  = ", spSA)
}