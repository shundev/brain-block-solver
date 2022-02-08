package solver

import (
	"fmt"
	"os"

	"petromino/color"
	"petromino/constant"
	"petromino/mino"
	"petromino/state"
)

type (
	Solver struct {
		minos []*mino.Mino
		num   int
		tried int
		found int
		debug bool
	}
)

func New(debug bool) *Solver {
	minos := []*mino.Mino{mino.T0, mino.T1, mino.T2, mino.T3, mino.T4, mino.T5, mino.T6, mino.T7, mino.T8, mino.T9, mino.T10, mino.T11}
	s := &Solver{minos, len(minos), 0, 0, debug}
	return s
}

func (s *Solver) Solve() {
	initSt := state.New()
	s.tried = 0
	s.found = 0
	s.dfs(initSt)
}

func (s *Solver) dfs(state *state.State) {
	s.tried++

	if s.debug {
		s.indent(state.Idx)
		fmt.Println("---------->")
		state.Print(state.Idx)
	}

	if state.IsComplete() {
		s.found++
		fmt.Println(color.Sprintf(color.Green, "Found."))
		fmt.Printf("Tried: %d, Found: %d", s.tried, s.found)
		os.Exit(0)
		return
	}

	// すべて完了
	if state.Idx == s.num-1 {
		return
	}

	idx := state.Idx + 1
	orig := s.minos[idx]
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(orig, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r1 := orig.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r1, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r2 := r1.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r2, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r3 := r2.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r3, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r4 := orig.Flip()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r4, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r5 := r4.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r5, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r6 := r5.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r6, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}

	r7 := r6.Rotate()
	for y := -4; y < constant.H+4; y++ {
		for x := -4; x < constant.W+4; x++ {
			newS, ok := state.Apply(&r7, y, x)
			if ok && !newS.Invalid() {
				newS.Idx = idx
				s.dfs(newS)
			}
		}
	}
}

func (s *Solver) indent(num int) {
	for i := 0; i < num+1; i++ {
		fmt.Print("            ")
	}
}
