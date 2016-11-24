package utils

import (
	"testing"
	"fmt"
)

func TestPointers(t *testing.T) {
	var i int = 10
	var p *int = &i
	fmt.Println(p)
	fmt.Println(*p)
}
