package textjustification

import (
	"strings"
)

func fullJustify(words []string, maxWidth int) []string {
	var res []string

	for i := 0; i < len(words); i++ {
		l, j := 0, i // l is counter for length of words
		for ; j < len(words); j++ {
			if l+len(words[j])+j-i > maxWidth { // j-i is spaces between words
				break // check if words with spaces between suit into row
			}

			l += len(words[j])
		}

		count := j - i - 1        // count of spaces between words
		spaces := maxWidth - l    // left spaces to fill row
		isLast := j == len(words) // is the last row
		row := make([]byte, 0, maxWidth)

		if count == 0 || isLast { // should be left-justifed
			row = append(row, []byte(strings.Join(words[i:j], " "))...)
			row = append(row, []byte(strings.Repeat(" ", spaces-count))...)
		} else {
			estimateSpace := spaces / count
			rem := spaces % count
			if rem == 0 {
				sep := strings.Repeat(" ", estimateSpace)
				row = append(row, []byte(strings.Join(words[i:j], sep))...)
			} else {
				sep := strings.Repeat(" ", estimateSpace+1)                                    // add extra space to align
				row = append(row, []byte(strings.Join(words[i:i+rem+1], sep))...)              // i+rem+1 because, words one more than spaces
				row = append(row, []byte(sep[:len(sep)-1])...)                                 // add new delimeter
				row = append(row, []byte(strings.Join(words[i+rem+1:j], sep[:len(sep)-1]))...) // join left words
			}
		}

		res = append(res, string(row))
		i = j - 1
	}

	return res
}
