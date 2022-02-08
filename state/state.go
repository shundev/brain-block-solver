package state

import (
	"fmt"
	"petromino/color"
	"petromino/constant"
	"petromino/dsu"
	"petromino/mino"
	"strings"
)

type (
	State struct {
		Map [][]int
		Idx int
	}
)

var (
	colorMap = []color.Color{
		color.Red,
		color.Green,
		color.Orange,
		color.Yellow,
		color.Blue,
		color.Purple,
		color.Cyan,
		color.DarkGray,
		color.LightRed,
		color.LightGreen,
		color.LightBlue,
		color.LightPurple,
		color.LightCyan,
		color.LightGray,
		color.White,
	}
)

func New() *State {
	mp := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}
	return &State{mp, -1}
}

func (s *State) Fill(num int) {
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			s.Map[y][x] = num
		}
	}
}

func (s *State) IsComplete() bool {
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			if s.Map[y][x] == -1 {
				return false
			}
		}
	}

	return true
}

func (s *State) Copy() *State {
	cp := New()
	cp.Idx = s.Idx
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			cp.Map[y][x] = s.Map[y][x]
		}
	}
	return cp
}

func (s *State) Apply(mino *mino.Mino, y, x int) (*State, bool) {
	cp := s.Copy()

	// i)   out and o -> false
	// ii)  out and . -> ok
	// iii) not out and -1 and o -> ok
	// iv) not out and -1 and . -> ok
	// v) not out and >=0 and o -> false
	// vi) not out and >=0 and . -> ok
	for dy := 0; dy < 5; dy++ {
		for dx := 0; dx < 5; dx++ {
			newY := y + dy
			newX := x + dx
			if out(newY, newX) {
				// i)
				if mino.Map[dy][dx] == 'o' {
					return cp, false
				}

				// ii)
				continue
			}

			// iii) iv)
			if cp.Map[newY][newX] == -1 {
				if mino.Map[dy][dx] == 'o' {
					cp.Map[newY][newX] = mino.ID
					continue
				}
			}

			// v) cp.Map >= 0
			if mino.Map[dy][dx] == 'o' {
				return cp, false
			}

			// vi)
			// do nothing
		}
	}
	return cp, true
}

func out(y, x int) bool {
	return outH(y) || outW(x)
}

func outH(y int) bool {
	return y < 0 || constant.H <= y
}

func outW(x int) bool {
	return x < 0 || constant.W <= x
}

func (s *State) String() string {
	var sb strings.Builder
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			sb.WriteString(fmt.Sprintf("%d\t", s.Map[y][x]))
		}
		sb.WriteRune('\n')
	}

	return sb.String()
}

func (s *State) Print(indent int) {
	var sb strings.Builder
	for y := 0; y < constant.H; y++ {
		sb.WriteString(s.indent(indent))
		for x := 0; x < constant.W; x++ {
			if s.Map[y][x] == -1 {
				sb.WriteString(color.Sprintf(color.Black, "■ "))
			} else if s.Map[y][x] >= len(colorMap) {
				sb.WriteString(color.Sprintf(color.White, "■ "))
			} else {
				sb.WriteString(color.Sprintf(colorMap[s.Map[y][x]], "■ "))
			}
		}
		sb.WriteRune('\n')
	}

	fmt.Println(sb.String())
}

// Designates whether state is invalid.
// If invalid, no need to tranverse
func (s *State) Invalid() bool {
	return s.algorithm1()
}

// Every empty space (connected unoccupied squares) must have at least 5 squares.
func (s *State) algorithm1() bool {
	ds := dsu.New()
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			ds.Add(y*constant.W + x)
		}
	}

	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			if y > 0 && s.Map[y-1][x] == s.Map[y][x] {
				ds.Union(y*constant.W+x, (y-1)*constant.W+x)
			}

			if y < constant.H-1 && s.Map[y+1][x] == s.Map[y][x] {
				ds.Union(y*constant.W+x, (y+1)*constant.W+x)
			}

			if x > 0 && s.Map[y][x-1] == s.Map[y][x] {
				ds.Union(y*constant.W+x, y*constant.W+x-1)
			}

			if x < constant.W-1 && s.Map[y][x+1] == s.Map[y][x] {
				ds.Union(y*constant.W+x, y*constant.W+x+1)
			}
		}
	}

	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			if s.Map[y][x] != -1 {
				continue
			}

			if ds.Size(y*constant.W+x) <= 4 {
				return true
			}
		}
	}

	return false
}

func (s *State) indent(num int) string {
	var sb strings.Builder
	for i := 0; i < num+1; i++ {
		sb.WriteString("            ")
	}

	return sb.String()
}
