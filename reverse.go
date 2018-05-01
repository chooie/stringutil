package stringutil

// Reverse : reverses a string
func Reverse(originalString string) string {
	runeArrayToReverse := toRuneArray(originalString)
	inlineReverseRuneArray(runeArrayToReverse)
	return string(runeArrayToReverse)
}

func inlineReverseRuneArray(runeArray []rune) {
	middleOfArray := len(runeArray) / 2
	formerIndex := 0
	latterIndex := len(runeArray) - 1
	for formerIndex < middleOfArray {
		runeArray[formerIndex], runeArray[latterIndex] = runeArray[latterIndex], runeArray[formerIndex]
		formerIndex++
		latterIndex--
	}
}

func toRuneArray(theString string) []rune {
	return []rune(theString)
}
