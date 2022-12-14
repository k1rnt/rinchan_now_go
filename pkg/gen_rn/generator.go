package gen_rn

import (
	"fmt"
	"strings"
)

type RinchanGenerator struct {
	str string
}

func NewRinchanGenerator(str string) *RinchanGenerator {
	return &RinchanGenerator{
		str: str,
	}
}

func (g *RinchanGenerator) Generate() {
	bytes := []byte(g.str)
	s := 0

	for _, e := range bytes {
		if s < int(e) {
			k := int(e) - s
			t := rn_cal(k, string(incr))
			if len(t) < k {
				fmt.Print(t)
			} else {
				fmt.Print(strings.Repeat(string(incr), k))
			}
		} else if int(e) < s {
			k := s - int(e)
			t := rn_cal(k, string(decr))
			if len(t) < k {
				fmt.Print(t)
			} else {
				fmt.Print(strings.Repeat(string(decr), k))
			}
		}
		fmt.Print(string(putchar))
		s = int(e)
	}

}

func rn_cal(n int, sign string) string {
	if n <= 0 {
		return ""
	}
	if n <= 8 {
		return strings.Repeat(sign, n)
	}

	min := []int{4 * n}
	for p := 1; p <= n; p++ {
		q := n / p
		r := n % p

		k := 6 + p + q + r
		if k < min[0] {
			min = []int{k, p, q, r}
		}
	}

	return string(next) + strings.Repeat(string(incr), min[1]) + string(loopstart) + string(prev) + strings.Repeat(sign, min[2]) + string(next) + string(decr) + string(loopend) + string(prev) + strings.Repeat(sign, min[3])
}
