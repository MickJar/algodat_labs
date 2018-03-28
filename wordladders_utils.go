package main

import "fmt"
import "strings"

func checkLetters(word1 string, word2 string) bool{
	arr1 := strings.Split(word1,"")
	arr2 := strings.Split(word2,"")

	counter := 0
	for i:= 1; i <= 4; i++{
		for j:= 0; j < len(arr2); j++{
			if arr2[j] == arr1[i]{
	            arr1[i] = ""
	            arr2[j] = ""
	            counter++
	            break
	        }
	    }
	}
	if(counter == 4){
		return true
	}else{
		return false
	}
}


func main(){

	check := checkLetters("there", "other")
	fmt.Println(check)
}