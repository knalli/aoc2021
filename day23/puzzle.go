package day23

import (
	"aoc2021/day05"
	"container/heap"
	"fmt"
	"github.com/knalli/aoc"
	"math"
	"strings"
)

type Grid struct {
	Data [][]int
}

func parseGrid(lines []string) *Grid {
	data := make([][]int, 0)
	length := 0
	positions := make(map[string][]day05.Point)
	for y, line := range lines {
		if length == 0 {
			length = len(line)
		}
		data2 := make([]int, 0)
		for x, c := range line {
			if c == '#' {
				data2 = append(data2, int(c))
			} else if c == ' ' {
				// fill:  not reachable
				data2 = append(data2, '#')
			} else {
				if c == 'A' || c == 'B' || c == 'C' || c == 'D' {
					s := string(c)
					if _, exist := positions[s]; !exist {
						positions[s] = make([]day05.Point, 0)
					}
					positions[s] = append(positions[s], day05.Point{X: x, Y: y})
					data2 = append(data2, int(c))
				} else {
					data2 = append(data2, '.')
				}
			}
		}
		for len(data2) < length {
			// fill: not reachable
			data2 = append(data2, '#')
		}
		data = append(data, data2)
	}
	return &Grid{
		Data: data,
	}
}

func (g *Grid) ToString() string {
	result := ""
	for _, line := range g.Data {
		for _, c := range line {
			result += fmt.Sprintf("%c", c)
		}
		result += "\n"
	}
	return result
}

func (g *Grid) ToStringWithPods(pods map[day05.Point]Pod) string {
	result := ""
	for y, line := range g.Data {
		for x, c := range line {
			if pod, exist := pods[day05.Point{X: x, Y: y}]; exist {
				result += fmt.Sprintf("%c", pod.Name)
			} else {
				result += fmt.Sprintf("%c", c)
			}
		}
		result += "\n"
	}
	return result
}

func (g *Grid) Each(f func(x int, y int, v int)) {
	for y := range g.Data {
		for x := range g.Data[y] {
			f(x, y, g.Data[y][x])
		}
	}
}

func (g *Grid) Valid(x int, y int) bool {
	return 0 <= y && y < len(g.Data) && 0 <= x && x < len(g.Data[0])
}

func (g *Grid) Adjacents(x, y int, f func(ax int, ay int, av int)) {
	for _, dy := range []int{-1, 0, 1} {
		for _, dx := range []int{-1, 0, 1} {
			if dx == 0 && dy == 0 {
				continue
			}
			ay := y + dy
			ax := x + dx
			if g.Valid(ax, ay) {
				f(ax, ay, g.Data[ay][ax])
			}
		}
	}
}

func (g *Grid) Get(x int, y int) int {
	return g.Data[y][x]
}

func (g *Grid) Set(x int, y int, v int) {
	g.Data[y][x] = v
}

type Candidate struct {
	Origin  day05.Point
	Targets []day05.Point
}

type State struct {
	Pods   map[day05.Point]Pod
	Energy int
}

type Pod struct {
	Id    int
	Name  int
	State PodState
}

type PodState int

const (
	STATE_INITIAL PodState = 0
	STATE_WAIT    PodState = 2
	STATE_FINAL   PodState = 4
)

func solve1(lines []string) error {
	grid := parseGrid(lines)

	minEnergy := solveViaPriorityQueue(grid)
	aoc.PrintSolution(fmt.Sprintf("Minimal energy used = %d", minEnergy))

	return nil
}

func solve2(lines []string) error {

	// patch input
	lines2 := make([]string, len(lines)+2)
	for i := range lines2 {
		if i < 3 {
			lines2[i] = lines[i]
		} else if i == 3 {
			lines2[i] = "  #D#C#B#A#"
		} else if i == 4 {
			lines2[i] = "  #D#B#A#C#"
		} else {
			lines2[i] = lines[i-2]
		}
	}
	lines = lines2

	grid := parseGrid(lines)

	minEnergy := solveViaPriorityQueue(grid)
	aoc.PrintSolution(fmt.Sprintf("Minimal energy used = %d", minEnergy))

	return nil
}

