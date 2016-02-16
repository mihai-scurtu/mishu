package main

import (
	"fmt"
	"math/rand"
	"time"
	"github.com/mihai-scurtu/mishu/mishu"
	"net/http"
	"os"
	"log"
)



func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	rand.Seed(time.Now().UnixNano())
	words := mishu.GetWordList("./data/raw/*")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		target := 64
		poem := mishu.GeneratePoem(words, target)

		measure := 8
		syl := 0
		for _, word := range poem {
			if syl > measure {
				syl = 0;
				fmt.Fprint(w, "\n")
			}

			fmt.Fprintf(w, "%s ", word)
			syl += len(mishu.Syllables(word));
		}

		fmt.Fprint(w, "\n")
	});

	log.Fatal(http.ListenAndServe(":" + port, nil))
}
