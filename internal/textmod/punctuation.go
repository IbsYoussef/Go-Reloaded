package textmod

import (
	"regexp"
	"strings"
)

func ModifyPunctuationAndQuotes(text string) string {
	re := regexp.MustCompile(`\s*([.,!;:?]+)\s*`)
	text = re.ReplaceAllString(text, "$1 ")

	reStartQuote := regexp.MustCompile(`' `)
	text = reStartQuote.ReplaceAllString(text, " '")

	reEndQuote := regexp.MustCompile(` '`)
	return strings.TrimSpace(reEndQuote.ReplaceAllString(text, "'"))
}
