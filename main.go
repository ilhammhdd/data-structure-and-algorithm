package main

import (
	"fmt"
	"io/fs"
	"math/rand"
	"os"
)

func main() {
	numFreq := make(map[int32]int32)

	const n = 5000_000
	const maxFreq = 100

	countMills := 1
	result := make([]byte, 0, n)
	var randomized int32
	for i := 0; i < n; i++ {
		if i%1_000_000 == 0 {
			fmt.Printf("generated %d\n", 1_000_000*countMills)
			countMills++
		}
		randomized = rand.Int31n(n)
		for freq, ok := numFreq[randomized]; ok; {
			if freq < maxFreq {
				numFreq[randomized]++
			} else {
				randomized = rand.Int31n(n)
			}
		}
		result = append(result, []byte(fmt.Sprintf(",%d", randomized))...)
	}

	os.WriteFile("random-numbers.txt", result, fs.FileMode(os.O_WRONLY|os.O_RDONLY))
}
