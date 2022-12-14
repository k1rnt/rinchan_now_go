package main

import "github.com/k1rnt/rinchan_now_go/pkg/gen_rn"

func main() {
	str := "rinchan now! rinchan now! rinchan rinchan rinchan now!"
	generator := gen_rn.NewRinchanGenerator(str)
	generator.Generate()
}
