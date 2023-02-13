package html_read

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type product struct {
	Name           string
	Price          string
	Stars          float64
	Reviews        int
	Description    string
	NumberOfPeople int
}

var dataProducts = []product{
	{"Strawberries", "$2.00", 4.0, 251, "Lorem ipsum dolor sit amet, consectetur adipiscing elit.", 5},
	{"Onions", "$2.80", 5.0, 123, "Morbi sit amet erat vitae purus consequat vehicula nec sit amet purus.", 4},
	{"Tomatoes", "$3.10", 4.5, 235, "Curabitur tristique odio et nibh auctor, ut sollicitudin justo condimentum.", 3},
	{"Courgette", "$1.20", 4.0, 251, "Phasellus at leo a purus consequat ornare ac aliquam quam.", 2},
	{"Broccoli", "$3.80", 3.5, 123, "Maecenas sed ante sagittis, dapibus dui quis, hendrerit orci.", 2},
	{"Potatoes", "$3.00", 2.5, 235, "Vivamus malesuada est et tellus porta, vel consectetur orci dapibus.", 1},
}

func starRating(inputStars, inputStars2 float64, stars []product) {
	for i := 0; i < len(dataProducts); i++ {
		if inputStars <= dataProducts[i].Stars && inputStars2 >= dataProducts[i].Stars {
			stars = append(stars, dataProducts[i])
		}
	}
	fmt.Printf(`Here is your: %v`, stars)
}

func priceRating(inputPrice, inputPrice2 string, prices []product) {
	for i := 0; i < len(dataProducts); i++ {
		if inputPrice <= dataProducts[i].Price && inputPrice2 >= dataProducts[i].Price {
			prices = append(prices, dataProducts[i])
		}
	}
	fmt.Printf(`Here is your: %v`, prices)
}
func Filter_words() {

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

	var input string
	fmt.Println("What do you want to search?")
	fmt.Scan(&input)

	f := func(i int, sel *goquery.Selection) bool {
		//We retrieve all words starting with 'w'.
		return strings.HasPrefix(sel.Text(), input)
	}

	var words []string
	/*
		This is a predicate function that returns a boolean true for all words that begin with 'w'.
	*/
	doc.Find("li").FilterFunction(f).Each(func(_ int, sel *goquery.Selection) {
		words = append(words, sel.Text())
	})

	fmt.Println(words)

	var prices []product
	var stars []product
	var inputPrice, inputPrice2 string
	var inputStars, inputStars2 float64

	fmt.Println("Type the range of your price")
	fmt.Scanln(&inputPrice)
	fmt.Scanln(&inputPrice2)
	priceRating(inputPrice, inputPrice2, prices)

	fmt.Println()

	fmt.Println("Type the range of your stars")
	fmt.Scanln(&inputStars)
	fmt.Scanln(&inputStars2)
	starRating(inputStars, inputStars2, stars)
}
