package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"

	html "golang.org/x/net/html"
)

const HTMLTemplate = `
<!DOCTYPE html>
<html lang="en">
	<head>
	<meta charset="utf-8">
	<meta name="color-scheme" content="light dark" />
	<link rel="manifest" href="/manifest.json" crossOrigin="use-credentials">
	</head>
	<body>
		<h1>Hello world!</h1>
		<p>Lemme try this shit</p>
		<p>Lemme try again</p>
		<p>anaa again</p>
		<p>again...</p>
		<p>And figure out how to deal with it.</p>
		<img src="test.png" />
	</body>
	<script>console.log("Test log")</script>
</html>
`

func visitor(node *html.Node, PwordsCount, PimgsCount *int) {
	if node.Type == html.TextNode {
		*PwordsCount += len(strings.Fields(node.Data))

	} else if node.Type == html.ElementNode && node.Data == "img" {
		*PimgsCount++
	}

	for count := node.FirstChild; count != nil; count = count.NextSibling {
		visitor(count, PwordsCount, PimgsCount)
	}
}

func analyzeDOM(doc *html.Node) (int, int) {
	var wordsCount, imgsCount int

	visitor(doc, &wordsCount, &imgsCount)
	return wordsCount, imgsCount
}

func main() {
	doc, err := html.Parse(bytes.NewReader([]byte(HTMLTemplate)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Parse failed: %s\n", err)
		os.Exit(-1)
	}

	wordsCount, imgsCount := analyzeDOM(doc)
	fmt.Printf("%d word/s | %d image/s", wordsCount, imgsCount)

}



