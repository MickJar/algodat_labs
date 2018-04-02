package main

import (
				"fmt"
				"os"
				"strconv"
				"time"
)



func fib(n float64, cache map[float64]float64) float64 {
		var retVal float64
		if n == 1 || n == 2 {
				retVal = 1
		} else if val, ok := cache[n]; ok {
				return val
		} else {
				cache[n-1] = fib(n-1, cache)
				cache[n-2] = fib(n-2, cache)
				retVal = cache[n-1] + cache[n-2]
		}
		return retVal
}

func main() {

				//argsWithoutProg := os.Args[1:]
				p := fmt.Println
				timeBefore := time.Now()
				n, _ := strconv.ParseFloat(os.Args[1], 64)
				p("Starting with Fib of ", n)

				cache := make(map[float64]float64, 0)
				p(fib(n, cache))
				timeAfter := time.Now()
				p("Time Taken: ", timeAfter.Sub(timeBefore))
}
