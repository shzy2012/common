package tools

import (
	"reflect"
	"testing"
)

func TestDistinct(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "no duplicates",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "with duplicates",
			input:    []int{1, 2, 2, 3, 3, 4, 5, 5},
			expected: []int{1, 2, 3, 4, 5},
		},
		{
			name:     "all duplicates",
			input:    []int{1, 1, 1, 1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Distinct(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Distinct() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDistinctString(t *testing.T) {
	input := []string{"a", "b", "a", "c", "b", "d"}
	expected := []string{"a", "b", "c", "d"}

	result := Distinct(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Distinct() = %v, want %v", result, expected)
	}
}

func TestContains(t *testing.T) {
	tests := []struct {
		name     string
		slice    []int
		value    int
		expected bool
	}{
		{
			name:     "value exists",
			slice:    []int{1, 2, 3, 4, 5},
			value:    3,
			expected: true,
		},
		{
			name:     "value does not exist",
			slice:    []int{1, 2, 3, 4, 5},
			value:    6,
			expected: false,
		},
		{
			name:     "empty slice",
			slice:    []int{},
			value:    1,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.value, tt.slice)
			if result != tt.expected {
				t.Errorf("Contains() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestContainsString(t *testing.T) {
	slice := []string{"apple", "banana", "cherry"}

	if !Contains("banana", slice) {
		t.Error("Contains() should return true for existing string")
	}

	if Contains("orange", slice) {
		t.Error("Contains() should return false for non-existing string")
	}
}

func TestReverse(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "even length",
			input:    []int{1, 2, 3, 4},
			expected: []int{4, 3, 2, 1},
		},
		{
			name:     "odd length",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{5, 4, 3, 2, 1},
		},
		{
			name:     "empty slice",
			input:    []int{},
			expected: []int{},
		},
		{
			name:     "single element",
			input:    []int{1},
			expected: []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a copy for testing since Reverse modifies in place
			input := make([]int, len(tt.input))
			copy(input, tt.input)

			result := Reverse(input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Reverse() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestReverseCopy(t *testing.T) {
	input := []int{1, 2, 3, 4, 5}
	expected := []int{5, 4, 3, 2, 1}

	result := ReverseCopy(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("ReverseCopy() = %v, want %v", result, expected)
	}

	// Original slice should remain unchanged
	if !reflect.DeepEqual(input, []int{1, 2, 3, 4, 5}) {
		t.Error("ReverseCopy() should not modify original slice")
	}
}

func TestReverseString(t *testing.T) {
	input := []string{"a", "b", "c", "d"}
	expected := []string{"d", "c", "b", "a"}

	result := Reverse(input)
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Reverse() = %v, want %v", result, expected)
	}
}

func BenchmarkDistinct(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i % 100 // Creates some duplicates
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Distinct(input)
	}
}

func BenchmarkContains(b *testing.B) {
	slice := make([]int, 1000)
	for i := range slice {
		slice[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		Contains(999, slice)
	}
}

func BenchmarkReverse(b *testing.B) {
	input := make([]int, 1000)
	for i := range input {
		input[i] = i
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Create a copy for each iteration since Reverse modifies in place
		testInput := make([]int, len(input))
		copy(testInput, input)
		Reverse(testInput)
	}
}
