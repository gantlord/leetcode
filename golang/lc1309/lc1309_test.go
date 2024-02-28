package lc1309

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

// Generate predictable input
func generatePredictableInput(size int) string {
	var result strings.Builder
	for i := 0; i < size; i++ {
		if i%2 == 0 {
			result.WriteString(fmt.Sprintf("%d", 1+rand.Intn(9)))
		} else {
			result.WriteString(fmt.Sprintf("%d#", 10+rand.Intn(17)))
		}
	}
	return result.String()
}

// Generate unpredictable input
func generateUnpredictableInput(size int) string {
	rand.Seed(time.Now().UnixNano())
	var result strings.Builder
	for i := 0; i < size; i++ {
		if rand.Intn(2) == 0 {
			result.WriteString(fmt.Sprintf("%d", 1+rand.Intn(9)))
		} else {
			result.WriteString(fmt.Sprintf("%d#", 10+rand.Intn(17)))
		}
	}
	return result.String()
}

func BenchmarkFreqAlphabetsPredictable(b *testing.B) {
	input := generatePredictableInput(10000) // Adjust size for your needs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		freqAlphabets(input)
	}
}

func BenchmarkFreqAlphabetsUnpredictable(b *testing.B) {
	input := generateUnpredictableInput(10000) // Adjust size for your needs
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		freqAlphabets(input)
	}
}
