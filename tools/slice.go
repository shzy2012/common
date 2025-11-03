package tools

// Distinct removes duplicate elements from a slice while preserving order.
// It works with any comparable type.
func Distinct[T comparable](slice []T) []T {
	if len(slice) == 0 {
		return slice
	}

	keys := make(map[T]bool, len(slice))
	result := make([]T, 0, len(slice))

	for _, entry := range slice {
		if !keys[entry] {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return result
}

// Contains checks if a value exists in a slice.
// It works with any comparable type.
func Contains[T comparable](value T, slice []T) bool {
	for _, ele := range slice {
		if ele == value {
			return true
		}
	}
	return false
}

func InArray[T comparable](slice []T, item T) bool {
	for _, ele := range slice {
		if ele == item {
			return true
		}
	}
	return false
}

func RemoveArray[T comparable](slice []T, item T) []T {
	for i, ele := range slice {
		if ele == item {
			slice = append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}

// Reverse reverses the order of elements in a slice.
// The original slice is modified in place.
func Reverse[T any](arr []T) []T {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

// ReverseCopy creates a new slice with reversed order.
// The original slice remains unchanged.
func ReverseCopy[T any](arr []T) []T {
	result := make([]T, len(arr))
	for i, j := 0, len(arr)-1; i < len(arr); i, j = i+1, j-1 {
		result[i] = arr[j]
	}
	return result
}
