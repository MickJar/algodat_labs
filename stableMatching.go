package main


import (
				"bufio"
				"fmt"
				"os"
				"strings"
				"strconv"
)

type stack []int

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
		// fmt.Println("noVals:", noVals)
		// fmt.Println("noVals0:", noVals[0])
		// fmt.Println("noVals1:", noVals[1])
		noVals[1] = strings.TrimSuffix(noVals[1], "\n")
		num, aerr := strconv.Atoi(noVals[1])

		if aerr != nil {
						fmt.Println("Weird error:", aerr)
		}
		// fmt.Println(num)

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
		// fmt.Println("Men:\n ", men)
		// fmt.Println("Women:\n ", women)


		preferences := createPreferences(men, women)
		// preferences := [][]int{{2,4,6,8}, {1,7,5,3}, {6,4,2,8}, {7,3,1,5}, {6,4,8,2}, {5,3,7,1}, {4,8,6,2}, {7,5,1,3}}
		// preferences := [][]int{{6, 4, 2}, {3, 5, 1}, {2, 6, 4}, {5, 1, 3}, {6, 4, 2}, {1, 5, 3}}
	   	engaged := make(map[int]int, 0)
	    s := make(stack,0)
	   	s = createArray(len(preferences))
	    length := len(s)
		m := 0

	    for ;length > 0;  {
		    s, m = s.Pop()
		    prefListLength := len(preferences[m-1])
			for i := 0; i < prefListLength; i++ {
		   		w := preferences[m-1][i]
		   		if value, ok := engaged[w]; ok {
		   			//value borde inte vara med
				    value = value+1
				    if getIndex(preferences[w-1], engaged[w]) > getIndex(preferences[w-1], m){
				    	s = append(s, engaged[w])
				    	engaged[w] = m
				    	break
				    }
				} else {
					engaged[w] = m
					break
				}
			}   
		    
		    length = len(s)
	    }
	   	for index := range engaged {
	   		fmt.Print(strings.TrimSuffix(men[get(men, engaged[index])].name, "\n"))
	   		fmt.Print(" -- ")
	   		fmt.Print(women[get(women, index)].name)


	   	}
	   	//fmt.Println(engaged)

}

func createPreferences(men []person, women []person) [][]int{
	preferences := make([][]int, len(men)+len(women))
	for i := range preferences{
		preferences[i] = make([]int, 0)
	}
	j := 0
	for i := 0; i < len(men)+len(women); i= i+2{
		preferences[i] = append(preferences[i], men[j].pref...)
		preferences[i+1] = append(preferences[i+1],women[j].pref...)
		j++
	}
	// fmt.Println(preferences)
	return preferences
}
func (s stack) Push(v int) stack {
    return append(s, v)
}

func (s stack) Pop() (stack, int) {
	if len(s) > 0{
	    l := len(s)
	    return  s[:l-1], s[l-1]
	}else{
		return s, -1
	}
}

func contains(s []int, e int) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func getIndex(list []int, person int) int{
	for index, element:= range list{
		if element == person {
			return index
		}
	}
	return -1
}
func remove(s []int, i int) []int {
		if len(s) > 0{
	    s[len(s)-1], s[i] = s[i], s[len(s)-1]
	    return s[:len(s)-1]
	}
	return s
}

func createArray(nbr int) stack{
	s := make(stack,0)
	for i:= 1; i < nbr; i= i+2{
		s = s.Push(i)
	}
	return s
}
