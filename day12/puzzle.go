package day12

import (
	"fmt"
	"github.com/knalli/aoc"
	"strings"
)

type Edge struct {
	From string
	To   string
}

type Node struct {
	Id   int
	Name string
	Big  bool
}

type Wrapper struct {
	NodeId int
	Path   []int
}

func cloneIntArray(ints []int) []int {
	res := make([]int, len(ints))
	for i, n := range ints {
		res[i] = n
	}
	return res
}

func parseNodes(lines []string) ([]Node, map[string]int, [][]int) {
	nodes := make([]Node, 0)
	refs := make(map[string]int)
	neighbours := make([][]int, 0)
	areNeighbours := func(from int, to int) bool {
		for _, nn := range neighbours[from] {
			if nn == to {
				return true
			}
		}
		return false
	}
	for _, line := range lines {
		split := strings.Split(line, "-")
		from := split[0]
		to := split[1]
		if _, found := refs[from]; !found {
			refs[from] = len(refs)
			big := true
			if from[0] > 'Z' {
				big = false
			}
			nodes = append(nodes, Node{
				Id:   len(nodes),
				Name: from,
				Big:  big,
			})
			neighbours = append(neighbours, []int{})
		}
		if _, found := refs[to]; !found {
			refs[to] = len(refs)
			big := true
			if to[0] > 'Z' {
				big = false
			}
			nodes = append(nodes, Node{
				Id:   len(nodes),
				Name: to,
				Big:  big,
			})
			neighbours = append(neighbours, []int{})
		}
		if !areNeighbours(refs[from], refs[to]) {
			neighbours[refs[from]] = append(neighbours[refs[from]], refs[to])
			neighbours[refs[to]] = append(neighbours[refs[to]], refs[from])
		}
	}
	return nodes, refs, neighbours
}

func pathToString(nodes []Node, path []int) string {
	res := ""
	for _, n := range path {
		res += nodes[n].Name + ","
	}
	return res[0 : len(res)-1]
}

func solve1(lines []string) error {
	nodes, refs, neighbours := parseNodes(lines)

	startId := refs["start"]
	endId := refs["end"]

	paths := bfs(nodes, neighbours, startId, endId, false)

	fmt.Printf("Found %d paths from the start to the end\n", len(paths))
	//for _, path := range paths {
	//	fmt.Println(pathToString(nodes, path))
	//}

	return nil
}

func solve2(lines []string) error {
	nodes, refs, neighbours := parseNodes(lines)

	startId := refs["start"]
	endId := refs["end"]

	paths := bfs(nodes, neighbours, startId, endId, true)

	fmt.Printf("Found %d paths from the start to the end\n", len(paths))
	//for _, path := range paths {
	//fmt.Println(pathToString(nodes, path))
	//}

	return nil
}

func bfs(nodes []Node, neighbours [][]int, startId int, endId int, part2 bool) [][]int {

	q := aoc.NewQueue()
	q.Add(Wrapper{
		NodeId: startId,
		Path:   []int{startId},
	})

	paths := make([][]int, 0)

	for !q.IsEmpty() {
		nodeWrapper := q.Head().(Wrapper)
		nodeId := nodeWrapper.NodeId
		nodePath := nodeWrapper.Path

		//fmt.Printf("node=%s, path=%s\n", nodes[nodeId].Name, pathToString(nodes, nodeWrapper.Path))

		// is this the end node?
		if nodeId == endId {
			//fmt.Println("Found end")
			paths = append(paths, nodePath)
			continue
		}

		for _, nextNodeId := range neighbours[nodeId] {
			if nextNodeId == startId {
				continue
			}
			nextNode := nodes[nextNodeId]
			if !part2 { // part1: small cave exactly once; big caves infinitely
				// check if this node can be visited (again)
				if !nextNode.Big {
					visited := 0
					for _, np := range nodePath {
						if np == nextNodeId {
							visited++
						}
					}
					if visited == 1 {
						continue
					}
				}
				nextPath := cloneIntArray(nodePath)
				nextPath = append(nextPath, nextNodeId)
				q.Add(Wrapper{
					NodeId: nextNodeId,
					Path:   nextPath,
				})
			} else {
				// part2: one small cave exactly two times, rest once; big caves infinitely
				if !nextNode.Big {
					visited := make(map[int]int)
					for _, p := range nodePath {
						if p == startId {
							//ignore
							continue
						}
						visited[p] = visited[p] + 1
					}
					// cave already visited? check if 2nd time is possible
					if visited[nextNodeId] == 1 {
						// check if already another small cave was two times visited
						anotherCaveAlreadyVisitedTwice := false
						for nodeVisitedId, v := range visited {
							if nodeVisitedId != nextNodeId && !nodes[nodeVisitedId].Big && v > 1 {
								anotherCaveAlreadyVisitedTwice = true
								break
							}
						}
						if anotherCaveAlreadyVisitedTwice {
							// no option
							continue
						}
					} else if visited[nextNodeId] > 1 {
						continue
					}
				}
				nextPath := cloneIntArray(nodePath)
				nextPath = append(nextPath, nextNodeId)
				q.Add(Wrapper{
					NodeId: nextNodeId,
					Path:   nextPath,
				})
			}
		}
	}
	return paths
}
