package day04

type BingoBoard struct {
	values [][]int
	marked [][]bool
}

func NewBingoBoard(values [][]int) *BingoBoard {
	return &BingoBoard{
		values: values,
	}
}

func (b *BingoBoard) Draw(n int) bool {
	for y := 0; y < len(b.values); y++ {
		for x := 0; x < len(b.values[y]); x++ {
			// TODO same number again?
			if b.values[y][x] == n {
				b.applyDraw(x, y)
				return true
			}
		}
	}
	return false
}

func (b *BingoBoard) applyDraw(px int, py int) {
	// lazy init
	if len(b.marked) == 0 {
		b.marked = make([][]bool, len(b.values))
		for y := 0; y < len(b.values); y++ {
			b.marked[y] = make([]bool, len(b.values[y]))
		}
	}
	b.marked[py][px] = true
}

func (b *BingoBoard) HasMarkedLine() bool {
	if len(b.marked) == 0 {
		return false
	}
	return b.HasMarkedRow() || b.HasMarkedCol()
}

func (b *BingoBoard) HasMarkedRow() bool {
	if len(b.marked) == 0 {
		return false
	}
	for y := 0; y < len(b.marked); y++ {
		marked := true
		for x := 0; x < len(b.marked[y]); x++ {
			if !b.marked[y][x] {
				marked = false
				break
			}
		}
		if marked {
			return true
		}
	}
	return false
}

func (b *BingoBoard) HasMarkedCol() bool {
	if len(b.marked) == 0 {
		return false
	}
	for x := 0; x < len(b.marked[0]); x++ {
		marked := true
		for y := 0; y < len(b.marked[x]); y++ {
			if !b.marked[y][x] {
				marked = false
				break
			}
		}
		if marked {
			return true
		}
	}
	return false
}

func (b *BingoBoard) SumUnmarkedValues() int {
	result := 0
	for y := 0; y < len(b.values); y++ {
		for x := 0; x < len(b.values[y]); x++ {
			if !b.marked[y][x] {
				result += b.values[y][x]
			}
		}
	}
	return result
}
