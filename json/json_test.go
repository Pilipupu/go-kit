package json

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	s := Get(`{"a": "1213"}`, "a").String()
	fmt.Println(s)
}
