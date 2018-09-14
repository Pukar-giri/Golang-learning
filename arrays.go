/*
*Author:-Pukar Giri
*Created On:-14Th September 2018 at 12:45
*File Name:-arrays.go
*Project Name:-learning
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
*/
package array

import "fmt"

func array() {
	//arrays in go are of static length
	var a[5]int
	var b[10]string
	//arrays can be defined at declaration by usingtype infrence operator with {} braces
	names:= [5] string {
"hello","hi","namaskar","Pukar","Giri"}
	fmt.Println(names[3])
	//declaring but not using variable is not allowed in go if you do not need some return value we can use _ as dummy variable instead
}