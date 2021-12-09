package mt19937

import (
	"math"
)

const (
	n = 624
	m = 397
	w = 32
	r = 31
	u = 11
	s = 7
	t = 15
	l = 18
	f = 1812433253
	a = uint32(0x9908B0DF)
	b = uint32(0x9D2C5680)
	c = uint32(0xEFC60000)
	d = uint32(0xFFFFFFFF)
)

var lowerMask = int(math.Pow(2, r) - 1)
var upperMask = int(math.Pow(2, w-r)-1) << r

type Generator struct {
	index  int
	states []uint32
}

func New(seed uint32) Generator {
	g := Generator{}
	g.index = n
	g.states = make([]uint32, n)
	g.states[0] = seed

	for i := 1; i < n; i++ {
		g.states[i] = uint32((f*(int(g.states[i-1])^(int(g.states[i-1])>>(w-2))) + i) & int(math.Pow(2, w)-1))
	}

	return g
}

func (g *Generator) UInt32() uint32 {
	if g.index >= n {
		if g.index > n {
			panic("states index is out of range for mt19937")
		}

		g.twist()
	}

	res := g.temper(g.states[g.index])
	g.index++
	return res
}

func (g *Generator) twist() {
	for i := 0; i < n; i++ {
		temp := uint32((int(g.states[i]) & upperMask) + (int(g.states[(i+1)%n]) & lowerMask))
		tempSh := temp >> 1

		if temp%2 != 0 {
			tempSh ^= a
		}

		g.states[i] = g.states[(i+m)%n] ^ tempSh
	}

	g.index = 0
}

func (g *Generator) temper(val uint32) uint32 {
	y := val ^ ((val >> u) & d)
	y = y ^ ((y << s) & b)
	y = y ^ ((y << t) & c)
	z := y ^ (y >> l)

	return z
}
