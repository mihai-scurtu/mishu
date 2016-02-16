package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/mihai-scurtu/mishu/mishu"
)



func main() {
	rand.Seed(time.Now().Unix())
	words := mishu.GetWordList("./data/raw/*")

	target := 64
	poem := mishu.GeneratePoem(words, target)

	measure := 8
	syl := 0
	for _, word := range poem {
		if syl > measure {
			syl = 0;
			fmt.Print("\n")
		}

		fmt.Printf("%s ", word)
		syl += len(mishu.Syllables(word));
	}

	fmt.Print("\n")
}
