package jumble

import (
	"errors"
	"math"
	"math/rand"
	"time"
)

// Characters are defined as UTF-8 code points.
var codePoints = [][]rune{
	{0x41, 0x5a},
	{0x61, 0x7a},
	{0xc0, 0xd6},
	{0xd8, 0xf6},
	{0xf8, 0xff},
}

// Text takes a string representing a text, identifies the words and
// jumbles the characters of those words.
func Text(str string) (string, error) {
	if len(str) == 0 {
		return "", errors.New("Text has zero length")
	}

	text := []rune(str)
	indices, err := Indices(str)
	if err != nil {
		return str, nil
	}

	var jumbledWord []rune
	for _, i := range indices {
		if i[1]-i[0] > 2 {
			jumbledWord = Word(text[i[0]:i[1]+1], time.Now().Unix())
			for j := i[0] + 1; j < i[1]; j++ {
				text[j] = jumbledWord[j-i[0]]
			}
		}
	}
	return string(text), nil
}

// Indices takes a string and returns a 2-dimensional int array with the start
// and end indices of each word.
func Indices(str string) ([][]int, error) {
	if len(str) == 0 {
		return nil, errors.New("String has zero length")
	}

	runes := []rune(str)
	indices := make([][]int, 0, 8)
	word := false
	var character bool
	var start, end int
	// We compare UTF-8 code points
	for i := 0; i < len(runes); i++ {
		character = false
		for _, s := range codePoints {
			if runes[i] >= s[0] && runes[i] <= s[1] {
				character = true
				break
			}
		}
		if !word && character {
			start = i
			word = true
		} else if word && !character {
			end = i - 1
			word = false
			indices = append(indices, []int{start, end})
		}
	}
	if word {
		indices = append(indices, []int{start, len(runes) - 1})
	}
	if len(indices) == 0 {
		return nil, errors.New("String has no words")
	}
	return indices, nil
}

// Word changes the given rune array in that it keeps the first and the last
// rune and randomly jumbles the runes of sub rune array str[1 : len(str)-1].
func Word(str []rune, seed int64) []rune {
	if len(str) == 0 {
		return nil
	}
	if len(str) < 4 {
		return str
	}

	substr := str[1 : len(str)-1]
	rand.Seed(seed)
	newstr := make([]rune, 0, 8)
	newstr = append(newstr, str[0])
	for len(substr) > 0 {
		// r = random index
		r := int(math.Trunc(float64(len(substr)) * rand.Float64()))
		newstr = append(newstr, substr[r])
		substr = append(substr[0:r], substr[r+1:]...)
	}
	newstr = append(newstr, str[len(str)-1])

	return newstr
}
