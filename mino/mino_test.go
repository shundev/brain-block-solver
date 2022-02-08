package mino

import "testing"

func TestMinoSplit(t *testing.T) {
	input := `
ooo..
.o...
.o...
.....
.....
`
	want := [][]byte{
		{'o', 'o', 'o', '.', '.'},
		{'.', 'o', '.', '.', '.'},
		{'.', 'o', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}

	mp := split(input)
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if want[y][x] != mp[y][x] {
				t.Fatalf("Wrong (y, x) = (%d, %d)\n", y, x)
			}
		}
	}
}

func TestMinoRotate(t *testing.T) {
	input := `
ooo..
.o...
.o...
.....
.....
`
	want := [][]byte{
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
		{'o', '.', '.', '.', '.'},
		{'o', 'o', 'o', '.', '.'},
		{'o', '.', '.', '.', '.'},
	}

	mino := &Mino{1, split(input)}
	rotated := mino.Rotate()

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if want[y][x] != rotated.Map[y][x] {
				t.Fatalf("Wrong (y, x) = (%d, %d)\n", y, x)
			}
		}
	}

	if rotated.ID != mino.ID {
		t.Fatalf("Wrong ID want=%d, got=%d\n", mino.ID, rotated.ID)
	}
}

func TestMinoFlip(t *testing.T) {
	input := `
ooo..
.o...
.o...
.....
.....
`
	want := [][]byte{
		{'.', '.', 'o', 'o', 'o'},
		{'.', '.', '.', 'o', '.'},
		{'.', '.', '.', 'o', '.'},
		{'.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.'},
	}

	mino := &Mino{1, split(input)}
	rotated := mino.Flip()

	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if want[y][x] != rotated.Map[y][x] {
				t.Fatalf("Wrong (y, x) = (%d, %d)\n", y, x)
			}
		}
	}

	if rotated.ID != mino.ID {
		t.Fatalf("Wrong ID want=%d, got=%d\n", mino.ID, rotated.ID)
	}
}
