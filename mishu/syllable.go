package mishu
import (
	"regexp"
)

var vowels = []byte{'a', 'e', 'i', 'o', 'u'}

func IsVowel(r byte) bool {
	for _, v := range vowels {
		if v == r {
			return true
		}
	}

	return false
}

func IsConsonant(r byte) bool {
	return !IsVowel(r)
}

func Syllables(word string) []string {
	for i, r := range []byte(word) {
		if i > 0 {
			prev := word[i-1]

			next := byte('0')
			if (i < len(word) - 1) {
				next = word[i+1]
			}

			if IsConsonant(r) {
				if (r == 'c' || r == 'g') && (next == 'h' || next == 'i' || next == 'e') {
					return append([]string{word[0:i]}, Syllables(word[i:])...)
				} else if IsVowel(prev) && regexp.MustCompile("^.(c|g)(i|e|hi|he)").MatchString(word[i:]) {
					return append([]string{word[0:i+1]}, Syllables(word[i+1:])...)
				} else if IsVowel(prev) && IsVowel(next) {
					return append([]string{word[0:i]}, Syllables(word[i:])...)
				} else if IsVowel(prev) && regexp.MustCompile("^(b|c|d|f|g|h|p|t|v)(l|r)").MatchString(word[i:]) {
					return append([]string{word[0:i]}, Syllables(word[i:])...)
				} else if IsVowel(prev) && IsConsonant(next) && i < len(word) - 2 && IsVowel(word[i + 2]) {
					return append([]string{word[0:i+1]}, Syllables(word[i+1:])...)
				} else if IsVowel(prev) && i < len(word) - 3 && IsConsonant(next) && IsConsonant(word[i + 2]) && IsVowel(word[i + 3]) {
					return append([]string{word[0:i+1]}, Syllables(word[i+1:])...)
				}
			}
		}
	}

	return []string{word}
}
