package main
import (
				"fmt"
				"bufio"
				"os"
				_"strconv"
				"strings"
				"io"
				"math"
				//"time"
)

var enum = map[string]int{
		"A": 0,
		"R": 1,
		"N": 2,
		"D": 3,
		"C": 4,
		"Q": 5,
		"E": 6,
		"G": 7,
		"H": 8,
		"I": 9,
		"L": 10,
		"K": 11,
		"M": 12,
		"F": 13,
		"P": 14,
		"S": 15,
		"T": 16,
		"W": 17,
		"Y": 18,
		"V": 19,
		"B": 20,
		"Z": 21,
		"X": 22,
		"STAR": 23,
}

var blosum = [][]int{
    { 4, -1, -2, -2, 0, -1, -1, 0, -2, -1, -1, -1, -1, -2, -1, 1, 0, -3, -2, 0, -2, -1, 0 },
    { -1, 5, 0, -2, -3, 1, 0, -2, 0, -3, -2, 2, -1, -3, -2, -1, -1, -3, -2, -3, -1, 0, -1 },
    { -2, 0, 6, 1, -3, 0, 0, 0, 1, -3, -3, 0, -2, -3, -2, 1, 0, -4, -2, -3, 3, 0, -1 },
    { -2, -2, 1, 6, -3, 0, 2, -1, -1, -3, -4, -1, -3, -3, -1, 0, -1, -4, -3, -3, 4, 1, -1 },
    { 0, -3, -3, -3, 9, -3, -4, -3, -3, -1, -1, -3, -1, -2, -3, -1, -1, -2, -2, -1, -3, -3, -2 },
    { -1, 1, 0, 0, -3, 5, 2, -2, 0, -3, -2, 1, 0, -3, -1, 0, -1, -2, -1, -2, 0, 3, -1 },
    { -1, 0, 0, 2, -4, 2, 5, -2, 0, -3, -3, 1, -2, -3, -1, 0, -1, -3, -2, -2, 1, 4, -1 },
    { 0, -2, 0, -1, -3, -2, -2, 6, -2, -4, -4, -2, -3, -3, -2, 0, -2, -2, -3, -3, -1, -2, -1, },
    { -2, 0, 1, -1, -3, 0, 0, -2, 8, -3, -3, -1, -2, -1, -2, -1, -2, -2, 2, -3, 0, 0, -1 },
    { -1, -3, -3, -3, -1, -3, -3, -4, -3, 4, 2, -3, 1, 0, -3, -2, -1, -3, -1, 3, -3, -3, -1 },
    { -1, -2, -3, -4, -1, -2, -3, -4, -3, 2, 4, -2, 2, 0, -3, -2, -1, -2, -1, 1, -4, -3, -1 },
    { -1, 2, 0, -1, -3, 1, 1, -2, -1, -3, -2, 5, -1, -3, -1, 0, -1, -3, -2, -2, 0, 1, -1 },
    { -1, -1, -2, -3, -1, 0, -2, -3, -2, 1, 2, -1, 5, 0, -2, -1, -1, -1, -1, 1, -3, -1, -1 },
    { -2, -3, -3, -3, -2, -3, -3, -3, -1, 0, 0, -3, 0, 6, -4, -2, -2, 1, 3, -1, -3, -3, -1 },
    { -1, -2, -2, -1, -3, -1, -1, -2, -2, -3, -3, -1, -2, -4, 7, -1, -1, -4, -3, -2, -2, -1, -2 },
    { 1, -1, 1, 0, -1, 0, 0, 0, -1, -2, -2, 0, -1, -2, -1, 4, 1, -3, -2, -2, 0, 0, 0 },
    { 0, -1, 0, -1, -1, -1, -1, -2, -2, -1, -1, -1, -1, -2, -1, 1, 5, -2, -2, 0, -1, -1, 0 },
    { -3, -3, -4, -4, -2, -2, -3, -2, -2, -3, -2, -3, -1, 1, -4, -3, -2, 11, 2, -3, -4, -3, -2 },
    { -2, -2, -2, -3, -2, -1, -2, -3, 2, -1, -1, -2, -1, 3, -3, -2, -2, 2, 7, -1, -3, -2, -1 },
    { 0, -3, -3, -3, -1, -2, -2, -3, -3, 3, 1, -2, 1, -1, -2, -2, 0, -3, -1, 4, -3, -2, -1 },
    { -2, -1, 3, 4, -3, 0, 1, -1, 0, -3, -4, 0, -3, -3, -2, 0, -1, -4, -3, -3, 4, 1, -1 },
    { -1, 0, 0, 1, -3, 3, 4, -2, 0, -3, -3, 1, -1, -3, -1, 0, -1, -3, -2, -2, 1, 4, -1 },
    { 0, -1, -1, -1, -2, -1, -1, -1, -1, -1, -1, -1, -1, -1, -2, 0, 0, -2, -1, -1, -1, -1, -1 } };

