// Package baseconverter converts any number from an arbitrary base to another arbitrary base.
// From binary up to base62. Not only from base2 to base36 like in standard library. The package
// also contains two additional functions to convert from an arbitrary base to decimal and from
// decimal to an arbitrary base.
package baseconverter

import (
	"fmt"
	"math"
	"strconv"
)

// List of all digits from base2 to base62.
const allDigits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Flag of the negative number.
var flagNeg uint

// Conversion converts number from an arbitrary base to another arbitrary base.
func Conversion(toConv string, fromBase, toBase int) (string, error) {
	// Converting the number to decimal first.
	decNum, err := ConvertToDec(toConv, fromBase)
	if err != nil {
		return "", err
	}
	// Converting the number from decimal to the declared base(toBase).
	newBaseNum, err := ConvertFromDec(decNum, toBase)
	if err != nil {
		return "", err
	}
	return newBaseNum, nil
}

// ConvertToDec converts an arbitrary base number to decimal base number.
func ConvertToDec(numS string, fromBase int) (string, error) {
	if fromBase < 2 {
		return "", fmt.Errorf("Invalid base: %v", fromBase)
	}
	if numS == "" {
		return "", fmt.Errorf("Nothing to convert")
	}

	// If the number is negative: flagNeg = 1 and "saving" the negative sign.
	num := []rune(numS)
	if num[0] == rune('-') {
		flagNeg = 1
		num = append(num[1:])
	}

	// Converting the number to decimal base.
	var decNum int64
	iFromEnd := len(num) - 1
	for _, vRune := range num {
		vInt, err := letterToNum(vRune, fromBase)
		if err != nil {
			return "", err
		}
		decNum += int64(vInt) * int64(math.Pow(float64(fromBase), float64(iFromEnd)))
		iFromEnd--
	}
	if flagNeg == 1 {
		decNum *= -1
	}

	return fmt.Sprint(decNum), nil
}

// ConvertFromDec converts decimal base number to an arbitrary base number.
func ConvertFromDec(numS string, toBase int) (string, error) {
	if toBase < 2 {
		return "", fmt.Errorf("Invalid base: %v", toBase)
	}
	if numS == "" {
		return "", fmt.Errorf("Nothing to convert")
	}
	// Converting string to int64
	num, err := strconv.ParseInt(numS, 10, 64)
	if err != nil {
		return "", fmt.Errorf("Invalid number: %v", numS)
	}
	// If original number is equal to "0...0".
	if num == 0 {
		return "0", nil
	}
	// If the number is negativ: flagNeg = 1 and "saving" the negative sign.
	if num < 0 {
		flagNeg = 1
		num *= -1
	}
	// Converting the number to needed base (toBase).
	var newNum string
	for num != 0 {
		reminder := num % int64(toBase)
		num = num / int64(toBase)
		reminderHex := numToLetter(reminder, toBase)
		newNum += reminderHex
	}

	if flagNeg != 0 {
		// If original number was negative - making new one negative too.
		newNum += "-"
		flagNeg = 0
	}
	// Reversing the number.
	numRunes := []rune(newNum)
	left := 0
	right := len(numRunes) - 1
	for left < len(numRunes)/2 {
		numRunes[left], numRunes[right] = numRunes[right], numRunes[left]
		left++
		right--
	}

	return string(numRunes), nil
}

// letterToNum converts rune(it can be digit(0-9) or letter(a-z,A-Z) to int.
func letterToNum(digit rune, base int) (int, error) {
	if base > 10 {
		for i, v := range allDigits[10:] {
			if digit == v {
				return i + 10, nil
			}
		}
	}
	digitInt, err := strconv.Atoi(string(digit))
	if err != nil {
		return 0, fmt.Errorf("Invalid digit %v", string(digit))
	}
	return digitInt, nil
}

// numToLetter converts int64 digit to string with digits(0-9) or/and letters(a-z,A-Z).
func numToLetter(digit int64, base int) string {
	if base > 10 {
		for i, v := range allDigits[10:base] {
			if digit == int64(i+10) {
				return string(v)
			}
		}
	}
	return fmt.Sprint(digit)
}
