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
			children []*node
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
func mutateChildren(a *[]node, n node) {
		
}
//func BFS_helper(nodes []nodes, layer int)
//{
//}
func get(a *[]*node, toFind string, change bool) *node {
				for _, element := range *a {
								//fmt.Println("Comparing: ", element.name, " and ", toFind)
								if strings.EqualFold(element.name, toFind) {
												//fmt.Println("Match Found")
												element.visited = change
												return element
								}
				}
				return nil
}
func printPred(pred map[string]node, node node, count int) int {
		if _, ok := pred[node.name]; !ok {
				fmt.Println(node.name)
				return count
		}
		fmt.Print(node.name, "<-")
		count++
		return printPred(pred, pred[node.name], count)
}
func BFS(nodes []*node, root string, toFind string) int {
	//node[0].visited = true
	//fmt.Println(nodes)
	if strings.EqualFold(root, toFind) {
		return 0
	}
	for _, n := range nodes {
			n.visited = false;
	}
	//fmt.Println(nodes)
	rootNode := get(&nodes, root, true)

	//fmt.Println("root: ", rootNode.name, "rootString: ", root)
	//rootNode.visited = true
	layer := 1
	queue := make([]node, 0)
	pred := make(map[string]node,0)
	queue = append(queue, *rootNode)
	for {
		v := queue[0]
		// Discard top element
		queue = queue[1:]
		//fmt.Println("v: ", v.name)

		for _,element := range v.children {
				if element.visited == false {
						element.visited = true
						elementTree := get(&nodes, element.name, true)
						//queue = append(queue, *elementTree)
						//fmt.Println("Visiting Node:", *element)
						for _ , child := range elementTree.children {
							queue = append(queue, *child)
						}
						pred[element.name] = v
						//pred = append(pred, *element)
						if strings.EqualFold(toFind, element.name) {
								//fmt.Println("return layer", layer)
								//for _, p := range pred {
								//		fmt.Print(p.name, " -> ")
								//}
								fmt.Println(pred)
								return printPred(pred, *element, 0)

								//fmt.Println(pred)
								//return len(pred)
								//return layer
						}
				}
		}
		// Is empty ?
		if len(queue) < 1 {
						break
		}
		//fmt.Println("UPDATING LAYER WITH CURRNODE: ", v)
		layer++
	}
	return -1
}

func main() {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Enter text; ")
		words := make([]string, 0)
		//children := make(map[string]string, 0)
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
		nodes := make([]*node, 0)
		for _, w := range words {
			nodes = append(nodes, &node{w, nil, false})
		}
		var currNode *node
		for i := 0; i < len(words); i++ {
				currNode = get(&nodes, words[i], false)
				for j:=0; j < len(words); j++ {
					if j == i {
									continue
					} else {
							isMatching := checkWords(words[i], words[j])
							if isMatching {
									//newNode := node{words[j], make([]node, 0), false}
									//fmt.Println(currNode.children)
									//fmt.Println("local Reference: ", &currNode)
									//fmt.Println("nodes reference: ", get(&nodes, words[j], false))
									currNode.children = append(currNode.children, get(&nodes, words[j], false))
									//fmt.Println(currNode.children)
									//fmt.Println(get(&nodes, words[i], false))
									//children[currNode.name] += newNode.name
							}
					}
				}
				//mutateChildren(nodes, currNode)
				//nodes = append(nodes, currNode)
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


