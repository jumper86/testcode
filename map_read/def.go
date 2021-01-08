package map_read

const fileName = "./2.obj"

type Triangle struct {
	A Point
	B Point
	C Point

	maxX float64
	minX float64
	maxZ float64
	minZ float64
}

type Point struct {
	X float64
	Y float64
	Z float64
}

//二维向量叉乘
func CrossProduct(p1, p2 Point) float64 {
	return p1.X*p2.Z - p1.Z*p2.X
}

//二维向量减法
func Minus(p1, p2 Point) Point {
	return Point{
		X: p1.X - p2.X,
		Z: p1.Z - p2.Z,
	}
}

func IsPointInTriangle(p Point, t Triangle) bool {
	pa := Minus(t.A, p)
	pb := Minus(t.B, p)
	pc := Minus(t.C, p)
	v1 := CrossProduct(pa, pb)
	v2 := CrossProduct(pb, pc)
	v3 := CrossProduct(pc, pa)
	return v1*v2 >= 0 && v1*v3 >= 0 && v2*v3 >= 0
}

var Draw [][]int32

var TopPoints []Point
var Triangles []Triangle

func simpleCheckNotInTriangle(p Point, t Triangle) bool {
	if p.X < t.minX || p.X > t.maxX || p.Z < t.minZ || p.Z > t.maxZ {
		return true
	}
	return false
}

func getRectangleXY(t Triangle) (float64, float64, float64, float64) {
	maxX := t.A.X
	minX := t.A.X
	maxZ := t.A.Z
	minZ := t.A.Z

	if t.B.X > maxX {
		maxX = t.B.X
	}
	if t.C.X > maxX {
		maxX = t.C.X
	}
	if t.B.X < minX {
		minX = t.B.X
	}
	if t.C.X < minX {
		minX = t.C.X
	}

	if t.B.Z > maxZ {
		maxZ = t.B.Z
	}
	if t.C.Z > maxZ {
		maxZ = t.C.Z
	}
	if t.B.Z < minZ {
		minZ = t.B.Z
	}
	if t.C.Z < minZ {
		minZ = t.C.Z
	}

	return maxX, minX, maxZ, minZ

}
