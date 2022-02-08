package mino

import "strings"

type (
	Mino struct {
		ID  int
		Map [][]byte
	}
)

const (
	t0 = `
ooooo
.....
.....
.....
.....
`
	t1 = `
ooo..
o.o..
.....
.....
.....
`
	t2 = `
o....
ooo..
.o...
.....
.....
`

	t3 = `
.o...
ooo..
.o...
.....
.....
`

	t4 = `
o....
oo...
.oo..
.....
.....
`

	t5 = `
oooo.
o....
.....
.....
.....
`

	t6 = `
ooo..
..oo.
.....
.....
.....
`

	t7 = `
oooo.
.o...
.....
.....
.....
`

	t8 = `
o....
ooo..
..o..
.....
.....
`

	t9 = `
ooo..
o....
o....
.....
.....
`

	t10 = `
oo...
oo...
o....
.....
.....
`

	t11 = `
ooo..
.o...
.o...
.....
.....
`
	t100 = `
ooooo
.....
.....
.....
.....
`
)

var (
	T0  = &Mino{0, split(t0)}
	T1  = &Mino{1, split(t1)}
	T2  = &Mino{2, split(t2)}
	T3  = &Mino{3, split(t3)}
	T4  = &Mino{4, split(t4)}
	T5  = &Mino{5, split(t5)}
	T6  = &Mino{6, split(t6)}
	T7  = &Mino{7, split(t7)}
	T8  = &Mino{8, split(t8)}
	T9  = &Mino{9, split(t9)}
	T10 = &Mino{10, split(t10)}
	T11 = &Mino{11, split(t11)}
)

func (m Mino) Rotate() Mino {
	mp := [][]byte{
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			mp[y][x] = m.Map[x][5-y-1]
		}
	}

	return Mino{m.ID, mp}
}

func (m Mino) Flip() Mino {
	mp := [][]byte{
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			mp[y][x] = m.Map[y][5-x-1]
		}
	}

	return Mino{m.ID, mp}
}

func split(s string) [][]byte {
	ss := strings.Split(strings.TrimSpace(s), "\n")
	ret := [][]byte{}
	for _, line := range ss {
		ret = append(ret, []byte(line))
	}

	return ret
}