func solveViaFifoQueue(grid *Grid) int {
	q := aoc.NewQueue()

	validRoomsXCoords := map[int]int{
		'A': 3,
		'B': 5,
		'C': 7,
		'D': 9,
	}

	costFactors := map[int]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}

	{
		pods := make(map[day05.Point]Pod)
		grid.Each(func(x int, y int, v int) {
			if 'A' <= v && v <= 'D' {
				pos := day05.Point{X: x, Y: y}
				pod := Pod{
					Id:    len(pods),
					Name:  v,
					State: STATE_INITIAL,
				}
				// calibrate pods which are already final
				if validRoomsXCoords[pod.Name] == pos.X && pos.Y == 3 {
					pod.State = STATE_FINAL
				}
				pods[pos] = pod
				grid.Set(x, y, '.') // clear grid
			}
		})
		// verify
		fmt.Printf("%s\n", grid.ToStringWithPods(pods))
		q.Add(State{
			Pods: pods,
		})
	}

	filterPods := func(pods map[day05.Point]Pod, state PodState) map[day05.Point]Pod {
		result := make(map[day05.Point]Pod)
		for p, pod := range pods {
			if pod.State == state {
				result[p] = pod
			}
		}
		return result
	}

	minEnergy := math.MaxInt64

	for !q.IsEmpty() {
		state := q.Head().(State)

		if len(filterPods(state.Pods, STATE_FINAL)) == len(state.Pods) {
			if state.Energy < minEnergy {
				fmt.Printf("Found an energy minimum = %d\n", state.Energy)
				minEnergy = state.Energy
			}
			continue
		}

		if state.Energy >= minEnergy {
			// abort this one, because it cannot improve anymore
			continue
		}

		// rule #3: a pod can stay waiting; but if starts moving, it will/must reach the target
		for pos, pod := range filterPods(state.Pods, STATE_WAIT) {
			invalid := false
			targetX := validRoomsXCoords[pod.Name]
			targetY := -1
			costFactor := costFactors[pod.Name]
			cost := 0
			// pos.X == targetX IS NOT possible (rule #1)
			if pos.X > targetX {
				for x := pos.X - 1; x >= targetX; x-- {
					if _, used := state.Pods[day05.Point{X: x, Y: pos.Y}]; used {
						invalid = true
						break
					}
					cost += costFactor
				}
			} else if pos.X < targetX {
				for x := pos.X + 1; x <= targetX; x++ {
					if _, used := state.Pods[day05.Point{X: x, Y: pos.Y}]; used {
						invalid = true
						break
					}
					cost += costFactor
				}
			}
			if invalid {
				continue
			}
			if _, used := state.Pods[day05.Point{X: targetX, Y: pos.Y + 1}]; used {
				continue
			} else {
				cost += costFactor
				targetY = pos.Y + 1
			}
			if user, used := state.Pods[day05.Point{X: targetX, Y: pos.Y + 2}]; used {
				if user.State != STATE_FINAL {
					// there is a pod which isn't at it final place
					continue
				}
			} else {
				cost += costFactor
				targetY = pos.Y + 2
			}

			//fmt.Printf("? %s (%d/%d) => (%d/%d)\n", string(int32(pod.Name)), pos.X, pos.Y, targetX, targetY)
			newPods := make(map[day05.Point]Pod)
			for oldPos, oldPod := range state.Pods {
				if pod.Id == oldPod.Id {
					newPods[day05.Point{X: targetX, Y: targetY}] = Pod{
						Id:    pod.Id,
						Name:  pod.Name,
						State: STATE_FINAL,
					}
				} else {
					newPods[oldPos] = oldPod
				}
			}
			q.Add(State{
				Pods:   newPods,
				Energy: state.Energy + cost,
			})
		}

		for pos, pod := range filterPods(state.Pods, STATE_INITIAL) {
			invalid := false
			// a non-final pod has always targetX != pos.X
			costFactor := costFactors[pod.Name]
			targetY := 1
			cost := 0
			for y := pos.Y - 1; y >= targetY; y-- {
				if _, used := state.Pods[day05.Point{X: pos.X, Y: y}]; used {
					invalid = true
					break
				}
				cost += costFactor
			}
			if invalid {
				continue
			}
			for _, d := range []int{-1, 1} {
				costBase := cost
				for x := pos.X + d; ; x += d {
					if _, used := state.Pods[day05.Point{X: x, Y: targetY}]; used {
						break
					}
					if grid.Get(x, targetY) == '#' {
						break
					}
					costBase += costFactor

					if x != 3 && x != 5 && x != 7 && x != 9 { // rule #1
						//fmt.Printf("? %s (%d/%d) => (%d/%d)\n", string(int32(pod.Name)), pos.X, pos.Y, x, targetY)
						newPods := make(map[day05.Point]Pod)
						for oldPos, oldPod := range state.Pods {
							if pod.Id == oldPod.Id {
								newPods[day05.Point{X: x, Y: targetY}] = Pod{
									Id:    pod.Id,
									Name:  pod.Name,
									State: STATE_WAIT,
								}
							} else {
								newPods[oldPos] = oldPod
							}
						}
						q.Add(State{
							Pods:   newPods,
							Energy: state.Energy + costBase,
						})
					}
				}
			}
		}
	}

	return minEnergy
}

