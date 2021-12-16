package day15

import "fmt"

type Cave struct {
	values [][]int
}

func (c *Cave) Init(maxX int, maxY int) {
	c.values = make([][]int, maxY)
	for i := 0; i < maxX; i++ {
		c.values[i] = make([]int, maxX)
	}
}

func (c *Cave) Each(f func(dx int, dy int)) {
	for y := 0; y < len(c.values); y++ {
		for x := 0; x < len(c.values[y]); x++ {
			f(x, y)
		}
	}
}

func (c *Cave) GetMax() (int, int) {
	return len(c.values[0]) - 1, len(c.values) - 1
}

func (c *Cave) ToString() string {
	result := ""
	for y := 0; y < len(c.values); y++ {
		for x := 0; x < len(c.values[y]); x++ {
			result += fmt.Sprintf("%d", c.values[y][x])
		}
		result += "\n"
	}
	return result
}

func (c *Cave) Adjacents(dx int, dy int, f func(x int, y int)) {
	for _, y := range []int{dy - 1, dy, dy + 1} {
		for _, x := range []int{dx - 1, dx, dx + 1} {
			if y == dy && x == dx {
				continue
			}
			if y < 0 || y > len(c.values)-1 {
				continue
			}
			if x < 0 || x > len(c.values[0])-1 {
				continue
			}
			f(x, y)
		}
	}
}

func (c *Cave) AdjacentsHV(dx int, dy int, f func(x int, y int)) {
	c.Adjacents(dx, dy, func(x int, y int) {
		if dx == x || dy == y {
			f(x, y)
		}
	})
}

func (c *Cave) Set(x int, y int, v int) {
	c.values[y][x] = v
}

func (c *Cave) Get(x int, y int) int {
	return c.values[y][x]
}

func (c *Cave) Expand(size int) {
	values := make([][]int, len(c.values)*size)
	for y := 0; y < len(c.values); y++ {
		{
			oldRow := c.values[y]
			newRow := make([]int, len(oldRow)*size)
			for x := 0; x < len(oldRow); x++ {
				for dx := 0; dx < size; dx++ {
					n := oldRow[x] + dx
					if n > 9 {
						n -= 9
					}
					newRow[x+(dx*len(oldRow))] = n
				}
			}
			values[y] = newRow
		}
		for dy := 1; dy < size; dy++ {
			fromRow := values[y]
			newRow := make([]int, len(fromRow))
			for x := 0; x < len(fromRow); x++ {
				n := fromRow[x] + dy
				if n > 9 {
					n -= 9
				}
				newRow[x] = n
			}
			values[y+(dy*len(c.values))] = newRow
		}
	}
	c.values = values
}
