package main

import (
	"flag"

	"github.com/k1rnt/rinchan_now_go/pkg/gen_rn"
)

func main() {
	var str string
	flag.StringVar(&str, "str", "", "string to rn code")
	flag.Parse()

	generator := gen_rn.NewRinchanGenerator(str)
	generator.Generate()
}
