package goyml_test

import (
	"testing"

	"github.com/hypnguyen1209/goyml"
)

func TestParse(t *testing.T) {
	t.Parallel()
	data := []byte(
		`
foo: bar
num: 6
`)
	yq := goyml.Parse(data)
	foo, err := yq.String("foo")
	if err != nil {
		t.Fatalf("failed to parse string %v", foo)
	}
	num, err := yq.Int("num")
	if err != nil {
		t.Fatalf("failed to parse string %v", num)
	}
}
