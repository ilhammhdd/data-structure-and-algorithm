package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildLPS(t *testing.T) {
	assert.Equal(t, []int{0, 0, 1, 2, 3, 1}, buildLPS("ababaa"))
}
func TestBuildLPS_OneUniqueEnd(t *testing.T) {
	assert.Equal(t, []int{0, 0, 1, 2, 0}, buildLPS("ababc"))
}
func TestBuildLPS_OneUniqueMid(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0, 1, 2}, buildLPS("abcab"))
}
func TestBuildLPS_TwoGroups(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 0, 0, 0}, buildLPS("aaabbb"))
}
func TestBuildLPS_Distinct(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0, 0, 0, 0}, buildLPS("abcdef"))
}
func TestBuildLPS_AllEqual(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2}, buildLPS("aaa"))
}
func TestBuildLPS_OneChar(t *testing.T) {
	assert.Equal(t, []int{0}, buildLPS("a"))
}
func TestBuildLPS_Empty(t *testing.T) {
	assert.Empty(t, buildLPS(""))
}

func TestBuildLPS_A(t *testing.T) {
	assert.Equal(t, []int{0, 0, 0, 0, 1, 2, 3, 4}, buildLPS("abcdabcd"))
}

func TestMatchPatternKMP(t *testing.T) {
	assert.Equal(t, []int{0, 2}, MatchPatternKMP("abab", "ab"))
}
func TestMatchPatternKMP_Overlap(t *testing.T) {
	assert.Equal(t, []int{0, 9, 12}, MatchPatternKMP("aabaacaadaabaaba", "aaba"))
}
func TestMatchPatternKMP_Repeating(t *testing.T) {
	assert.Equal(t, []int{0, 4}, MatchPatternKMP("abcdabcd", "abcd"))
}
func TestMatchPatternKMP_AllMatchEven(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4}, MatchPatternKMP("aaaaaa", "aa"))
}
func TestMatchPatternKMP_AllMatchOdd(t *testing.T) {
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, MatchPatternKMP("aaaaaaa", "aa"))
}
func TestMatchPatternKMP_NoMatch(t *testing.T) {
	assert.Empty(t, MatchPatternKMP("abcdef", "gh"))
}
func TestMatchPatternKMP_PatternLonger(t *testing.T) {
	assert.Empty(t, MatchPatternKMP("abc", "abcd"))
}

func TestSlice(t *testing.T) {

}
