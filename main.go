package main

import (
	"fmt"
	"os"

	"github.com/nl-interpreter-engine/pkg/sentiment"
	"github.com/nl-interpreter-engine/web"
)

func main() {
	webClient := web.NewWebClient(nil)

	sentimentClient := sentiment.New(webClient.HTTPClient)

	file := os.Args[1]

	bytes, err := sentimentClient.InterpretSentiment(file)

	if err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}

	fmt.Printf("%s\n", bytes)
}
