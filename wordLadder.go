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
			visited bool
}
type pair struct {
			key string
			value string
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

//func BFS_helper(nodes []nodes, layer int)
//{
//}
func get(a []node, toFind string) *node {
				for _, element := range a {
								//fmt.Println("Comparing: ", element.name, " and ", toFind)
								if strings.EqualFold(element.name, toFind) {
												//fmt.Println("Match Found")
												return &element
								}
				}
				return nil
}

func BFS(nodes []node, root string, toFind string) int {
	//node[0].visited = true
	rootNode := get(nodes, root)

	//fmt.Println("root: ", rootNode.name, "rootString: ", root)
	rootNode.visited = true
	layer := 1
	queue := make([]node, 0)
	queue = append(queue, *rootNode)
	for {
		v := queue[0]
		// Discard top element
		queue = queue[1:]
		//fmt.Println("v: ", v.name)

		for _,element := range v.children {
				if element.visited == false {
						fmt.Println("Layer:", layer, "ELEMENT:", element.name)
						element.visited = true
						elementTree := get(nodes, element.name)
						elementTree.visited = true
						fmt.Println("Local copyi: ", elementTree)
						fmt.Println("Tree version: ",get(nodes, element.name))
						queue = append(queue, *elementTree)
						//maybe pred
						////fmt.Println("ToFind: ", toFind, "ele:", element.name)
						if strings.EqualFold(toFind, element.name) {
								//fmt.Println("return layer", layer)
								return layer
						}
				}
		}
		// Is empty ?
		if len(queue) < 1 {
						break
		}
		layer++
	}
	return -1
}

func main() {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Enter text; ")
		words := make([]string, 0)
		children := make(map[string]string, 0)
		pairs := make([]pair, 0)
		for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				keyValue := strings.Split(text, " ")
				keyValue[0] = strings.TrimSuffix(keyValue[0], "\n")
				if len(keyValue) > 1 {
						keyValue[1] = strings.TrimSuffix(keyValue[1], "\n")
						pairs = append(pairs, pair{keyValue[0], keyValue[1]})
				}else {
						words = append(words, keyValue[0])
				}
		}

		fmt.Println("*-*-*-*-*-* WORDS *-*-*-*-*-*")
		fmt.Println(words)
		fmt.Println("*-*-*-*-*-* PAIRS *-*-*-*-*-*")
		fmt.Println(pairs)
		nodes := make([]node, 0)
		for i := 0; i < len(words); i++ {
				currNode := node{words[i], make([]node, 0), false}
				for j:=0; j < len(words); j++ {
					if j == i {
									continue
					} else {
							isMatching := checkWords(words[i], words[j])
							if isMatching {
									newNode := node{words[j], make([]node, 0), false}
									currNode.children = append(currNode.children, newNode)
									children[currNode.name] += newNode.name
							}
					}
				}
				nodes = append(nodes, currNode)
		}
		//for _, element := range nodes {
		//	fmt.Println(element.name, ": ")
		//	for _, ele2 := range element.children {
		//		fmt.Print(ele2.name)
		//		fmt.Println(", ")
		//	}
		//	fmt.Println("__________")
		//}
		//for _, n := range
		//BFS(nodes,
		for _, p := range pairs {
				//fmt.Println("key:", p.key, "val:", p.value)
				fmt.Println(BFS(nodes, p.key, p.value))
		}
}


