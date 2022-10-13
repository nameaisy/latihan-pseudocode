package main

import "fmt"

func main() {
    
    var spRadius, spSA, spVolume

    fmt.Print("Masukan radius lingkaran: ")
    fmt.Scanln(&spRadius)

    spVolume = (4.0/3) * 3.14 * spRadius * spRadius * spRadius

    fmt.Println("Volume lingkaran  = ", spVolume)
}