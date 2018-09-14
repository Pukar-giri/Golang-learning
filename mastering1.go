/*
*Author:-Pukar Giri
*Created On:-14Th September 2018 at 12:59
*File Name:-mastering1.go
*Project Name:-learning
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
*/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {
	arr := []int{0}
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 100000000; i++{
	m:=rand.Int63()%10
	l:=len(arr)
	arr=append(arr,arr[l-1]+int(m))
}
	fmt.Println(arr)
	Search(arr,arr[rand.Int63()%100000000])
}
func Search(arr[] int,num int) int{

	n:=search(arr,num)
	if n<0{
		return -1
	}else{
		fmt.Println("just found that",num," lies at index",n)
		return n

	}

}

func search(arr[] int,num int) int{
	mid:=int(len(arr)/2)
	if mid==0 && num!=arr[mid]{
		fmt.Println("just found that",num,"does not exist")
		return int(math.Pow(-2,63))
	}else if num<arr[mid] {
		fmt.Println("just found that",num,"lies below",arr[mid])
		return search(arr[0:mid],  num)
	}else if num>arr[mid]{
		fmt.Println("just found that",num,"lies above",arr[mid])
		return mid+search(arr[mid:],num)
	}else{
		return mid
	}
}