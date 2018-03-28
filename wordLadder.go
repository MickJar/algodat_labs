package main

import (
				"bufio"
				"fmt"
				"os"
				"io"
				"strings"
				//"go_labs/wordladders_utils"
)

type node struct {
			name string
			children []node
}

func checkWords(word1 string, word2 string) bool{
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

func main() {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Enter text; ")
		words := make([]string, 0)
		children := make(map[string]string, 0)

		for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				words = append(words, text)
		}
		fmt.Println(words)
		nodes := make([]node, 0)
		for i := 0; i < len(words); i++ {
				currNode := node{words[i], make([]node, 0)}
				for j:=0; j < len(words); j++ {
					if j == i {
									continue
					} else {
							isMatching := checkWords(words[i], words[j])
							if isMatching {
									newNode := node{words[j], make([]node, 0)}
									currNode.children = append(currNode.children, newNode)
									children[currNode.name] += newNode.name
							}
					}
				}
				nodes = append(nodes, currNode)
		}
		for _, element := range nodes {
			
			fmt.Println(element.name, ": ")
			for _, ele2 := range element.children {
				fmt.Print(ele2.name)
			}
			fmt.Println("__________")
		}
}


