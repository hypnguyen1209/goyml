package main

import (
	"os"

	"github.com/hypnguyen1209/goyml"
)

func main() {
	test, _ := os.ReadFile("./test.yml")
	yq := goyml.Parse(test)
	d1, _ := yq.Int("age")
	d2, _ := yq.String("name")
	d3, _ := yq.ArrayOfStrings("work")
	println(d2, d1)
	for _, v := range d3 {
		println(v)
	}
}
