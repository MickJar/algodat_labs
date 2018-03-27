package main

import (
				"bufio"
				"fmt"
				"os"
				"strings"
				"strconv"
)
type person struct {
				no int
				name string
				pref []int
}
func get(a []person, perNo int) int {
				for index, element := range a {
								if element.no == perNo {
												return index
								}
				}
				return 1000
}
func main() {
		reader := bufio.NewReader(os.Stdin)
		//fmt.Print("Enter text; ")
		text, err  := reader.ReadString('\n')
		for ;strings.Contains(text, "#"); {
			text, err  = reader.ReadString('\n')
			if err != nil {
							fmt.Println("Stupid error:", err)
			}
		}
		if err != nil {
				fmt.Println("Stupid error")
		}
		noVals := strings.Split(text, "=");
		fmt.Println("noVals:", noVals)
		fmt.Println("noVals0:", noVals[0])
		fmt.Println("noVals1:", noVals[1])
		noVals[1] = strings.TrimSuffix(noVals[1], "\n")
		num, aerr := strconv.Atoi(noVals[1])

		if aerr != nil {
						fmt.Println("Weird error:", aerr)
		}
		fmt.Println(num)

		men := make([]person, 0)
		women := make([]person, 0)
		for i := 0; i < 2*num; i++ {
				text, err  := reader.ReadString('\n')
				if err != nil {
						fmt.Println("Stupid error")
				}
				nameNo := strings.Split(text, " ");
				if i % 2 == 0 {
					input, _ := strconv.Atoi(nameNo[0])
					men = append(men, person{input, nameNo[1], nil})
				} else {
					input, _ := strconv.Atoi(nameNo[0])
					women = append(women, person{input, nameNo[1], nil})
				}
		}
		for i := 0; i < 2*num; i++ {
				text, err  := reader.ReadString('\n')
				if strings.EqualFold(text, "\n") {
					text, err = reader.ReadString('\n')
				}
				if err != nil {
						fmt.Println("Stupid error")
				}
				perNo := strings.Split(text, ": ");
				attractedTo := strings.Split(perNo[1], " ");
				currPer, _ := strconv.Atoi(perNo[0]);

				for _, element := range attractedTo {
						if currPer % 2 != 0 {
							element = strings.TrimSuffix(element, "\n")
							input, _ := strconv.Atoi(element)
							currMan := get(men, currPer)
							men[currMan].pref =	append(men[currMan].pref, input)
						} else {
							element = strings.TrimSuffix(element, "\n")
							input, _ := strconv.Atoi(element)
							currMan := get(women, currPer)
							women[currMan].pref = append(women[currMan].pref, input)
						}
					// index is the index where we are
					// element is the element from someSlice for where we are
				}
		}
		fmt.Println("Men:\n ", men);
		fmt.Println("Women:\n ", women);

		freeMen := men
		freeWomen := women
		marriedCouples := make(person[][], 0)
		menLeft = len(freeMen)
		for ;menLeft > 0; {
				freeMen[0]
		}
}

