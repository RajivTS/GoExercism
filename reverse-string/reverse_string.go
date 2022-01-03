package reverse

func Reverse(input string) string {
	inputRunes := []rune(input)
	for left, right := 0, len(inputRunes)-1; left < right; left, right = left+1, right-1 {
		tmp := inputRunes[left]
		inputRunes[left] = inputRunes[right]
		inputRunes[right] = tmp
	}
	return string(inputRunes)
}
