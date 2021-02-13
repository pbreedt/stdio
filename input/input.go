package input

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInt reads an integer from standard in
func ReadInt(prompt string) (int, error) {
	inStr, err := ReadString(prompt)
	inNum, err := strconv.Atoi(inStr)
	if err != nil {
		return 0, err
	}
	return inNum, nil
}

// ReadFloat reads an float from standard in
func ReadFloat(prompt string) (float64, error) {
	inStr, err := ReadString(prompt)
	inNum, err := strconv.ParseFloat(inStr, 64)
	if err != nil {
		return 0, err
	}
	return inNum, nil
}

// ReadBool reads a boolean from standard in
func ReadBool(prompt string, allowYN bool) (bool, error) {
	inStr, err := ReadString(prompt)
	inBool, err := strconv.ParseBool(inStr)
	if err != nil {
		if allowYN {
			switch inStr {
			case "Y", "y", "yes", "YES", "Yes":
				return true, nil
			case "N", "n", "no", "NO", "No":
				return false, nil
			}
		}
		return false, err
	}
	return inBool, nil
}

// ReadString reads a string from standard in
func ReadString(prompt string) (string, error) {
	if len(prompt) > 0 {
		fmt.Print(prompt)
	}

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}

	input = strings.TrimSpace(input)

	return input, nil
}
