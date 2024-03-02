package graph

type vector2 struct {
	X float64
	Y float64
}

type vector3 struct {
	X float64
	Y float64
	Z float64
}

type col struct {
	pts   []vector3
	count int
	color string
}