func solveViaPriorityQueue(grid *Grid) int {
	pq := make(PriorityQueue, 0)

	validRoomsXCoords := map[int]int{
		'A': 3,
		'B': 5,
		'C': 7,
		'D': 9,
	}

	costFactors := map[int]int{
		'A': 1,
		'B': 10,
		'C': 100,
		'D': 1000,
	}

	{
		pods := make(map[day05.Point]Pod)
		grid.Each(func(x int, y int, v int) {
			if 'A' <= v && v <= 'D' {
				pos := day05.Point{X: x, Y: y}
				pod := Pod{
					Id:    len(pods),
					Name:  v,
					State: STATE_INITIAL,
				}
				// calibrate pods which are already final (beneath must be a wall)
				if validRoomsXCoords[pod.Name] == pos.X && grid.Get(pos.X, pos.Y+1) == '#' {
					pod.State = STATE_FINAL
				}
				pods[pos] = pod
				grid.Set(x, y, '.') // clear grid
			}
		})
		// verify
		fmt.Printf("%s\n", grid.ToStringWithPods(pods))

		heap.Init(&pq)
		heap.Push(&pq, &Item{
			value: State{
				Pods: pods,
			},
			priority: 0,
		})
	}

	knownEnergyCosts := make(map[string]int)

	podsToString := func(pods map[day05.Point]Pod) string {
		list := make([]string, 0)
		for len(list) != len(pods) {
			search := len(list)
			for pos, pod := range pods {
				if pod.Id == search {
					list = append(list, fmt.Sprintf("%s=%d/%d", string(int32(pod.Name)), pos.X, pos.Y))
					break
				}
			}
		}
		return strings.Join(list, ";")
	}

	filterPods := func(pods map[day05.Point]Pod, state PodState) map[day05.Point]Pod {
		result := make(map[day05.Point]Pod)
		for p, pod := range pods {
			if pod.State == state {
				result[p] = pod
			}
		}
		return result
	}

	minEnergy := math.MaxInt64

	for pq.Len() > 0 {
		state := heap.Pop(&pq).(*Item).value

		if len(filterPods(state.Pods, STATE_FINAL)) == len(state.Pods) {
			if state.Energy < minEnergy {
				fmt.Printf("Found an energy minimum = %d\n", state.Energy)
				minEnergy = state.Energy
			}
			continue
		}

		if state.Energy >= minEnergy {
			// abort this one, because it cannot improve anymore
			continue
		}

		// rule #3: a pod can stay waiting; but if starts moving, it will/must reach the target
		for pos, pod := range filterPods(state.Pods, STATE_WAIT) {
			invalid := false
			targetX := validRoomsXCoords[pod.Name]
			targetY := -1
			costFactor := costFactors[pod.Name]
			cost := 0
			// pos.X == targetX IS NOT possible (rule #1)
			if pos.X > targetX {
				for x := pos.X - 1; x >= targetX; x-- {
					if _, used := state.Pods[day05.Point{X: x, Y: pos.Y}]; used {
						invalid = true
						break
					}
					cost += costFactor
				}
			} else if pos.X < targetX {
				for x := pos.X + 1; x <= targetX; x++ {
					if _, used := state.Pods[day05.Point{X: x, Y: pos.Y}]; used {
						invalid = true
						break
					}
					cost += costFactor
				}
			}
			if invalid {
				continue
			}

			for y := pos.Y + 1; ; y++ {
				if grid.Get(targetX, y) == '#' {
					break
				}
				if user, used := state.Pods[day05.Point{X: targetX, Y: y}]; used {
					if user.State != STATE_FINAL {
						invalid = true
						break
					}
					break
				}
				cost += costFactor
				targetY = y
			}
			if invalid {
				continue
			}
			if targetY == -1 {
				continue
			}

			//fmt.Printf("? %s (%d/%d) => (%d/%d)\n", string(int32(pod.Name)), pos.X, pos.Y, targetX, targetY)
			newPods := make(map[day05.Point]Pod)
			for oldPos, oldPod := range state.Pods {
				if pod.Id == oldPod.Id {
					newPods[day05.Point{X: targetX, Y: targetY}] = Pod{
						Id:    pod.Id,
						Name:  pod.Name,
						State: STATE_FINAL,
					}
				} else {
					newPods[oldPos] = oldPod
				}
			}
			nextEnergy := state.Energy + cost
			if nextEnergy < minEnergy {
				nextId := podsToString(newPods)
				if knownEnergy, exist := knownEnergyCosts[nextId]; !exist || knownEnergy > nextEnergy {
					knownEnergyCosts[nextId] = nextEnergy
					heap.Push(&pq, &Item{
						value: State{
							Pods:   newPods,
							Energy: nextEnergy,
						},
						priority: nextEnergy,
					})
				}
			}
		}

		for pos, pod := range filterPods(state.Pods, STATE_INITIAL) {
			invalid := false
			// a non-final pod has always targetX != pos.X
			costFactor := costFactors[pod.Name]
			targetY := 1
			cost := 0
			for y := pos.Y - 1; y >= targetY; y-- {
				if _, used := state.Pods[day05.Point{X: pos.X, Y: y}]; used {
					invalid = true
					break
				}
				cost += costFactor
			}
			if invalid {
				continue
			}
			for _, d := range []int{-1, 1} {
				costBase := cost
				for x := pos.X + d; ; x += d {
					if _, used := state.Pods[day05.Point{X: x, Y: targetY}]; used {
						break
					}
					if grid.Get(x, targetY) == '#' {
						break
					}
					costBase += costFactor

					if x != 3 && x != 5 && x != 7 && x != 9 { // rule #1
						//fmt.Printf("? %s (%d/%d) => (%d/%d)\n", string(int32(pod.Name)), pos.X, pos.Y, x, targetY)
						newPods := make(map[day05.Point]Pod)
						for oldPos, oldPod := range state.Pods {
							if pod.Id == oldPod.Id {
								newPods[day05.Point{X: x, Y: targetY}] = Pod{
									Id:    pod.Id,
									Name:  pod.Name,
									State: STATE_WAIT,
								}
							} else {
								newPods[oldPos] = oldPod
							}
						}
						nextEnergy := state.Energy + costBase
						if nextEnergy < minEnergy {
							nextId := podsToString(newPods)
							if knownEnergy, exist := knownEnergyCosts[nextId]; !exist || knownEnergy > nextEnergy {
								knownEnergyCosts[nextId] = nextEnergy
								heap.Push(&pq, &Item{
									value: State{
										Pods:   newPods,
										Energy: nextEnergy,
									},
									priority: nextEnergy,
								})
							}
						}
					}
				}
			}
		}
	}

	return minEnergy
}
