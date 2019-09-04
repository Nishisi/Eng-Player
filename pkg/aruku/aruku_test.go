package aruku

import (
	"fmt"
	"testing"
)

func TestScrap(t *testing.T) {
	s, _ := scrap("infer")
	fmt.Println(s)
}
