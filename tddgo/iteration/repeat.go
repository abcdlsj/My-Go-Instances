package iteration

func Repeat(character string, repeatCount int) string {
	res := ""
	for i := 0; i < repeatCount; i++ {
		res += character
	}
	return res
}
