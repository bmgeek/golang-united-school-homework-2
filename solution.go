package square

// Define custom int type to hold sides number and update CalcSquare signature by replacing #yourTypeNameHere#

// Define constants to represent 0, 3 and 4 sides.  Test uses mnemos: SidesTriangle(==3), SidesSquare(==4), SidesCircle(==0)
// it's like:
// CalcSquare(10.0, SidesTriangle)
// CalcSquare(10.0, SidesSquare)
// CalcSquare(10.0, SidesCircle)

import (
	"math"
)

var (
	SidesSquare   = 4
	SidesTriangle = 3
	SidesCircle   = 0
)

func CalcSquare(sideLen float64, sidesNum int) float64 {
	switch sidesNum {
	case 3:
		square := (sideLen * math.Sqrt(3)) / 4
		return square
	case 4:
		square := sideLen * sideLen
		return square
	case 0:
		square := sideLen * math.Pi
		return square
	default:
		return 0
	}
}
