package square

import "math"

type sidesNumInt int

const (
	SidesTriangle sidesNumInt = 3
	SidesSquare   sidesNumInt = 4
	SidesCircle   sidesNumInt = 0
)

func CalcSquare(sideLen float64, sidesNum sidesNumInt) float64 {
	sideLenSqua := sideLen * sideLen
	switch sidesNum {
	case SidesTriangle:
		return sideLenSqua * math.Sqrt(3) / 4
	case SidesSquare:
		return sideLenSqua
	case SidesCircle:
		return math.Pi * sideLenSqua
	}

	return 0
}
