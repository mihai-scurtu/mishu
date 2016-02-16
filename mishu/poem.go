package mishu
import (
	"io/ioutil"
	"strings"
	"regexp"
	"path/filepath"
	"math/rand"
)

type Poem []string;

func GeneratePoem(words []string, count int) Poem {
	poem := make(Poem, 0)

	_, r := randomWord(words)
	word := words[r+1]
	poem = append(poem, words[r], word)

	for len(poem) < count {
		word = pickNext(word, words)

		poem = append(poem, word)
	}

	return poem
}

func GetWordList(pattern string) []string {
	files, _ := filepath.Glob(pattern)
	words := make([]string, 0)

	//	ch := make(chan []string)

	for _, fileName := range files {
		fromFile := getWordsFromFile(fileName)
		words = append(words, fromFile...)
	}

	//	for {
	//		append(words, <-ch...)
	//	}

	return words
}

func getWordsFromFile(fileName string) []string {
	byteContent, _ := ioutil.ReadFile(fileName)
	content := strings.ToLower(string(byteContent))

	// normalize whitespace
	r := regexp.MustCompile("\\s+")
	content = r.ReplaceAllLiteralString(content, " ")

	return strings.Split(content, " ")
}

func randomWord(words []string) (word string, pos int) {
	r := rand.Intn(len(words))

	return words[r], r
}

func pickNext(word string, words []string) string {
	var next string

	matches := make([]string, 0)

	for i, w := range words {
		if i < len(words)-1 && w == word {
			matches = append(matches, words[i+1])
		}
	}

	if len(matches) == 0 {
		next, _ = randomWord(words)
	} else {
		next, _ = randomWord(matches)
	}

	return next
}