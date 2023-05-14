package lib

func ReverseSlice[T any](slice []T) []T {
	new := make([]T, len(slice))
	copy(new, slice)
	for i, j := 0, len(new)-1; i < j; i, j = i+1, j-1 {
		new[i], new[j] = new[j], new[i]
	}
	return new
}
