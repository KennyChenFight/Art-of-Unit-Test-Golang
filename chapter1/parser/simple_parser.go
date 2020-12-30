package parser

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrInvalidParse = errors.New("just handle numbers for now")
)

type SimpleParser struct {
}

func NewSimpleParser() *SimpleParser {
	return &SimpleParser{}
}

func (s *SimpleParser) ParseAndSum(numbers string) (int, error) {
	if len(numbers) == 0 {
		return 0, nil
	}

	token := strings.Split(numbers, ",")
	var sum int
	for _, t := range token {
		num, err := strconv.Atoi(t)
		if err != nil {
			return 0, ErrInvalidParse
		}
		sum += num
	}
	return sum, nil
}
