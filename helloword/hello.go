package main

import (
	"fmt"

	//"golang.org/x/exp/slices"
)

var name string = "Yasel"

func main() {
	fmt.Printf("hello, world\n")

	fmt.Println(name)

	var age = [3]int{25, 30, 35}

	//array
	var name =[4]string{"pepe", "antonio", "juan", "luigi"}
		//name[1]= "pech"
	fmt.Println(age, len(age))
	fmt.Println(name, len(name))

//slices
var score =[]int{100, 50, 60}
score[0]= 40

score = append(score, 85)
score = append(score, 200)
fmt.Println(score, len(score))

//slices Range
rango1 := name[1:3]
rango2 := name[2:]   //despues del 2
rango3 := name[:3]    //antes del 3
fmt.Println(rango1)
fmt.Println(rango1, rango2, rango3)
rango1 = append(rango1, "koopa")

fmt.Println(rango1)

}
