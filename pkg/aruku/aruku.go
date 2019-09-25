package aruku

import "github.com/PuerkitoBio/goquery"

// An Aruku represents an alc.
type Aruku struct {
	soundURL string
}

// New returns a new Aruku given soundURL.
func New(soundURL string) *Aruku {
	return &Aruku{soundURL: soundURL}
}

// GetMeanAndFile is to get mean of word and audio file.
func (a *Aruku) GetMeanAndFile(word string) (string, error) {
	s, err := scrap(word)
	return s, err
}

// スクレイピングします
func scrap(w string) (string, error) {
	url := "https://eow.alc.co.jp/search?q="
	url += w
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return "", err
	}

	// resultsList の 最初の li タグ部分のみをとってくる
	selection := doc.Find("div#resultsList li").First()

	title := ""

	// TODO : html タグで ol で囲まれてても liで囲まれていないやつがある。その場合、要素が取れないので修正する必要あり
	selection.Find("span.wordclass,li").Each(func(i int, s *goquery.Selection) {
		if s.Is("li") {
			title += " " + s.Text() + "\n"
		} else {
			title += "\n" + s.Text() + "\n"
		}
	})

	return title, nil
}