func Max(a, b int) int {
    if a > b {
        return a
    }
    return b
}
func isMax(a, b int) bool {
    if a > b {
        return true
    }
    return false
}
type pair struct {
				key string
				value string
}
func MaxRow(row []int) int {
		max := int(math.Inf(1))
		retVal := -1
		for index, element := range row {
				if element > max {
						max = element
						retVal = index
				}
		}
		return retVal
}
func printArray(row []int) {
		pf := fmt.Printf
		pf("[")
		for i:=0; i<len(row);i++ {
				pf("%d\t", row[i])
		}
		fmt.Println("]")
}
func main() {

	reader := bufio.NewReader(os.Stdin)
	p := fmt.Println
	_, _ = reader.ReadString('\n')
	inputs := make(map[string]string, 0)
	comparisons := make([]pair, 0)
	alpha_gap := -4
	//for _, row := range blosum {
	//	p(row)
	//}
	for {
				text, err  := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				currInput := strings.Split(text, " ")
				if len(currInput) > 1 {
					currInput[1] = strings.TrimSuffix(currInput[1], "\n")
					currPair := pair{currInput[0], currInput[1]}
					comparisons = append(comparisons, currPair)
				} else {


					text = strings.TrimSuffix(text, "\n")
					next, _ := reader.ReadString('\n')
					next = strings.TrimSuffix(next, "\n")
					inputs[text] = next
				}
	}
	p("String Inputs\n",inputs)
	p("Comparisons\n",comparisons)
	// Remove one letter in X and Y
	// case 1 remove from X
	// case 2 remove from X and Y
	// case 3 remove from Y

	// OPT(i,j) = MIN(ALPHA(Xi,Yi) + P(
	for _, pair := range comparisons {
		p("COMPARING _______________:\n",pair.key," -- ",pair.value)
		//optAlignment := make(map[string]string, 0)
		if len(inputs[pair.key]) < len(inputs[pair.value]) {
				temp := pair.key
				pair.key = pair.value
				pair.value = temp
		}
		p("key afterswap:",pair.key, " -- ", pair.value)
		optm := make([][]int, len(inputs[pair.key])+1)
		for i:=0; i<len(optm); i++ {
				optm[i] = make([]int, len(inputs[pair.value])+1)
		}
		for m:=0; m<len(optm); m++ {
				optm[m][0] = m*alpha_gap
		}
		for n:=0; n<len(optm[0]); n++ {
				optm[0][n] = n*alpha_gap
		}
		myX := strings.Split(inputs[pair.key], "")
		myY := strings.Split(inputs[pair.value], "")
		//optm[0][0] = blosum[enum[myX[0]]][enum[myY[0]]]
		for i := 1; i<len(optm); i++ {
			for j := 1; j<len(optm[0]); j++ {
					//p(myX[i])
					//p(myY[j])
					//p(enum[myX[i]])
					//p(enum[myY[j]])

					alpha := blosum[enum[myX[i-1]]][enum[myY[j-1]]]
					diag  := alpha+optm[i-1][j-1]
					down := alpha_gap + optm[i-1][j]
					right := alpha_gap + optm[i][j-1]

					traverseMin := Max(diag, down)
					optm[i][j] = Max(right, traverseMin)
					//p("I:",i,"J:",j," The value chosen: ", optm[i][j])
			}
		}
		prevIndex := len(optm[0]) - 1
		output := ""
		var temp bool
		for k:=len(optm)-1; k>0; k-- {
						if prevIndex == 0{
								temp = false
						} else {
								temp = isMax(optm[k-1][prevIndex-1], optm[k-1][prevIndex])
								p("temp:", temp,"prev", prevIndex)
						}

						if temp {
								output = myY[prevIndex-1] + output
								prevIndex--
						} else {
								output = "_"+ output
						}
						//p("max index:", temp)
						//p("max val:", optm[k][temp])
						//p(optm[k])


		}
		//output += myY[0]
		p(output)
		p(strings.Join(myX, ""))
		//p(optm)
		for i:=0; i<len(optm); i++ {
						printArray(optm[i])
		}
		p("The answer is: ", optm[len(optm)-1][len(optm[0])-1])
	}

}
