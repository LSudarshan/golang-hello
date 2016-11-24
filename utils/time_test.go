package utils

import (
	"time"
	"testing"
	"fmt"
)

func TestToday(t *testing.T){
	fmt.Printf("Today: %s \n", time.Now())
	fmt.Printf("Tomorrow: %s \n", time.Now().AddDate(0, 0, 1))
}
