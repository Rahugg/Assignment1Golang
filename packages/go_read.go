package html_read

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func HTML_read() {

	data := `
<html lang="en">
<body>
<p>List of words</p>
<ul>
    <li>dark</li>
    <li>smart</li>
    <li>war</li>
    <li>cloud</li>
    <li>park</li>
    <li>cup</li>
    <li>worm</li>
    <li>water</li>
    <li>rock</li>
    <li>warm</li>
</ul>
<footer>footer for words</footer>
</body>
</html>
`

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(data))

	if err != nil {
		log.Fatal(err)
	}

	words := doc.Find("li").Map(func(i int, sel *goquery.Selection) string {
		return fmt.Sprintf("%d: %s", i+1, sel.Text())
	})

	fmt.Println(words)
}
