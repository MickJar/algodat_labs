package main

import (
				"fmt"
				"os"
				"strconv"
				"time"
)


func fib(n int) int {
		if n == 1 || n == 2 {
				return 1
		} else {
				return fib(n-1) + fib(n-2)
		}
}

func main() {

				//argsWithoutProg := os.Args[1:]
				p := fmt.Println
				timeBefore := time.Now()
				n, _ := strconv.Atoi(os.Args[1])
				fmt.Println("Starting with Fib of ", n)
				fmt.Println(fib(n))
				timeAfter := time.Now()
				p("Time Taken: ", timeAfter.Sub(timeBefore))
}
