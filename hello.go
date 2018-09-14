/*
*Author:-Pukar Giri
*Created On:-13Th September 2018 at 15:28
*File Name:-hello.go
*Project Name:-learning
*Licence:- MIT
*Email:-crazzy.lx75@gmail.com
*/
package hello

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func Hello() {
	response,err:=http.Get("https://geek-jokes.sameerkumar.website/api")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer response.Body.Close()
	contents,err:=ioutil.ReadAll(response.Body)
	if err != nil{
		fmt.Println(err)
		return
	}
	fmt.Println(string(contents))
}
