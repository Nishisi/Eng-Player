package aruku

import (
	"fmt"
	"testing"
)

func TestScrap(t *testing.T) {
	s, _ := scrap("respond")
	fmt.Println(s)
	s, _ = scrap("startling")
	fmt.Println(s)
	s, _ = scrap("day+one")
	fmt.Println(s)
}
