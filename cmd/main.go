package main

import (
	"fmt"
	"os"

	"github.com/k1rnt/rinchan_now_go/pkg/rinchannow"
)

const SIZE = 30000

func main() {
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: go run cmd/*.go <filename>")
		return
	}
	filename := args[1]
	b, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	program := string(b)

	rinchan := rinchannow.NewRinchanRuntime(SIZE)
	rinchan.Now(program)
}
