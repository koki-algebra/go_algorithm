package libs

func Chmin(a *int, b int) bool {
	if *a > b {
		*a = b
		return true
	}
	return false
}

func Chmax(a *int, b int) bool {
	if *a < b {
		*a = b
		return true
	}
	return false
}
