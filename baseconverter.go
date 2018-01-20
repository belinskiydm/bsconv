// Package baseconverter converts number from an arbitrary base to another arbitrary base.
// From binary up to base62.
package baseconverter

import (
	"fmt"
	"log"
	"math"
	"strconv"
)

// List of all digits from base2 to base62.
const allDigits = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Flag of the negativ number.
var flagNeg uint

// Conversion converts number from an arbitrary base to another arbitrary base.
func Conversion(toConv string, fromBase, toBase int) (string, error) {
	// Basic error checkings.
	if fromBase < 2 || toBase < 2 {
		return "error", fmt.Errorf("Invalid base: %v or %v", fromBase, toBase)
	}
	if len(toConv) == 0 {
		return "error", fmt.Errorf("Nothing to convert")
	}
	if fromBase == toBase {
		return toConv, nil
	}
	// If the number is negativ: flagNeg = 1 and "saving" the negative sign.
	numS := []rune(toConv)
	if numS[0] == rune('-') {
		flagNeg = 1
		numS = append(numS[1:])
	}
	// Checking if number exists and belongs to the declared base(fromBase).
	if len(numS) == 0 || numIsValid(numS, fromBase) != true {
		return "error", fmt.Errorf("Invalid number: %v", toConv)
	}
	// Converting the number to decimal first.
	decNum := ConvertToDec(numS, fromBase)
	// Converting the number from decimal to the declared base(toBase).
	newBaseNum := ConvertFromDec(decNum, toBase)

	return newBaseNum, nil
}

// ConvertToDec converts an arbitrary base number to decimal base number.
func ConvertToDec(num []rune, fromBase int) float64 {
	// If number base is already decimal - no need to convert.
	if fromBase == 10 {
		numF, err := strconv.ParseFloat(string(num), 64)
		if err != nil {
			log.Fatalln(err)
		}
		return numF
	}
	// Converting the number to decimal base.
	var decNum float64
	iFromEnd := len(num) - 1
	for _, vRune := range num {
		vInt := letterToNum(vRune, fromBase)
		decNum += float64(vInt) * math.Pow(float64(fromBase), float64(iFromEnd))
		iFromEnd--
	}
	return decNum
}

// ConvertFromDec converts decimal base number to an arbitrary base number.
func ConvertFromDec(num float64, toBase int) string {
	// If toBase is decimal - no need to convert.
	if toBase == 10 {
		if flagNeg != 0 {
			// If original number was negative - making new one negative too.
			num *= -1
			flagNeg = 0
		}
		return fmt.Sprint(num)
	}
	// If original number is equal to "0...0".
	if num == 0 {
		return "0"
	}

	// Converting the number to needed base (toBase).
	var newNum string
	for num != 0 {
		reminder := math.Mod(num, float64(toBase))
		num = num / float64(toBase)
		num, _ = math.Modf(num)
		reminderHex := numToLetter(reminder, toBase)
		newNum += reminderHex
	}

	if flagNeg != 0 {
		// If original number was negative - making new one negative too.
		newNum += "-"
		flagNeg = 0
	}
	// Reversing the number.
	newNum = numReverse([]rune(newNum))

	return newNum
}

// letterToNum converts rune(it can be digit(0-9) or letter(a-z,A-Z) to int.
func letterToNum(digit rune, base int) int {
	if base > 10 {
		for i, v := range allDigits[10:] {
			if digit == v {
				return i + 10
			}
		}
	}
	digitInt, err := strconv.Atoi(string(digit))
	if err != nil {
		log.Fatal(err)
	}
	return digitInt
}

// numToLetter converts float64 digit to string with digits(0-9) or/and letters(a-z,A-Z).
func numToLetter(digit float64, base int) string {
	if base > 10 {
		for i, v := range allDigits[10:base] {
			if int(digit) == i+10 {
				return string(v)
			}
		}
	}
	return fmt.Sprint(digit)
}

// numIsValid checks if number belongs to the declared base.
func numIsValid(num []rune, checkBase int) bool {
	for _, v := range num {
		flagFound := false
		for _, d := range allDigits[:checkBase] {
			if v == d {
				flagFound = true
				continue
			}
		}
		if flagFound != true {
			return false
		}
	}
	return true
}

// numReverse reverses the number.
func numReverse(num []rune) string {
	leftSide := 0
	rightSide := len(num) - 1
	for leftSide < len(num)/2 {
		num[leftSide], num[rightSide] = num[rightSide], num[leftSide]
		leftSide++
		rightSide--
	}
	return string(num)
}
