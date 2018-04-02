package main

import (
				"fmt"
				"os"
				"strconv"
				"time"
)



func binomialCoeff(n int, k int, cache[][]int) int {
		var retVal int
		//fmt.Println("n: ", n, " k: ",k)
		if k == 0 || k == n {
				retVal = 1
		} else if cache[n][k] != 0 {
				retVal = cache[n][k]
		} else {
				cache[n-1][k-1] = binomialCoeff(n-1,k-1, cache)
				cache[n-1][k] = binomialCoeff(n-1,k,cache)
				retVal = cache[n-1][k-1] + cache[n-1][k]
		}
		//fmt.Printf("C(%d,%d) is %d\n", n,k,retVal)
		return retVal
}

func main() {

				//argsWithoutProg := os.Args[1:]
				p := fmt.Println
				timeBefore := time.Now()
				n, _ := strconv.Atoi(os.Args[1])
				k, _ := strconv.Atoi(os.Args[2])
				p("Finding Binomial Coefficient C(",n,",",k,")")

				cache := make([][]int, n+1)
				for i:=0; i<n+1; i++ {
						cache[i] = make([]int, k+1)
				}
				//cache[n] = []int{0,0,0}
				//fmt.Println("cache:\n", cache)
				p(binomialCoeff(n, k, cache))
				timeAfter := time.Now()
				p("Time Taken: ", timeAfter.Sub(timeBefore))
}
