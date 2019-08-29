package aruku

// Aruku
type Aruku struct {
	soundURL string
}

func New(soundURL string) *Aruku {
	return &Aruku{soundURL: soundURL}
}

// GetMeanAndFile is to get mean of word and audio file.
func (a *Aruku) GetMeanAndFile() (string, string) {
	return "", ""
}
