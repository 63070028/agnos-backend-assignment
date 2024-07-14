package service

import (
	"fmt"
	"math"
	"unicode"

	"github.com/63070028/agnos-backend-assignment/model"
)

func MiminimumActions(password string, config model.ConfigStrongPassword) int {
	actions := 0
	passwordLength := len(password)

	missingL := config.MinLowerCase
	missingU := config.MinUpperCase
	missingD := config.MinDigit

	repeatCount := NumberReplaceCharacter(password, config.Repeat)

	fmt.Println(repeatCount)

	for _, r := range password {
		if unicode.IsLower(r) {
			missingL--
		}
		if unicode.IsUpper(r) {
			missingU--
		}
		if unicode.IsDigit(r) {
			missingD--
		}
	}

	missingTypes := missingL + missingU + missingD

	if passwordLength < config.MinLength {
		actions += int(math.Max(float64(config.MinLength-passwordLength), float64(missingTypes)))
	}

	return actions
}

func NumberReplaceCharacter(text string, repeat int) int {
	result := 0

	for i := 0; i < len(text); i++ {
		if i > repeat-1 {
			temp := text[i]
			counter := 1
			for j := 1; j < repeat; j++ {
				if temp == text[i-j] {
					counter++
				}
				if counter == repeat {
					result++
					i++
				}
			}
		}
	}

	return result
}
