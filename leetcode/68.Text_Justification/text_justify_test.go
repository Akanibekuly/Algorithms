package textjustification

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFullJustify(t *testing.T) {
	cases := []struct {
		name     string
		words    []string
		maxWidth int
		res      []string
	}{
		// {"one word", []string{"one"}, 5, []string{"one  "}},
		// {"1", []string{"This", "is", "an", "example", "of", "text", "justification."}, 16,
		// 	[]string{
		// 		"This    is    an",
		// 		"example  of text",
		// 		"justification.  "}},
		// {"2", []string{"What", "must", "be", "acknowledgment", "shall", "be"}, 16, []string{
		// 	"What   must   be",
		// 	"acknowledgment  ",
		// 	"shall be        ",
		// }},
		{
			"hard",
			[]string{"Science", "is", "what", "we", "understand", "well", "enough", "to", "explain", "to", "a", "computer.", "Art", "is", "everything", "else", "we", "do"},
			20,
			[]string{
				"Science  is  what we",
				"understand      well",
				"enough to explain to",
				"a  computer.  Art is",
				"everything  else  we",
				"do                  ",
			},
		},
		{
			"fallen test",
			[]string{"ask", "not", "what", "your", "country", "can", "do", "for", "you", "ask", "what", "you", "can", "do", "for", "your", "country"},
			16,
			[]string{
				"ask   not   what",
				"your country can",
				"do  for  you ask",
				"what  you can do",
				"for your country",
			},
		},
	}

	for _, c := range cases {
		require.Equal(t, c.res, fullJustify(c.words, c.maxWidth))
	}
}
