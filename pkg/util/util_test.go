/*
   Copyright 2021 VMware, Inc.
   SPDX-License-Identifier: Apache-2.0
*/

package util

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	allowedLowerLettersList := strings.Split(VCDLowerLetters, "")
	allowedUpperLettersList := strings.Split(VCDUpperLetters, "")
	allowedDigitsList := strings.Split(VCDDigits, "")
	allowedSymbolsList := strings.Split(VCDSymbols, "")

	lst := append(allowedLowerLettersList, allowedUpperLettersList...)
	lst = append(lst, allowedDigitsList...)
	lst = append(lst, allowedSymbolsList...)

	allowedNumDigits := 2
	allowedNumSymbols := 3
	passwd, err := GeneratePassword(10, allowedNumDigits, allowedNumSymbols, false, false)
	assert.NoError(t, err, "should be able to generate a password with the character set")
	fmt.Printf("password is [%s]", passwd)

	charSet := NewSet(lst)
	digitsSet := NewSet(allowedDigitsList)
	symbolsSet := NewSet(allowedSymbolsList)
	numDigits := 0
	numSymbols := 0

	for _, c := range strings.Split(passwd, "") {
		if !charSet.Contains(c) {
			assert.Fail(t, "character set [%v] does not contain character [%v]", charSet, c)
		}
		if digitsSet.Contains(c) {
			numDigits++
		}
		if symbolsSet.Contains(c) {
			numSymbols++
		}
	}

	assert.LessOrEqual(t, allowedNumDigits, numDigits, "password has unexpected number of digits")
	assert.LessOrEqual(t, allowedNumSymbols, numSymbols, "password has unexpected number of symbols")

	return
}
