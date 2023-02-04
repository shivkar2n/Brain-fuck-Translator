package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var A [30000]int8
	dp := 0
	ip := 0

	for i := 0; i < len(A); i++ {
		A[i] = int8(0)
	}

	inputBytes, _ := ioutil.ReadFile(os.Args[1])
	inputString := string(inputBytes)

	for ; ip < len(inputString); ip++ {
		switch inputString[ip] {
		case '>':
			dp++

		case '<':
			dp--

		case '+':
			A[dp]++

		case '-':
			A[dp]--

		case '.':
			fmt.Printf("%c", A[dp])

		case ',':
			fmt.Scanf("%c", A[dp])

		case '[':
			if A[dp] == 0 {
				stack := 0
				for j := ip + 1; ; j++ {
					if inputString[j] == '[' {
						stack++
					} else if inputString[j] == ']' {
						if stack != 0 {
							stack--
						} else {
							ip = j
							break
						}
					}
				}
			}

		case ']':
			if A[dp] != 0 {
				stack := 0
				for j := ip - 1; ; j-- {
					if inputString[j] == ']' {
						stack++
					} else if inputString[j] == '[' {
						if stack != 0 {
							stack--
						} else {
							ip = j
							break
						}
					}
				}
			}
		}
	}
}
