/*
*Author:-Pukar Giri
*Created On:-14Th September 2018 at 12:04
*File Name:-datatypes.go
*Project Name:-learning
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
*/
package datatypes



func datatype() {
	//Beauty of this language is the use of diffrent sized datatype right in its declaration making it easier and highly intuitive
	var a int8
	var b int16
	var c int32
	var d int64
	//using only the int  allows go to choose any size for the variable for int it is either 32 or 64 bits based on your architecture ie x86 or x64
	var e int
	//float data type can be 32 or 64 bit long
	var f float32
	var g float64
	//another datatype is unsigned integer which has similar sizes to int but lacks the sign bit
	var h uint8
	var i uint16
	var j uint32
	var k uint64
	var l uint
	//go also provides a string datatype which is flexible and allows you to change length.
	var m string
	m="a"
	m="hello world"
	//there is a boolean variable in go that does not require size
	var truth bool
	truth=true
	//there is a complex no variable for strongly scientific application
	var n complex64
	n=5+6i
	var o complex128
	o=12546+123456897i
	//when you are inside any function  go allows us to use the type infrence using the type infrence operator ie :=
	num1:=5
	num2:=500.5
	num3:=5+6i
	name:="Pukar Giri"
	truth:=true

}