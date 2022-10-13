package main

import "fmt"

var area int   // Variable declare outside the main function

func main(){
    var l,b int //Declaration of multiple Variables 
   
    fmt.Print("Masukan lebar persegi: ")
    fmt.Scan(&l)
	fmt.Print("Masukan panjang persegi: ")
    fmt.Scan(&p)
    area = l*p
    fmt.Print("Luas persegi : ",area)
}