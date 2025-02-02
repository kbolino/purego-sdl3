package sdl

type PenAxis uint32

const (
	PenAxisPressure PenAxis = iota
	PenAxisXtilt
	PenAxisYtilt
	PenAxisDistance
	PenAxisRotation
	PenAxisSlider
	PenAxisTangentialPressure
	PenAxisCount
)
