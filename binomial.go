package main

import (
				"fmt"
				"os"
				"strconv"
				"time"
)



func binomialCoeff(n float64, k float64) float64 {
		var retVal float64
		//fmt.Printf("C(%f,%f)\n", n,k)
		if k == 0 || k == n {
				retVal = 1
		} else {
				retVal = binomialCoeff(n-1, k-1) + binomialCoeff(n-1, k)
		}
		//fmt.Printf("C(%f,%f) is %f\n", n,k,retVal)
		return retVal
}

func main() {

				//argsWithoutProg := os.Args[1:]
				p := fmt.Println
				timeBefore := time.Now()
				n, _ := strconv.ParseFloat(os.Args[1], 64)
				k, _ := strconv.ParseFloat(os.Args[2], 64)
				p("Finding Binomial Coefficient C(",n,",",k,")")

				p(binomialCoeff(n, k))
				timeAfter := time.Now()
				p("Time Taken: ", timeAfter.Sub(timeBefore))
}
