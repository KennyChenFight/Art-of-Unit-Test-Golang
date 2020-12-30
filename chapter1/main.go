package main

import (
	"fmt"
	"github.com/KennyChenFight/Art-of-Unit-Test-Golang/chapter1/parser"
)

func main() {
	simpleParser := parser.NewSimpleParser()
	numbers := ""
	sum, err := simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "1"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "1,2"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "1 2"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "1, 2"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "1,2 3"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)

	numbers = "n"
	sum, err = simpleParser.ParseAndSum(numbers)
	fmt.Printf("sum=%d, err=%v\n", sum, err)
}
