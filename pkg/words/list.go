package words

import (
	"encoding/json"
	"os"
)

// Word List
type WordList struct {
	Lists []struct {
		Word        string `json:"word"`
		Translation string `json:"translation"`
	} `json:"lists"`
}

// load
func (wl *WordList) loadListFile(file string) error {
	listFile, err := os.Open(file)
	defer listFile.Close()
	if err != nil {
		return err
	}
	jsonParser := json.NewDecoder(listFile)
	jsonParser.Decode(wl)
	return nil
}
