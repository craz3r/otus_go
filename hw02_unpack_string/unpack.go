package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func isDigit(input rune) bool {
	return input >= '0' && input <= '9'
}

func Unpack(input string) (string, error) {
	if input == "" {
		return "", nil
	}

	var sb strings.Builder
	inputR := []rune(input)

	i := 0

	for i < len(inputR) {
		symbol := inputR[i]

		if isDigit(symbol) {
			return "", ErrInvalidString
		}

		if i+1 < len(inputR) && isDigit(inputR[i+1]) {
			count := inputR[i+1]
			iCount, err := strconv.Atoi(string(count))
			if err != nil {
				return "", ErrInvalidString
			}

			sb.WriteString(strings.Repeat(string(symbol), iCount))
			i += 2
			continue
		}

		sb.WriteString(string(symbol))
		i += 1
	}

	return sb.String(), nil
}
