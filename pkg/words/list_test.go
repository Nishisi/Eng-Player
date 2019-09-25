package words

import (
	"fmt"
	"testing"
)

func TestLoadWordList(t *testing.T) {
	wl := &WordList{}
	_ = wl.LoadWordList("./_testdata/list")
	fmt.Println(wl)
}
