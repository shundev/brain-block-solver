package state

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"petromino/constant"
	"petromino/mino"
)

func TestInitState(t *testing.T) {
	st := New()
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			if st.Map[y][x] != -1 {
				t.Fatalf("Wrong state: (%d, %d)\n", y, x)
			}
		}
	}
}

func TestIsComplete(t *testing.T) {
	st := New()
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			st.Map[y][x] = 1
		}
	}

	if !st.IsComplete() {
		t.Fatalf("Wrong IsComplete")
	}

	st.Map[constant.H-1][constant.W-1] = -1

	if st.IsComplete() {
		t.Fatalf("Wrong IsComplete")
	}
}

func TestApplyOk(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, 0, 0)
	want := [][]int{
		{1, 1, 1, -1, -1, -1},
		{1, -1, 1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}

	if !ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}

	if err := eq(got.Map, want); err != nil {
		t.Fatal(err)
	}
}

func TestApplyOkOut(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, 8, 3)
	want := [][]int{
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, 1, 1, 1},
		{-1, -1, -1, 1, -1, 1},
	}

	if !ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}

	if err := eq(got.Map, want); err != nil {
		t.Fatal(err)
	}
}

func TestApplyOkOutLeft(t *testing.T) {
	init := New()
	mp := `
.....
.oooo
....o
.....
.....
`
	mino := &mino.Mino{9, split(mp)}
	got, ok := init.Apply(mino, -1, -1)
	want := [][]int{
		{9, 9, 9, 9, -1, -1},
		{-1, -1, -1, 9, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
		{-1, -1, -1, -1, -1, -1},
	}

	if !ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}

	if err := eq(got.Map, want); err != nil {
		t.Fatal(err)
	}
}

func TestApplyOkNoOverwrap(t *testing.T) {
	init := New()
	init.Fill(0)
	init.Map[1][2] = -1
	init.Map[1][3] = -1
	init.Map[1][4] = -1
	init.Map[2][2] = -1
	init.Map[2][4] = -1
	got, ok := init.Apply(mino.T1, 1, 2)
	want := [][]int{
		{0, 0, 0, 0, 0, 0},
		{0, 0, 1, 1, 1, 0},
		{0, 0, 1, 0, 1, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0},
	}

	if !ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}

	if err := eq(got.Map, want); err != nil {
		t.Fatal(err)
	}
}

func TestApplyNgOverwrap(t *testing.T) {
	init := New()
	init.Map[0][0] = 1
	got, ok := init.Apply(mino.T1, 0, 0)
	if ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}
}

func TestApplyNgOutRight(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, 8, 4)

	if ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}
}

func TestApplyNgOutDown(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, 9, 3)

	if ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}
}

func TestApplyNgOutUp(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, -1, 0)

	if ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}
}

func TestApplyNgOutLeft(t *testing.T) {
	init := New()
	got, ok := init.Apply(mino.T1, 0, -1)

	if ok {
		fmt.Println(got.String())
		t.Fatalf("Wrong ok")
	}
}

func eq(got, want [][]int) error {
	for y := 0; y < constant.H; y++ {
		for x := 0; x < constant.W; x++ {
			if got[y][x] != want[y][x] {
				return errors.New(fmt.Sprintf("Wrong (%d, %d): want=%d, got=%d", y, x, want[y][x], got[y][x]))
			}
		}
	}

	return nil
}

func split(s string) [][]byte {
	ss := strings.Split(strings.TrimSpace(s), "\n")
	ret := [][]byte{}
	for _, line := range ss {
		ret = append(ret, []byte(line))
	}

	return ret
}
