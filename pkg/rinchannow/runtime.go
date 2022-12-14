package rinchannow

import (
	"fmt"
	"strings"
)

type RinchanRuntime struct {
	memory  []byte
	pointer int
}

func NewRinchanRuntime(size int) *RinchanRuntime {
	return &RinchanRuntime{
		memory:  make([]byte, size),
		pointer: 0,
	}
}

func (r *RinchanRuntime) Now(code string) {
	c := strings.Split(code, "")
	for i := 0; i < len(c); i++ {
		switch c[i] {
		case string(next):
			r.next()
		case string(prev):
			r.prev()
		case string(incr):
			r.incr()
		case string(decr):
			r.decr()
		case string(putchar):
			r.putchar()
		case string(getchar):
			r.getchar()
		case string(loopstart):
			if r.memory[r.pointer] == 0 {
				loop := 1
				for loop > 0 {
					i++
					if c[i] == string(loopstart) {
						loop++
					} else if c[i] == string(loopend) {
						loop--
					}
				}
			}
		case string(loopend):
			if r.memory[r.pointer] != 0 {
				loop := 1
				for loop > 0 {
					i--
					if c[i] == string(loopstart) {
						loop--
					} else if c[i] == string(loopend) {
						loop++
					}
				}
			}
		}
	}
}

func (r *RinchanRuntime) next() {
	r.pointer++
}

func (r *RinchanRuntime) prev() {
	r.pointer--
}

func (r *RinchanRuntime) incr() {
	r.memory[r.pointer]++
}

func (r *RinchanRuntime) decr() {
	r.memory[r.pointer]--
}

func (r *RinchanRuntime) putchar() {
	fmt.Printf("%c", r.memory[r.pointer])
}

func (r *RinchanRuntime) getchar() {
	_, err := fmt.Scanf("%c", &r.memory[r.pointer])
	if err != nil {
		return
	}
}
