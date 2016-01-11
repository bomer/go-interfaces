package shapes

import "strconv"

//First Object, Square
type Square struct {
	Side int
}

func (s *Square) Area() string {
	return strconv.Itoa(s.Side * s.Side)
}
