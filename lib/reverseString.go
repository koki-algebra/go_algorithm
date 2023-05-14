package lib

func ReverseString(s string) string {
	return string(ReverseSlice([]rune(s)))
}
