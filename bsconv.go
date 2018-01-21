// Package bsconv converts any number from an arbitrary base to another arbitrary base.
// From binary up to base62. This package can work with numbers bigger than int64. It also
// contains two additional functions to convert from an arbitrary base to decimal and from
// decimal to an arbitrary base.
package bsconv

import (
	"fmt"
	"math/big"
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

// ConvertToDec converts an arbitrary base number to decimal base number larger than int64.
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
	// After removing "-" we have to check our number again.
	if len(num) == 0 {
		return "", fmt.Errorf("Nothing to convert")
	}
	// Converting the number to decimal base using math/big library in order to get
	// numbers bigger than int64.
	bDecNum := big.NewInt(0)
	iFromEnd := len(num) - 1
	for _, vRune := range num {
		vInt, err := letterToNum(vRune, fromBase)
		if err != nil {
			return "", err
		}
		bVint := big.NewInt(int64(vInt))
		bFromBase := big.NewInt(int64(fromBase))
		bIfromEnd := big.NewInt(int64(iFromEnd))
		bPow := big.NewInt(0)
		mod := big.NewInt(0)
		bRight := big.NewInt(0)
		bDecNum.Add(bDecNum, bRight.Mul(bPow.Exp(bFromBase, bIfromEnd, mod), bVint))
		iFromEnd--
	}
	if flagNeg == 1 {
		// If original number was negative - making new one negative too.
		bDecNum.Neg(bDecNum)
	}

	return bDecNum.String(), nil
}

// ConvertFromDec converts decimal base number (larger than int64) to an arbitrary base number.
func ConvertFromDec(numS string, toBase int) (string, error) {
	if toBase < 2 {
		return "", fmt.Errorf("Invalid base: %v", toBase)
	}
	if numS == "" {
		return "", fmt.Errorf("Nothing to convert")
	}
	// Converting string to the number bigger than int64.
	bNum := big.NewInt(0)
	bNum, ok := bNum.SetString(numS, 10)
	if ok != true {
		return "", fmt.Errorf("Invalid number: %v", numS)
	}
	bZero := big.NewInt(0)
	// If original number is equal to "0...0".
	if bNum.Cmp(bZero) == 0 {
		return "0", nil
	}
	// If the number is negativ: flagNeg = 1 and "saving" the negative sign.
	if bNum.Cmp(bZero) == -1 {
		flagNeg = 1
		bNum.Neg(bNum)
	}
	// Converting the number to needed base (toBase).
	var newNum string
	bToBase := big.NewInt(int64(toBase))
	bReminder := big.NewInt(0)
	for bNum.Cmp(bZero) != 0 {
		bNum.DivMod(bNum, bToBase, bReminder)
		reminderHex := numToLetter(bReminder.Int64(), toBase)
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

// letterToNum converts rune(it can be digit(0-9) or letter(a-z,A-Z)) to int.
func letterToNum(digit rune, base int) (int, error) {
	if base > 10 {
		for i, v := range allDigits[10:base] {
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
