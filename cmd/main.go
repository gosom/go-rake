package main

import (
	"log"

	"github.com/gosom/go-rake"
)

var TXT string = `
Herzlich Willkommen auf unserer Website
Liebe Gäste,

Wir haben Betriebsferiern vom Di. den 07. August bis zum Di. den 21. August 2018.
Ab Mittwoch, den 22. August 2018 sind wir wieder für Sie da und freuen uns auf Ihren Busuch!
 

Herzlich Willkommen in unserem Schnellrestaurant. Mit erlesenen Speisen aus dem Fernen Osten und ausgesuchten Spezialitäten aus unserer Heimatregion möchten wir Sie bewirten.

erleben Sie bei uns die einfachen Geheimnisse großer Kochkunst. Durch erstklassige Zutaten sowie Geschick bei der Kombination und bei der Zubereitung werden Kleinigkeiten zu Hochgenüssen. Die Zutat ist unser Weg! Der Zubereitung folgen der Geschmack, die Konsistenz und die Eigenart des Gerichts. Genießen Sie den Geschmack, der aus dem Wasser kommt - Nudeln und Reis, die mehr als nur Beilagen sind - Fleisch, das frittiert, gewokt und gegrillt wird und zum Schluss ein Dessert von reicher Süße.

Wir wünschen Ihnen einen guten Aufenthalt in unserem Haus und wünschen Ihnen einen guten Appetit.

hello world! of warcraft warcraft

Ihr China Haus Team 
`

func main() {
	rake, err := rake.New(rake.STOPWORDS)
	if err != nil {
		log.Fatal(err)
	}
	keywords, err := rake.Extract(TXT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(keywords)
}
