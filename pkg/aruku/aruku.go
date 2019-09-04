package aruku

import "github.com/PuerkitoBio/goquery"

// Aruku
type Aruku struct {
	soundURL string
}

func New(soundURL string) *Aruku {
	return &Aruku{soundURL: soundURL}
}

// GetMeanAndFile is to get mean of word and audio file.
func (a *Aruku) GetMeanAndFile(word string) (string, string) {
	s, _ := scrap(word)
	return s, ""
}

// scraping
func scrap(w string) (string, error) {
	url := "https://eow.alc.co.jp/search?q="
	url += w
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	selection := doc.Find("div#resultsList")

	var title string
	selection.Find("li").Each(func(i int, s *goquery.Selection) {
		if i == 0 {
			title = s.Find("div").Text()
		}
	})

	return title, nil
}
