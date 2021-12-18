package day18

import (
	"fmt"
	"github.com/knalli/aoc"
)

type SfPairType int

const (
	SFPAIR_NESTED  SfPairType = 1
	SFPAIR_LITERAL SfPairType = 2
)

type SfPair struct {
	Type SfPairType

	Left  *SfPair
	Right *SfPair

	Value int
}

func (p *SfPair) ToString() string {
	if p.Type == SFPAIR_NESTED {
		return fmt.Sprintf("[%s,%s]", p.Left.ToString(), p.Right.ToString())
	} else {
		return fmt.Sprintf("%d", p.Value)
	}
}

func (p *SfPair) Add(s *SfPair) *SfPair {
	return &SfPair{
		Type:  SFPAIR_NESTED,
		Left:  p,
		Right: s,
	}
}

func (p *SfPair) Reduce() {
	//fmt.Printf("From: %s\n", p.ToString())
	for {
		if p.explodeIfNested([]*SfPair{p}, 4) {
			//fmt.Printf("After explode: %s\n", p.ToString())
			continue
		}
		if p.splitNumber(10) {
			//fmt.Printf("After split: %s\n", p.ToString())
			continue
		}
		break
	}
}

func (p *SfPair) explodeIfNested(parents []*SfPair, level int) bool {
	if p.Type == SFPAIR_NESTED {
		if level < len(parents) {
			if p.Left.Type != SFPAIR_LITERAL {
				panic("invalid type left")
			}
			if p.Right.Type != SFPAIR_LITERAL {
				panic("invalid type right")
			}
			{
				// find the first regular number to the left
				s := aoc.NewStack()
				for _, pp := range parents {
					s.Add(pp)
				}
				for !s.IsEmpty() {
					pi := s.Head().(*SfPair)
					if pi.Type == SFPAIR_LITERAL {
						pi.Value += p.Left.Value
						break
					}

					// forbidden (exclude the "wrong side")
					if pi.Left == p {
						continue
					}
					{
						forbidden := false
						for _, pp := range parents {
							if pi.Left == pp {
								forbidden = true
								break
							}
						}
						if forbidden {
							continue
						}
					}

					s.Add(pi.Left)
					if pi.Right != p {
						s.Add(pi.Right)
					}
				}
			}
			{
				// find the first regular number to the right
				s := aoc.NewStack()
				for _, pp := range parents {
					s.Add(pp)
				}
				for !s.IsEmpty() {
					pi := s.Head().(*SfPair)
					if pi.Type == SFPAIR_LITERAL {
						pi.Value += p.Right.Value
						break
					}

					// forbidden (exclude the "wrong side")
					if pi.Right == p {
						continue
					}
					{
						forbidden := false
						for _, pp := range parents {
							if pi.Right == pp {
								forbidden = true
								break
							}
						}
						if forbidden {
							continue
						}
					}

					s.Add(pi.Right)
					if pi.Left != p {
						s.Add(pi.Left)
					}
				}
			}
			p.Type = SFPAIR_LITERAL
			p.Value = 0
			p.Left = nil
			p.Right = nil
			return true
		}
		nextParents := make([]*SfPair, len(parents))
		for i, pp := range parents {
			nextParents[i] = pp
		}
		nextParents = append(nextParents, p)
		if p.Left.explodeIfNested(nextParents, level) {
			return true
		}
		return p.Right.explodeIfNested(nextParents, level)
	}
	return false
}

func (p *SfPair) splitNumber(threshold int) bool {
	if p.Type == SFPAIR_LITERAL {
		if p.Value >= threshold {
			p.Type = SFPAIR_NESTED
			p.Left = &SfPair{
				Type:  SFPAIR_LITERAL,
				Value: p.Value / 2,
			}
			p.Right = &SfPair{
				Type:  SFPAIR_LITERAL,
				Value: p.Value - p.Left.Value,
			}
			p.Value = -1
			return true
		}
	} else {
		if p.Left.splitNumber(threshold) {
			return true
		}
		return p.Right.splitNumber(threshold)
	}
	return false
}

func parseSfPair(input string) *SfPair {
	if input[0] == '[' && input[len(input)-1] == ']' {
		i := 1
		opens := 0
		found := false
		for {
			c := input[i]
			if c == '[' {
				opens++
			} else if c == ']' {
				opens--
			} else if c == ',' && opens == 0 {
				found = true
			}
			if opens == 0 && found {
				break
			}
			i++
		}
		return &SfPair{
			Type:  SFPAIR_NESTED,
			Left:  parseSfPair(input[1:i]),
			Right: parseSfPair(input[i+1 : len(input)-1]),
		}
	} else {
		return &SfPair{
			Type:  SFPAIR_LITERAL,
			Value: aoc.ParseInt(input),
		}
	}
}

func (p *SfPair) Magnitude() int {
	if p.Type == SFPAIR_LITERAL {
		return p.Value
	} else {
		return 3*p.Left.Magnitude() + 2*p.Right.Magnitude()
	}
}
