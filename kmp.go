package main

func MatchPatternKMP(txt, pattern string) []int {
	result := []int{}
	if len(txt) < len(pattern) {
		return result
	}

	lps := buildLPS(pattern)
	txtIdx := 0
	patIdx := 0

	for txtIdx < len(txt) {
		if txt[txtIdx] == pattern[patIdx] {
			if patIdx == len(pattern)-1 {
				result = append(result, txtIdx-patIdx)
				patIdx = 0
			} else {
				patIdx++
				txtIdx++
			}
		} else if patIdx == 0 {
			txtIdx++
		} else {
			patIdx = lps[patIdx-1]
		}
	}
	return result
}

func buildLPS(pattern string) []int {
	lps := make([]int, len(pattern))
	if len(pattern) == 0 {
		return lps
	}

	lps[0] = 0
	lenPrevLPS := 0
	idx := 1

	for idx < len(lps) {
		if pattern[idx] == pattern[lenPrevLPS] {
			lenPrevLPS++
			lps[idx] = lenPrevLPS
			idx++
		} else {
			if lenPrevLPS > 0 {
				lenPrevLPS = lps[lenPrevLPS-1]
			} else {
				lps[idx] = 0
				idx++
			}
		}
	}
	return lps
}
