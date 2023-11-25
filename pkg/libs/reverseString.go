package libs

func ReverseString(s string) string {
	return string(ReverseSlice([]rune(s)))
}
