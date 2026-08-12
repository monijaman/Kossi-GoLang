package productreview

import (
	"html"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var reviewTagPattern = regexp.MustCompile(`<[^>]*>`)

// moderateReview accepts plain-text user reviews only. The blocked list can
// be extended without a code change using REVIEW_BLOCKED_WORDS (comma-separated).
func moderateReview(input string) (string, string) {
	text := html.UnescapeString(input)
	text = reviewTagPattern.ReplaceAllString(text, " ")
	text = strings.Map(func(r rune) rune {
		if unicode.IsControl(r) && r != '\n' && r != '\r' && r != '\t' {
			return -1
		}
		return r
	}, text)
	text = strings.Join(strings.Fields(text), " ")
	if len([]rune(text)) < 3 {
		return "", "Comment must contain at least 3 characters"
	}
	if len([]rune(text)) > 5000 {
		return "", "Comment must be 5000 characters or fewer"
	}

	blocked := os.Getenv("REVIEW_BLOCKED_WORDS")
	if blocked == "" {
		blocked = "fuck,shit,bitch,cunt,asshole,slut,whore"
	}
	lower := strings.ToLower(text)
	for _, word := range strings.Split(blocked, ",") {
		word = strings.TrimSpace(strings.ToLower(word))
		if word != "" && strings.Contains(lower, word) {
			return "", "Comment contains language that is not allowed"
		}
	}
	return text, ""
}
