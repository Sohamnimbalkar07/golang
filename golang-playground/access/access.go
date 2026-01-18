package data

var i int

var Str string = "Initial Name"

type Point struct {
	X int
	Y int
}

// pointer receiver to modify the original Point
// no need to deference p as Go does it automatically ( deferencing = *p.X = *(p).X + dx)
func (p *Point) Move(dx, dy int) {
	p.X += dx
	p.Y += dy
}

func modifyI() {
	i = 20
}

func ModifyName() {
	Str = "Modified Name"
}
