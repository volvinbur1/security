package mt19937

import (
	"math"
)

const (
	N = 624
	m = 397
	W = 32
	r = 31
	U = 11
	S = 7
	T = 15
	L = 18
	f = 1812433253
	a = 0x9908B0DF
	B = 0x9D2C5680
	C = 0xEFC60000
	D = 0xFFFFFFFF
)

var lowerMask = int(math.Pow(2, r) - 1)
var upperMask = int(math.Pow(2, W-r)-1) << r

type Generator struct {
	index  int
	states []uint32
}

func New(seed uint32) Generator {
	g := Generator{}
	g.index = N
	g.states = make([]uint32, N)
	g.states[0] = seed

	for i := 1; i < N; i++ {
		g.states[i] = uint32((f*(int(g.states[i-1])^(int(g.states[i-1])>>(W-2))) + i) & int(math.Pow(2, W)-1))
	}

	return g
}

func (g *Generator) UInt32() uint32 {
	if g.index >= N {
		if g.index > N {
			panic("states index is out of range for mt19937")
		}

		g.twist()
	}

	res := g.temper(g.states[g.index])
	g.index++
	return res
}

func (g *Generator) SetStates(states []uint32) {
	g.states = states
}

func (g *Generator) twist() {
	for i := 0; i < N; i++ {
		temp := uint32((int(g.states[i]) & upperMask) + (int(g.states[(i+1)%N]) & lowerMask))
		tempSh := temp >> 1

		if temp%2 != 0 {
			tempSh ^= a
		}

		g.states[i] = g.states[(i+m)%N] ^ tempSh
	}

	g.index = 0
}

func (g *Generator) temper(val uint32) uint32 {
	val ^= (val >> U) & D
	val ^= (val << S) & B
	val ^= (val << T) & C
	val ^= val >> L
	return val
}
