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

func main() {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Enter text; ")
		words := make([]string, 0)
		for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				words = append(words, text)
		}
		fmt.Println(words)

		for i := 0; i < len(words); i++ {
				currNode := node{words[i], make([]node, 0)}
				for j:=0; j < len(words); j++ {
					if j == i {
									continue
					} else {
							isMatching := strings.ContainsRune(words[j],rune(words[i][1])) &&
														strings.ContainsRune(words[j],rune(words[i][2])) &&
														strings.ContainsRune(words[j],rune(words[i][3])) &&
														strings.ContainsRune(words[j],rune(words[i][4]))

							if isMatching {
									newNode := node{words[j], make([]node, 0)}
									currNode.children = append(currNode.children, newNode)
							}
					}
				}
		}
}


